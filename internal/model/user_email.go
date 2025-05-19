package model

import "time"

type UserEmail struct {
	ID          string    `json:"id"`
	User        string    `json:"user"`
	Address     string    `json:"address"`
	VerifyToken string    `json:"verify_token"`
	VerifiedAt  string    `json:"verified_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
