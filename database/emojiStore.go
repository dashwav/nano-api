package database

import (
	"github.com/dashwav/nano-api/models"
	"github.com/go-pg/pg"
	"net/http"
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

type Resp struct {
	EmojiId int64
	EmojiName string
	EmojiCount int
}

func (rd *Resp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
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

func (s *EmojiStore) GetAll(animated bool, days int) ([]*Resp, error) {
	emoji := models.Emoji{IsAnimated:animated}
	response := []*Resp{}
	negDays := 0 - days
	date := time.Now().AddDate(0, 0, negDays)
	err := s.db.Model(&emoji).
		Where("animated = ?", animated).
		Where("logtime > ?", date).
		Group("emoji_id").
		Group("emoji_name").
		Column("emoji_id").
		Column("emoji_name").
		ColumnExpr("count(*) as emoji_count").
		Select(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}