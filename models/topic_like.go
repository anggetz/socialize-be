package models

import (
	"time"
)

type TopicsLike struct {
	tableName struct{} `pg:"topics_like"` // maps to table name

	ID        int64     `pg:"id,pk"`
	Author    string    `pg:"author"`
	TopicsID  int64     `pg:"topics_id"`
	CreatedAt time.Time `pg:"created_at"`

	// Optional: to preload the topic this like belongs to
	// Topic     *Topic    `pg:"rel:has-one,fk:topics_id"`
}
