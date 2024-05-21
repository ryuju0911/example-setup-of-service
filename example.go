package exampleapi

import (
	"context"
	example "example/gen/example"
	"log"
)

// example service example implementation.
// The example methods log the requests and return zero values.
type examplesrvc struct {
	logger *log.Logger
}

// NewExample returns the example service implementation.
func NewExample(logger *log.Logger) example.Service {
	return &examplesrvc{logger}
}

// CreateSomeData implements CreateSomeData.
func (s *examplesrvc) CreateSomeData(ctx context.Context, p *example.CreateSomeDataPayload) (res *example.SomeData, err error) {
	res = &example.SomeData{}
	s.logger.Print("example.CreateSomeData")
	return
}
