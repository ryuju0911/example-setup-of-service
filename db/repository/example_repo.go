package repository

import (
	"context"
	"time"
)

type ExampleRepo interface {
	CreateSomeData(ctx context.Context, id string, content string) (*SomeData, error)
}

// SomeData represents an entry of 'example' table.
type SomeData struct {
	Id        string
	Content   string
	CreatedAt time.Time
}
