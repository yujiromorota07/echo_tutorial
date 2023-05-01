package entity

import "time"

type Todo struct {
	ID        uint32
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
