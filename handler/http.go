package handler

import (
	"net/http"
	"time"

	"github.com/mafuyuk/ddd-go-api-template/domain/service"
	mw "github.com/mafuyuk/ddd-go-api-template/handler/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/render"
	"gopkg.in/go-playground/validator.v9"
)

var rendering = render.New(render.Options{})
var validate = validator.New()

// Services アプリケーションサービスをまとめる構造体
type Services struct {
	UserService service.UserService
}

// New ドメインのサービスを追加したServerの構造体を返す関数
func New(addr string, services *Services) *Server {
	return &Server{
		Server: http.Server{
			Addr: addr,
		},
		Services: *services,
	}
}

// Server HTTPのサーバそのものを表す構造体
type Server struct {
	http.Server
	Services
}

// ListenAndServe はServerを起動するメソッド
func (s *Server) ListenAndServe() error {
	r := chi.NewRouter()

	// CORS対応
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	// 公式提供のmiddleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// 独自のmiddleware
	r.Use(mw.AuthMiddleware)

	// ルーティング
	r.Route("/v1", func(r chi.Router) {

		// User
		r.Route("/users", func(r chi.Router) {
			r.Post("/", s.registerUser)
		})
	})

	s.Handler = r
	return s.Server.ListenAndServe()
}
