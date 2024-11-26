// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/song": {
            "post": {
                "description": "Submit a new song's details and store it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Post a new song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group name",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song name",
                        "name": "song",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The posted song details",
                        "schema": {
                            "$ref": "#/definitions/models.FullSong"
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters"
                    },
                    "500": {
                        "description": "Request failed"
                    }
                }
            }
        },
        "/song/": {
            "get": {
                "description": "Retrieve the text for a specific song by its group and song name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Get song text",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group name",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song name",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Verse number (default 1)",
                        "name": "verse",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Text of the song",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters"
                    },
                    "500": {
                        "description": "Internal error"
                    }
                }
            },
            "delete": {
                "description": "Delete the song specified by group and song name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Delete a song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group name",
                        "name": "group",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Song name",
                        "name": "song",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song successfully deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters"
                    },
                    "500": {
                        "description": "Internal error"
                    }
                }
            },
            "patch": {
                "description": "Update the details of a song, identified by group and song name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Update an existing song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group name",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song name",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "description": "Updated song details",
                        "name": "songDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FullSong"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The updated song details",
                        "schema": {
                            "$ref": "#/definitions/models.FullSong"
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters"
                    },
                    "500": {
                        "description": "Internal error"
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Retrieve a list of songs with optional limit, page, and filter parameters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Get all songs with optional filters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Limit the number of results",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Specify the page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter songs by field (song, group, release)",
                        "name": "filter",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of songs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.FullSong"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.FullSong": {
            "description": "FullSong represents a song in some endpoints",
            "type": "object",
            "properties": {
                "group": {
                    "description": "The name of the group",
                    "type": "string"
                },
                "link": {
                    "description": "The link to song",
                    "type": "string"
                },
                "releaseDate": {
                    "description": "The release date of song",
                    "type": "string"
                },
                "song": {
                    "description": "The name of the song",
                    "type": "string"
                },
                "text": {
                    "description": "The text of song",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Music Lib",
	Description:      "Sever for manage music lib.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
