package api

import (
	"net/http"
	"time"

	"github.com/dashwav/nano-api/api/app"
	"github.com/dashwav/nano-api/database"
	"github.com/dashwav/nano-api/logging"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// New configures application resources and routes.
func New(enableCORS bool) (*chi.Mux, error) {
	logger := logging.NewLogger()

	// Initialize DB w/ default values
	db, err := database.DBConn()
	if err != nil {
		logger.WithField("module", "database").Error(err)
		return nil, err
	}

	appAPI, err := app.NewAPI(db)
	if err != nil {
		logger.WithField("module", "app").Error(err)
		return nil, err
	}


	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(logging.NewStructuredLogger(logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if enableCORS {
		r.Use(corsConfig().Handler)
	}

	r.Group(func(r chi.Router) {
		r.Mount("/v1", appAPI.Router())
	})

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			logger.WithField("module", "router").Error(err)
		}
	})

	return r, nil
}

func corsConfig() *cors.Cors {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	return cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	})
}