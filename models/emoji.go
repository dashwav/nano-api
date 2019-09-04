package models

import "time"

// Profile holds specific application settings linked to an Account.
type Emoji struct {
	tableName struct{} `sql:"nanochan.emojis"`
	ID        int64 `sql:"emoji_id,pk" json:"id"`
	Name string `sql:"emoji_name"`
	MessageId int64 `sql:"message_id,pk"`
	ChannelId int64 `sql:"channel_id"`
	ChannelName string `sql:"channel_name"`
	UserId	int64 `sql:"user_id,pk"`
	UserName string `sql:"user_name"`
	TargetId  int64 `sql:"target_id"`
	TargetName string `sql:"target_name"`
	IsReaction bool `sql:"reaction,pk"`
	IsAnimated bool `sql:"animated,pk"`
	LogTime time.Time `sql:"logtime"`
}