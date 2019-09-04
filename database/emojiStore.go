package database

import (
	"github.com/dashwav/nano-api/models"
	"github.com/go-pg/pg"
)

type EmojiStore struct {
	db *pg.DB
}

func NewEmojiStore(db *pg.DB) *EmojiStore {
	return &EmojiStore{
		db: db,
	}
}

func (s *EmojiStore) Get(id int64) (*models.Emoji, error) {
	e := models.Emoji{ID: id}
	err := s.db.Model(&e).Where("emoji_id = ?", id).Select()
	return &e, err
}