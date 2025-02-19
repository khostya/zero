// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/edit/{id}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Post News",
                        "name": "news",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.PostNews"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.PostNews"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint32",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "uint32",
                        "description": "size",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.ListNews"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.News": {
            "type": "object",
            "properties": {
                "Content": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "http.ListNews": {
            "type": "object",
            "properties": {
                "News": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.News"
                    }
                },
                "Success": {
                    "type": "boolean"
                }
            }
        },
        "http.PostNews": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "content": {
                    "type": "string"
                },
                "title": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "zero",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
