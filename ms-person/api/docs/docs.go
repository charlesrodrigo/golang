// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Charles Rodrigo",
            "email": "charlesrodrigo@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/person": {
            "get": {
                "description": "get all person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "get all person",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.GetPersonRequest"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "create person",
                "parameters": [
                    {
                        "description": "Person Data",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreatePersonRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/person/{id}": {
            "get": {
                "description": "get person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "get person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id person",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetPersonRequest"
                        }
                    }
                }
            },
            "put": {
                "description": "update person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "update person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id person",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Person Data",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreatePersonRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreatePersonRequest"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "delete person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id person",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Address": {
            "type": "object",
            "required": [
                "city",
                "country",
                "neighborhood",
                "state",
                "street",
                "zipcode"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        },
        "controllers.CreatePersonRequest": {
            "type": "object",
            "required": [
                "address",
                "email",
                "name"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/controllers.Address"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.GetPersonRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/controllers.Address"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "http://localhost:8080/swagger/index.html"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Crud Person API",
	Description:      "This is a crud of person.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}