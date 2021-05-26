package main

import (
	"net/http"
	"time"
)

// swagger:model ThingResponse
type ThingResponse struct {
	// The UUID of a thing
	// example: 6204037c-30e6-408b-8aaa-dd8219860b4b
	UUID string `json:"uuid"`

	// The Name of a thing
	// example: Some name
	Name string `json:"name"`

	// The Value of a thing
	// example: Some value
	Value string `json:"value"`

	// The last time a thing was updated
	// example: 2021-05-25T00:53:16.535668Z
	Updated time.Time `json:"updated"`

	// The time a thing was created
	// example: 2021-05-25T00:53:16.535668Z
	Created time.Time `json:"created"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	// The error message
	// example: An error occurred
	Error string `json:"error"`
}

// swagger:parameters get-thing update-thing delete-thing
type _ struct {
	// The UUID of a thing
	// in:path
	UUID string `json:"uuid"`
}

// swagger:route GET /thing/{uuid} Thing get-thing
//
// This is the summary for getting a thing by its UUID
//
// This is the description for getting a thing by its UUID. Which can be longer.
//
// responses:
//   200: ThingResponse
//   404: ErrorResponse
//   500: ErrorResponse
func (s *Server) GetThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

// swagger:model CreateThing
type CreateThing struct {
	// The name for a thing
	// example: Some name
	// required: true
	Name string `json:"name"`
	// The value for a thing
	// example: Some value
	// required: true
	Value string `json:"value"`
}

// swagger:parameters create-thing
type _ struct {
	// The body to create a thing
	// in:body
	// required: true
	Body CreateThing
}

// swagger:route POST /thing/new Thing create-thing
//
// This is the summary for creating a thing
//
// This is the description for creating a thing. Which can be longer.
//
// responses:
//   200: ThingResponse
//   404: ErrorResponse
//   500: ErrorResponse
func (s *Server) CreateThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

// swagger:model UpdateThing
type UpdateThing struct {
	// The new value for a thing
	// example: Updated value
	// required: true
	Value string `json:"value" validate:"required"`
}

// swagger:parameters update-thing
type _ struct {
	// The body to update a thing
	// in:body
	// required: true
	Body UpdateThing
}

// swagger:route PUT /thing/{uuid} Thing update-thing
//
// This is the summary for updating a thing
//
// This is the description for updating a thing. Which can be longer.
//
// responses:
//   200: ThingResponse
//   404: ErrorResponse
//   500: ErrorResponse
func (s *Server) UpdateThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

// swagger:route DELETE /thing/{uuid} Thing delete-thing
//
// This is the summary for deleting a thing
//
// This is the description for deleting a thing. Which can be longer.
//
// produces:
// - text/plain
//
// responses:
//   200:
//   500: ErrorResponse
func (s *Server) DeleteThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

// swagger:model ThingsResponse
type ThingsResponse struct {
	Total  int             `json:"total"`
	Page   int             `json:"page"`
	Limit  int             `json:"limit"`
	Things []ThingResponse `json:"things"`
}

// swagger:parameters list-things
type _ struct {
	// Page
	// in:query
	Page int
	// Limit (max 100)
	// in:query
	Limit int
}

// swagger:route GET /thing/ Thing list-things
//
// This is the summary for listing things
//
// This is the description for listing things. Which can be longer.
//
// responses:
//   200: ThingsResponse
//   404: ErrorResponse
//   500: ErrorResponse
func (s *Server) ListThings(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
