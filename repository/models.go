// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repository

import (
	"time"

	"github.com/google/uuid"
)

type Boardroom struct {
	ID        uuid.UUID
	Name      string
	Createdat time.Time
	Updatedat time.Time
}
