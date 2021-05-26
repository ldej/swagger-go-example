// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Laurence de Jong",
            "url": "https://ldej.nl/",
            "email": "support@ldej.nl"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/thing": {
            "get": {
                "description": "This is the description for listing things. Which can be longer.",
                "tags": [
                    "Thing"
                ],
                "summary": "This is the summary for listing things",
                "operationId": "list-things",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit (max 100)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ThingsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/thing/new": {
            "post": {
                "description": "This is the description for creating a thing. Which can be longer.",
                "tags": [
                    "Thing"
                ],
                "summary": "This is the summary for creating a thing",
                "operationId": "create-thing",
                "parameters": [
                    {
                        "description": "The body to create a thing",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateThing"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ThingResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/thing/{uuid}": {
            "get": {
                "description": "This is the description for getting a thing by its UUID. Which can be longer.",
                "tags": [
                    "Thing"
                ],
                "summary": "This is the summary for getting a thing by its UUID",
                "operationId": "get-thing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The UUID of a thing",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ThingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "This is the description for updating a thing. Which can be longer.",
                "tags": [
                    "Thing"
                ],
                "summary": "This is the summary for updating a thing",
                "operationId": "update-thing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The UUID of a thing",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The body to update a thing",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateThing"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ThingResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "This is the description for deleting a thing. Which can be longer.",
                "tags": [
                    "Thing"
                ],
                "summary": "This is the summary for deleting a thing",
                "operationId": "delete-thing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The UUID of a thing",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateThing": {
            "type": "object",
            "required": [
                "name",
                "value"
            ],
            "properties": {
                "name": {
                    "description": "The name for a thing",
                    "type": "string",
                    "example": "Some name"
                },
                "value": {
                    "description": "The value for a thing",
                    "type": "string",
                    "example": "Some value"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "An error occurred"
                }
            }
        },
        "ThingResponse": {
            "type": "object",
            "properties": {
                "created": {
                    "description": "The time a thing was created",
                    "type": "string",
                    "format": "date-time",
                    "example": "2021-05-25T00:53:16.535668Z"
                },
                "name": {
                    "description": "The Name of a thing",
                    "type": "string",
                    "example": "Some name"
                },
                "updated": {
                    "description": "The last time a thing was updated",
                    "type": "string",
                    "format": "date-time",
                    "example": "2021-05-25T00:53:16.535668Z"
                },
                "uuid": {
                    "description": "The UUID of a thing",
                    "type": "string",
                    "example": "6204037c-30e6-408b-8aaa-dd8219860b4b"
                },
                "value": {
                    "description": "The Value of a thing",
                    "type": "string",
                    "example": "Some value"
                }
            }
        },
        "ThingsResponse": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "format": "int64"
                },
                "page": {
                    "type": "integer",
                    "format": "int64"
                },
                "things": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ThingResponse"
                    }
                },
                "total": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "UpdateThing": {
            "type": "object",
            "required": [
                "value"
            ],
            "properties": {
                "value": {
                    "description": "The new value for a thing",
                    "type": "string",
                    "example": "Update value"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "A thing server",
	Description: "A thing server",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
