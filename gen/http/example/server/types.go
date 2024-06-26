// Code generated by goa v3.16.2, DO NOT EDIT.
//
// example HTTP server types
//
// Command:
// $ goa gen example/design

package server

import (
	example "example/gen/example"

	goa "goa.design/goa/v3/pkg"
)

// CreateSomeDataRequestBody is the type of the "example" service
// "CreateSomeData" endpoint HTTP request body.
type CreateSomeDataRequestBody struct {
	ID      *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
}

// CreateSomeDataResponseBody is the type of the "example" service
// "CreateSomeData" endpoint HTTP response body.
type CreateSomeDataResponseBody struct {
	ID        string `form:"id" json:"id" xml:"id"`
	Content   string `form:"content" json:"content" xml:"content"`
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// CreateSomeDataBadRequestResponseBody is the type of the "example" service
// "CreateSomeData" endpoint HTTP response body for the "BadRequest" error.
type CreateSomeDataBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewCreateSomeDataResponseBody builds the HTTP response body from the result
// of the "CreateSomeData" endpoint of the "example" service.
func NewCreateSomeDataResponseBody(res *example.SomeData) *CreateSomeDataResponseBody {
	body := &CreateSomeDataResponseBody{
		ID:        res.ID,
		Content:   res.Content,
		CreatedAt: res.CreatedAt,
	}
	return body
}

// NewCreateSomeDataBadRequestResponseBody builds the HTTP response body from
// the result of the "CreateSomeData" endpoint of the "example" service.
func NewCreateSomeDataBadRequestResponseBody(res *goa.ServiceError) *CreateSomeDataBadRequestResponseBody {
	body := &CreateSomeDataBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateSomeDataPayload builds a example service CreateSomeData endpoint
// payload.
func NewCreateSomeDataPayload(body *CreateSomeDataRequestBody) *example.CreateSomeDataPayload {
	v := &example.CreateSomeDataPayload{
		ID:      *body.ID,
		Content: *body.Content,
	}

	return v
}

// ValidateCreateSomeDataRequestBody runs the validations defined on
// CreateSomeDataRequestBody
func ValidateCreateSomeDataRequestBody(body *CreateSomeDataRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "body"))
	}
	return
}
