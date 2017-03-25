package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"context"
	"strconv"
	"sync"
)

func main() {
	r := chi.NewRouter()

	r.Route("/games", func(r chi.Router) {
		r.Get("/", ListGames)
		r.Post("/", CreateGame)

		r.Route("/:gameID", func(r chi.Router) {
			r.Use(GameCtx)
			r.Get("/", GetGame)
			r.Post("/hit", Hit)
			r.Post("/stand", Stand)
		})
	})

	http.ListenAndServe(":3333", r)
}

var games = []*Game{}
var gamesLock = sync.Mutex{}

func ListGames(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, games)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	gamesLock.Lock()
	defer gamesLock.Unlock()
	game := NewGame()
	game.ID = len(games)
	games = append(games, &game)
	render.JSON(w, r, game)
}

func GameCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gamesLock.Lock()
		defer gamesLock.Unlock()
		gameID := chi.URLParam(r, "gameID")
		id, err := strconv.Atoi(gameID)

		if err != nil || id >= len(games) {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "game", games[id])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value("game").(*Game)
	game.Lock()
	defer game.Unlock()
	render.JSON(w, r, game)
}

func gameAction(w http.ResponseWriter, r *http.Request, handle func(g *Game)) {
	game := r.Context().Value("game").(*Game)
	game.Lock()
	defer game.Unlock()
	if game.IsFinished {
		http.Error(w, "Game is finished", 400)
		return
	}
	handle(game)
	render.JSON(w, r, game)
}
func Hit(w http.ResponseWriter, r *http.Request) {
	gameAction(w, r, func (g *Game) {
		g.Hit()
	})
}

func Stand(w http.ResponseWriter, r *http.Request) {
	gameAction(w, r, func (g *Game) {
		g.Stand()
	})
}

