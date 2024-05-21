package service

import (
	"context"
	"database/sql"
	"example/db/repository"
	example "example/gen/example"
	"log"
)

// example service example implementation.
// The example methods log the requests and return zero values.
type exampleSvc struct {
	someDataRepo repository.ExampleRepo
	logger       *log.Logger
}

// NewExample returns the example service implementation.
func NewExampleSvc(db *sql.DB, logger *log.Logger) example.Service {
	someDataRepo := repository.NewExampleRepoImpl(db)
	return &exampleSvc{someDataRepo, logger}
}

// CreateSomeData implements CreateSomeData.
func (s *exampleSvc) CreateSomeData(ctx context.Context, p *example.CreateSomeDataPayload) (res *example.SomeData, err error) {
	res = &example.SomeData{}
	s.logger.Print("example.CreateSomeData")
	return
}
