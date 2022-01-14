// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/example-api/v1/{tenant}/users": {
            "get": {
                "description": "list user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "list user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "current user id",
                        "name": "current_user_id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "tenant id",
                        "name": "tenant",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maximum": 500,
                        "minimum": 1,
                        "type": "integer",
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "page number",
                        "name": "current",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "username ilike",
                        "name": "username__ilike",
                        "in": "query"
                    },
                    {
                        "maxLength": 100,
                        "minLength": 1,
                        "type": "string",
                        "description": "username ilike",
                        "name": "age",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.ListUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/httpx.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/httpx.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Single user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tenant id",
                        "name": "tenant",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.UserInfoResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/httpx.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/httpx.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/example-api/v1/{tenant}/users/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Single user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tenant id",
                        "name": "tenant",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.UserInfoResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/httpx.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/httpx.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ListUserResponse": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer",
                    "example": 3
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.UserInfoDTO"
                    }
                },
                "page_size": {
                    "type": "integer",
                    "example": 20
                },
                "total": {
                    "type": "integer",
                    "example": 120
                },
                "total_page": {
                    "type": "integer",
                    "example": 6
                }
            }
        },
        "dto.UserInfoDTO": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 18
                },
                "email": {
                    "type": "string",
                    "example": "daniel@trinity.com"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ],
                    "example": "male"
                },
                "id": {
                    "type": "string",
                    "example": "1479429646645936128"
                },
                "username": {
                    "type": "string",
                    "example": "Daniel"
                }
            }
        },
        "dto.UserInfoResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 18
                },
                "email": {
                    "type": "string",
                    "example": "daniel@trinity.com"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ],
                    "example": "male"
                },
                "id": {
                    "type": "string",
                    "example": "1479429646645936128"
                },
                "username": {
                    "type": "string",
                    "example": "Daniel"
                }
            }
        },
        "httpx.ErrorInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400001
                },
                "details": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "error detail1",
                        "error detail2"
                    ]
                },
                "message": {
                    "type": "string",
                    "example": "ErrInvalidRequest"
                }
            }
        },
        "httpx.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/httpx.ErrorInfo"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                },
                "trace_id": {
                    "type": "string",
                    "example": "1-trace-it"
                }
            }
        },
        "httpx.SuccessResponse": {
            "type": "object",
            "properties": {
                "result": {},
                "status": {
                    "type": "integer",
                    "example": 200
                },
                "trace_id": {
                    "type": "string",
                    "example": "1-trace-it"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
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
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "trinity-micro Example API",
	Description: "This is a sample server for trinity-micro",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}