// Package app ties together application resources and handlers.
package app

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/dashwav/nano-api/api/app/emoji"
	"github.com/dashwav/nano-api/database"
	"github.com/dashwav/nano-api/logging"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxProfile
)

// API provides application resources and handlers.
type API struct {
	Emoji *emoji.EmojiResource
}

// NewAPI configures and returns application API.
func NewAPI(db *pg.DB) (*API, error) {

	userStore := database.NewEmojiStore(db)
	emojiApi := emoji.NewEmojiResource(userStore)

	api := &API{
		Emoji: emojiApi,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/emojis", a.Emoji.Router())

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
