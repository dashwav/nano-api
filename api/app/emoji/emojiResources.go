package emoji

import (
	"github.com/dashwav/nano-api/database"
	"github.com/go-chi/render"
	"sort"
)

func NewEmojiListResponse(emojis []*database.Resp) []render.Renderer {
	sort.SliceStable(emojis, func(i, j int) bool {
		return emojis[i].EmojiCount > emojis[j].EmojiCount
	})
	list := []render.Renderer{}
	for _, emoji := range emojis {
		list = append(list, emoji)
	}
	return list
}

