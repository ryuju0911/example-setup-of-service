package service

import (
	"context"
	"database/sql"
	"example/db/repository"
	example "example/gen/example"
	"log"
	"time"
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
func (s *exampleSvc) CreateSomeData(
	ctx context.Context,
	p *example.CreateSomeDataPayload,
) (res *example.SomeData, err error) {
	some_data, err := s.someDataRepo.CreateSomeData(ctx, p.ID, p.Content)
	if err != nil {
		s.logger.Printf("example.CreateSomeData: failed (%s)", err)
		return nil, example.MakeBadRequest(err)
	}

	res = mapRepoSomeDataToSvcSomeData(some_data)
	s.logger.Print("example.CreateSomeData")
	return
}

func mapRepoSomeDataToSvcSomeData(some_data *repository.SomeData) *example.SomeData {
	return &example.SomeData{
		ID:        some_data.Id,
		Content:   some_data.Content,
		CreatedAt: some_data.CreatedAt.Format(time.RFC3339),
	}
}
