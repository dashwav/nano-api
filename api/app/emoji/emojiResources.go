package emoji

import (
	"github.com/dashwav/nano-api/models"
	"github.com/go-chi/render"
	"net/http"
)

type EmojiResponse struct {
	*models.Emoji
	Elapsed int64 `json:"elapsed"`
}

func NewEmojiResponse(emoji *models.Emoji) *EmojiResponse {
	resp := &EmojiResponse{Emoji: emoji}
	return resp
}

func (rd *EmojiResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}

func NewEmojiListResponse(emojis []*models.Emoji) []render.Renderer {
	list := []render.Renderer{}
	for _, emoji := range emojis {
		list = append(list, NewEmojiResponse(emoji))
	}
	return list
}

