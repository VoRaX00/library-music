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
        "/api/add": {
            "post": {
                "description": "Create a new music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "AddMusic",
                "operationId": "create-music",
                "parameters": [
                    {
                        "description": "Music info to add",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MusicToAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/delete": {
            "delete": {
                "description": "delete music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "DeleteMusic",
                "operationId": "delete-music",
                "parameters": [
                    {
                        "description": "Song and group for delete",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MusicToDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/getAll": {
            "get": {
                "description": "get all music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "GetAllMusic",
                "operationId": "get-all-music",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "$ref": "#/definitions/domain.Music"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/getText": {
            "get": {
                "description": "get text music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "GetTextMusic",
                "operationId": "get-text-music",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Song and group for get text",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MusicToGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/info": {
            "get": {
                "description": "get music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "GetMusic",
                "operationId": "get-music",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Song and group for get music",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MusicToGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Music"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/update": {
            "put": {
                "description": "update music",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "music"
                ],
                "summary": "UpdateMusic",
                "operationId": "update-music",
                "parameters": [
                    {
                        "description": "Music info to update",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MusicToUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Music": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "domain.MusicToAdd": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "domain.MusicToDelete": {
            "type": "object",
            "properties": {
                "song": {
                    "type": "string"
                }
            }
        },
        "domain.MusicToGet": {
            "type": "object",
            "properties": {
                "song": {
                    "type": "string"
                }
            }
        },
        "domain.MusicToUpdate": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8090",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Library music API",
	Description:      "API Server for Library music application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}