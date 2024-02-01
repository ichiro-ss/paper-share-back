package data

import "time"

type User struct {
	id        int       `json:"id"`
	name      string    `json:"name"`
	createdAt time.Time `json:"datetime"`
	updatedAt time.Time `json:"datetime"`
	deletedAt time.Time `json:"datetime"`
}
