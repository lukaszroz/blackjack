package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
)

var addr string

func TestMain(m *testing.M) {
	listenOnAvailablePort()
	retCode := m.Run()
	os.Exit(retCode)
}
func TestRootUrl(t *testing.T) {
	t.Parallel()
	resp := get(t, "http://"+addr)
	defer resp.Body.Close()
	assertEquals(t, "Expected status %v, got %v", http.StatusNotFound, resp.StatusCode)
}

func TestListGames(t *testing.T) {
	t.Parallel()
	resp := get(t, "http://"+addr+"/games")
	defer resp.Body.Close()
	assertEquals(t, "Expected status %v, got %v", http.StatusOK, resp.StatusCode)
	b := body(t, resp)
	var games []Game
	err := json.Unmarshal(b, &games)
	assertEquals(t, "Expected %v error, got %v", nil, err)
}

func TestCreateGame(t *testing.T) {
	t.Parallel()
	resp := post(t, "http://"+addr+"/games")
	defer resp.Body.Close()
	assertEquals(t, "Expected status %v, got %v", http.StatusOK, resp.StatusCode)
	b := body(t, resp)
	g := game(t, b)
	assertEquals(t, "Expected player to have %v cards, got %v", 2, len(g.Player.Cards))
	assertEquals(t, "Expected dealer to show %v cards, got %v", 1, len(g.Dealer.Cards))
	assertGameState(t, g, false, false, false)
}

func TestGetGame(t *testing.T) {
	t.Parallel()
	ID := getNewGameID(t)
	resp := get(t, fmt.Sprintf("http://%s/games/%d", addr, ID))
	defer resp.Body.Close()
	assertEquals(t, "Expected status %v, got %v", http.StatusOK, resp.StatusCode)
	b := body(t, resp)
	g := game(t, b)
	assertEquals(t, "Expected player to have %v cards, got %v", 2, len(g.Player.Cards))
	assertEquals(t, "Expected dealer to show %v cards, got %v", 1, len(g.Dealer.Cards))
	assertGameState(t, g, false, false, false)

}
func TestHit(t *testing.T) {
	t.Parallel()
	ID := getNewGameID(t)
	resp := post(t, fmt.Sprintf("http://%s/games/%d/hit", addr, ID))
	defer resp.Body.Close()
	assertEquals(t, "Expected status %v, got %v", http.StatusOK, resp.StatusCode)
	b := body(t, resp)
	g := game(t, b)
	assertEquals(t, "Expected player to have %v cards, got %v", 3, len(g.Player.Cards))
	assertEquals(t, "Expected dealer to show %v cards, got %v", 1, len(g.Dealer.Cards))

}
func TestBust(t *testing.T) {
	t.Parallel()
	ID := getNewGameID(t)
	url := fmt.Sprintf("http://%s/games/%d/hit", addr, ID)
	resp := post(t, url)
	defer resp.Body.Close()
	cardsN := 2
	for resp.StatusCode == http.StatusOK {
		cardsN++
		resp = post(t, url)
		defer resp.Body.Close()
	}
	assertEquals(t, "Expected status %v, got %v", http.StatusBadRequest, resp.StatusCode)
	resp = get(t, fmt.Sprintf("http://%s/games/%d", addr, ID))
	defer resp.Body.Close()
	b := body(t, resp)
	g := game(t, b)
	assertEquals(t, "Expected player to have %v cards, got %v", cardsN, len(g.Player.Cards))
	assertEquals(t, "Expected dealer to show %v cards, got %v", 1, len(g.Dealer.Cards))
	assertGameState(t, g, true, false, false)
}
func TestStand(t *testing.T) {
	t.Parallel()
	ID := getNewGameID(t)
	resp := post(t, fmt.Sprintf("http://%s/games/%d/stand", addr, ID))
	defer resp.Body.Close()
	b := body(t, resp)
	g := game(t, b)
	assertEquals(t, "Expected player to have %v cards, got %v", 2, len(g.Player.Cards))
	if 2 > len(g.Dealer.Cards) {
		t.Fatalf("Expected dealer to have more than %v cards, got %v", 2, len(g.Dealer.Cards))
	}
	if 17 > g.Dealer.Score.value {
		t.Fatalf("Expected dealer to have higher than %v score, got %v", 17, g.Dealer.Score.value)
	}
	assertIsFinished(t, g, true)
}
func BenchmarkHitAndStand(b *testing.B) {
	wins, plays := 0, 0
	for i := 0; i < b.N; i++ {
		i := 1
		for g := play(); !g.HasPlayerWon; g = play() {
			i++
		}
		wins++
		plays += i
	}
	b.Logf("Played %v games, won %v or %.2f%%", plays, wins, float64(wins)/float64(plays)*100)
}
func play() Game {
	ID := getNewGameID(nil)
	resp := post(nil, fmt.Sprintf("http://%s/games/%d/hit", addr, ID))
	defer resp.Body.Close()
	resp = post(nil, fmt.Sprintf("http://%s/games/%d/stand", addr, ID))
	defer resp.Body.Close()
	resp = get(nil, fmt.Sprintf("http://%s/games/%d", addr, ID))
	defer resp.Body.Close()
	body := body(nil, resp)
	return game(nil, body)
}
func getNewGameID(t *testing.T) int {
	resp := post(t, "http://"+addr+"/games")
	defer resp.Body.Close()
	b := body(t, resp)
	g := game(t, b)
	return g.ID
}
func listenOnAvailablePort() {
	hs := http.Server{Addr: ":0", Handler: newServer().router}
	ln, err := net.Listen("tcp", hs.Addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	go func() {
		hs.Serve(ln)
	}()
	addr = ln.Addr().String()
}

func assertEquals(t *testing.T, format string, expected, actual interface{}) {
	if expected != actual {
		t.Fatalf(format, expected, actual)
	}
}

func assertGameState(t *testing.T, g Game, finished bool, tie bool, won bool) {
	assertIsFinished(t, g, finished)
	assertEquals(t, "Expected IsTie to be %v, got %v", tie, g.IsTie)
	assertEquals(t, "Expected HasPlayerWon to be %v, got %v", won, g.HasPlayerWon)
}

func assertIsFinished(t *testing.T, g Game, f bool) {
	assertEquals(t, "Expected IsFinished to be %v, got %v", f, g.IsFinished)
}

func get(t *testing.T, url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	return resp
}

func body(t *testing.T, resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func post(t *testing.T, url string) *http.Response {
	resp, err := http.Post(url, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	return resp
}

func game(t *testing.T, buf []byte) Game {
	var g Game
	if err := json.Unmarshal(buf, &g); err != nil {
		t.Fatal(err)
	}
	return g
}
