package database

import (
	"github.com/dashwav/nano-api/models"
	"github.com/go-pg/pg"
	"time"
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

func (s *EmojiStore) GetTop(emojiId int64, days int) ([]*models.Emoji, error) {
	emojis := []*models.Emoji{}
	negDays := 0 - days
	date := time.Now().AddDate(0, 0, negDays)
	err := s.db.Model(&emojis).Where("emoji_id = ?", emojiId).Where("logtime > ?", date).Select()
	if err != nil {
		return nil, err
	}
	return emojis, nil
}