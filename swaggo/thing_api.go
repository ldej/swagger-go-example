package main

import (
	"net/http"
	"time"
)

type ThingResponse struct {
	// The UUID of a thing
	UUID string `json:"uuid" example:"6204037c-30e6-408b-8aaa-dd8219860b4b"`
	// The Name of a thing
	Name string `json:"name" example:"Some name"`
	// The Value of a thing
	Value string `json:"value" example:"Some value"`
	// The last time a thing was updated
	Updated time.Time `json:"updated" example:"2021-05-25T00:53:16.535668Z" format:"date-time"`
	// The time a thing was created
	Created time.Time `json:"created" example:"2021-05-25T00:53:16.535668Z" format:"date-time"`
} // @name ThingResponse

type ErrorResponse struct {
	// The error message
	Error string `json:"error" example:"An error occurred"`
} // @name ErrorResponse

// GetThing godoc
// @Summary This is the summary for getting a thing by its UUID
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Description and can continue over multiple lines
// @ID get-thing
// @Tags Thing
// @Param uuid path string true "The UUID of a thing"
// @Success 200 {object} ThingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /thing/{uuid} [get]
func (s *Server) GetThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

type CreateThing struct {
	// The name for a thing
	Name string `json:"name" validate:"required" example:"Some name"`
	// The value for a thing
	Value string `json:"value" validate:"required" example:"Some value"`
} // @name CreateThing

// CreateThing godoc
// @Summary This is the summary for creating a thing
// @Description This is the description for creating a thing. Which can be longer.
// @ID create-thing
// @Tags Thing
// @Param Body body CreateThing true "The body to create a thing"
// @Success 200 {object} ThingResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /thing/new [post]
func (s *Server) CreateThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

type UpdateThing struct {
	// The new value for a thing
	Value string `json:"value" validate:"required" example:"Update value"`
} // @name UpdateThing

// UpdateThing godoc
// @Summary This is the summary for updating a thing
// @Description This is the description for updating a thing. Which can be longer.
// @ID update-thing
// @Tags Thing
// @Param uuid path string true "The UUID of a thing"
// @Param Body body UpdateThing true "The body to update a thing"
// @Success 200 {object} ThingResponse
// @Failure 404,500 {object} ErrorResponse
// @Router /thing/{uuid} [put]
func (s *Server) UpdateThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

// DeleteThing godoc
// @Summary This is the summary for deleting a thing
// @Description This is the description for deleting a thing. Which can be longer.
// @ID delete-thing
// @Tags Thing
// @Param uuid path string true "The UUID of a thing"
// @Success 200
// @Failure 500 {object} ErrorResponse
// @Router /thing/{uuid} [delete]
func (s *Server) DeleteThing(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}

type ThingsResponse struct {
	Total  int             `json:"total" format:"int64"`
	Page   int             `json:"page" format:"int64"`
	Limit  int             `json:"limit" format:"int64"`
	Things []ThingResponse `json:"things"`
} // @name ThingsResponse

// ListThings godoc
// @Summary This is the summary for listing things
// @Description This is the description for listing things. Which can be longer.
// @ID list-things
// @Tags Thing
// @Param page query int false "Page"
// @Param limit query int false "Limit (max 100)"
// @Success 200 {object} ThingsResponse
// @Failure 500 {object} ErrorResponse
// @Router /thing [get]
func (s *Server) ListThings(w http.ResponseWriter, r *http.Request) {
	// Your implementation here
}
