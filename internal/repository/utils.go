package repository

import "github.com/aidarkhanov/nanoid"

func newID() *string {
	newID := nanoid.New()
	return &newID
}
