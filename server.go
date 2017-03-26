package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"context"
	"strconv"
	"sync"
	"github.com/pressly/chi/middleware"
)

func main() {
	http.ListenAndServe(":3333", newServer().router)
}

type server struct {
	sync.RWMutex
	router *chi.Mux
	games  []*Game
}

func newServer() server {
	s := server{}
	s.games = []*Game{}
	s.router = newRouter(&s)
	return s
}

const (
	game_id_key = "gameID"
	game_key    = "game_key"
)

func newRouter(s *server) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/games", func(r chi.Router) {
		r.Get("/", s.listGames)
		r.Post("/", s.createGame)

		r.Route("/:"+game_id_key, func(r chi.Router) {
			r.Use(s.gameCtx)
			r.Get("/", s.getGame)
			r.Post("/hit", s.hit)
			r.Post("/stand", s.stand)
		})
	})
	return r
}

func (s *server) listGames(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()
	for _, g := range s.games {
		g.RLock()
		defer g.RUnlock()
	}
	render.JSON(w, r, s.games)
}

func (s *server) createGame(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	defer s.Unlock()
	game := NewGame()
	game.ID = len(s.games)
	s.games = append(s.games, &game)
	render.JSON(w, r, game)
}

func (s *server) gameCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gameID := chi.URLParam(r, game_id_key)
		id, err := strconv.Atoi(gameID)
		s.RLock()
		defer s.RUnlock()
		if err != nil || id >= len(s.games) {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), game_key, s.games[id])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *server) getGame(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(game_key).(*Game)
	game.RLock()
	defer game.RUnlock()
	render.JSON(w, r, game)
}

func gameAction(w http.ResponseWriter, r *http.Request, handle func(g *Game) error) {
	game := r.Context().Value(game_key).(*Game)
	game.Lock()
	defer game.Unlock()
	if e := handle(game); e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	render.JSON(w, r, game)
}
func (s *server) hit(w http.ResponseWriter, r *http.Request) {
	gameAction(w, r, func(g *Game) error {
		return g.Hit()
	})
}

func (s *server) stand(w http.ResponseWriter, r *http.Request) {
	gameAction(w, r, func(g *Game) error {
		return g.Stand()
	})
}
