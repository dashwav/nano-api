package emoji

import (
	"github.com/dashwav/nano-api/models"
	"github.com/go-chi/chi"
	"net/http"
)

// UserStore defines database operations for a profile.
type EmojiStore interface {
	Get(id int64) (*models.Emoji, error)
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
	r.Get("/{emojiId}", rs.getEmoji)
	return r
}

func (rs *EmojiResource) getEmoji(w http.ResponseWriter, r *http.Request) {

}
