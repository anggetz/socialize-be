package models

import (
	"time"
)

type TopicsComment struct {
	tableName struct{} `pg:"topics_comment"`

	ID        int64     `pg:"id,pk" json:"id"`
	Comment   string    `pg:"comment" json:"comment"`
	Author    string    `pg:"author" json:"author"`
	CreatedAt time.Time `pg:"created_at" json:"created_at"`
	TopicsID  int64     `pg:"topics_id" json:"topics_id"`

	// Optional: Join with Topics
	Topic *Topic `pg:"rel:has-one,fk:topics_id"`
}
