package repository

import (
	"context"
	"database/sql"
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
	return nil, nil
}
