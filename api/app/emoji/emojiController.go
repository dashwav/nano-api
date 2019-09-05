package emoji

import (
	"github.com/dashwav/nano-api/database"
	"github.com/dashwav/nano-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

// UserStore defines database operations for a profile.
type EmojiStore interface {
	Get(id int64) (*models.Emoji, error)
	GetTop(id int64, days int) ([]*models.Emoji, error)
	GetAll(animated bool, days int) ([]*database.Resp, error)
}

// UserResource implements user management handler.
type EmojiResource struct {
	Store EmojiStore
}

// NewUserResource creates and returns a profile resource.
func NewEmojiResource(store EmojiStore) *EmojiResource {
	return &EmojiResource{
		Store: store,
	}
}

func (rs *EmojiResource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{emojiId}/top/{days}", rs.getEmoji)
	r.Get("/top/{days}", rs.getAllEmoji)
	return r
}

func (rs *EmojiResource) getEmoji(w http.ResponseWriter, r *http.Request) {
	emojiId, err := strconv.ParseInt(chi.URLParam(r, "emojiId"), 10, 64)
	if err != nil {
		panic(err)
	}
	days, err := strconv.Atoi(chi.URLParam(r, "days"))
	if err != nil {
		panic(err)
	}
	_, err = rs.Store.GetTop(emojiId, days)
	if err != nil {
		panic(err)
	}
	//err = render.RenderList(w, r, NewEmojiListResponse(emojis))
	if err != nil {
		panic(err)
	}

}

func (rs *EmojiResource) getAllEmoji(w http.ResponseWriter, r *http.Request) {
	days, err := strconv.Atoi(chi.URLParam(r, "days"))
	if err != nil {
		panic(err)
	}
	count, err := rs.Store.GetAll(false, days)
	if err != nil {
		panic(err)
	}
	err = render.RenderList(w, r, NewEmojiListResponse(count))
	if err != nil {
		panic(err)
	}

}
