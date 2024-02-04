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
        "license": {
            "name": "MIT",
            "url": "https://github.com/simple-retro/api/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "Show API health",
                "responses": {
                    "200": {
                        "description": "API metrics",
                        "schema": {
                            "$ref": "#/definitions/server.health"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/retrospective": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Create Retrospective",
                "parameters": [
                    {
                        "description": "Create Retrospective",
                        "name": "retrospective",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RetrospectiveCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/retrospective/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Get Retrospective by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Retrospective ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Delete Retrospective by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Retrospective ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Retrospective"
                ],
                "summary": "Update Retrospective by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Retrospective ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Retrospective",
                        "name": "retrospective",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RetrospectiveCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retrospective Object",
                        "schema": {
                            "$ref": "#/definitions/types.Retrospective"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.health": {
            "type": "object",
            "properties": {
                "cpu": {
                    "type": "number"
                },
                "memory": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.Answer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "types.Question": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Answer"
                    }
                },
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "types.Retrospective": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Question"
                    }
                }
            }
        },
        "types.RetrospectiveCreateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}