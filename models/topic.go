package models

import "time"

// Topic maps to the public.topics table
type Topic struct {
	tableName struct{} `pg:"topics"` // Optional: `public.topics` if schema needed

	ID        int64            `pg:"id,pk" json:"id"`
	Title     string           `pg:"title" json:"title"`
	Content   string           `pg:"content" json:"content"`
	CreatedAt time.Time        `json:"created_at"`
	Comments  []*TopicsComment `pg:"rel:has-many,fk:topics_id"`
}
