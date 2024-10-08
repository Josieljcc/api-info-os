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
        "/clients": {
            "get": {
                "description": "Get all clients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Get all clients",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Client"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Create a new client",
                "parameters": [
                    {
                        "description": "Client",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ClientResponse"
                        }
                    }
                }
            }
        },
        "/clients/{id}": {
            "get": {
                "description": "Get a client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Get a client",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.Client"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Update a client",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Client",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ClientResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Delete a client",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ClientResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/orders": {
            "get": {
                "description": "Get Orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Orders",
                "responses": {}
            },
            "post": {
                "description": "Create Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Order"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Get Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Order"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Delete Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/parts": {
            "get": {
                "description": "Get Parts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Part"
                ],
                "summary": "Get Parts",
                "responses": {}
            },
            "post": {
                "description": "Create Part",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Part"
                ],
                "summary": "Create Part",
                "parameters": [
                    {
                        "description": "Part",
                        "name": "part",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Part"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/parts/{id}": {
            "get": {
                "description": "Get Part",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Part"
                ],
                "summary": "Get Part",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Part",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Part"
                ],
                "summary": "Update Part",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Part",
                        "name": "part",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Part"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Part",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Part"
                ],
                "summary": "Delete Part",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/services": {
            "get": {
                "description": "Get Services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Get Services",
                "responses": {}
            },
            "post": {
                "description": "Create Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Create Service",
                "parameters": [
                    {
                        "description": "Service",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Service"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/services/{id}": {
            "get": {
                "description": "Get Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Get Service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Update Service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Service"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Delete Service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/technicians": {
            "get": {
                "description": "Get Technicians",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technician"
                ],
                "summary": "Get Technicians",
                "responses": {}
            },
            "post": {
                "description": "Create Technician",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technician"
                ],
                "summary": "Create Technician",
                "parameters": [
                    {
                        "description": "Technician",
                        "name": "technician",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Technician"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/technicians/{id}": {
            "get": {
                "description": "Get Technician",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technician"
                ],
                "summary": "Get Technician",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Technician",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technician"
                ],
                "summary": "Update Technician",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Technician",
                        "name": "technician",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Technician"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Technician",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technician"
                ],
                "summary": "Delete Technician",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "schemas.Client": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Order"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "schemas.ClientResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "schemas.Order": {
            "type": "object",
            "properties": {
                "clientID": {
                    "type": "string"
                },
                "comment": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "parts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Part"
                    }
                },
                "services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Service"
                    }
                },
                "status": {
                    "type": "string"
                },
                "technicianID": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "schemas.Part": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Order"
                    }
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "schemas.Service": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Order"
                    }
                },
                "price": {
                    "type": "number"
                },
                "time": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "schemas.Technician": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Order"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updatedAt": {
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
	Title:            "API de Controle de Ordens de Serviço",
	Description:      "Esta é a API para controle de ordens de serviço em uma assistência de manutenção de computadores.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
