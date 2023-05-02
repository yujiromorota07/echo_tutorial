package entity

import "time"

type TodoID uint32

type Todo struct {
	ID        uint32
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
