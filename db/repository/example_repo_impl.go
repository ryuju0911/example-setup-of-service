package repository

import (
	"context"
	"database/sql"
	"time"
)

// exampleRepoImpl implements example/repository/example_repo.
type exampleRepoImpl struct {
	db *sql.DB
}

// NewExampleRepoImpl returns the example repository implementation.
func NewExampleRepoImpl(db *sql.DB) ExampleRepo {
	return &exampleRepoImpl{db}
}

// CreateSomeData creates some data entry and inserts it to 'example' table.
func (r *exampleRepoImpl) CreateSomeData(ctx context.Context, id string, content string) (*SomeData, error) {
	query := "INSERT INTO example (id, content) VALUES ($1, $2) RETURNING created_at"
	var created_at time.Time

	err := r.db.QueryRowContext(ctx, query, id, content).Scan(&created_at)
	if err != nil {
		return nil, err
	}

	some_data := SomeData{
		Id:        id,
		Content:   content,
		CreatedAt: created_at,
	}

	return &some_data, nil
}
