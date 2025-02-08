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
        "/clients": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Создаёт нового клиента",
                "parameters": [
                    {
                        "description": "Payload для создания клиента",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http_handler.CreateClientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID созданного клиента",
                        "schema": {
                            "$ref": "#/definitions/http_handler.CreateClientResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/clients/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Обновляет данные клиента",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID клиента",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload для обновления клиента",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http_handler.UpdateClientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Удаляет клиента по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID клиента",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/clients/{photographerID}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Возвращает список клиентов фотографа",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID фотографа",
                        "name": "photographerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Client"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/debt": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Financial"
                ],
                "summary": "Добавляет задолженность для клиента фотографа",
                "parameters": [
                    {
                        "description": "Payload для добавления задолженности",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http_handler.AddDebtRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/debtors/{photographerID}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Financial"
                ],
                "summary": "Получает список должников фотографа",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID фотографа",
                        "name": "photographerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список задолженностей",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Debt"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/incomes/{photographerID}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Financial"
                ],
                "summary": "Получает детализированный список доходов фотографа",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID фотографа",
                        "name": "photographerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список платежей и общий доход",
                        "schema": {
                            "$ref": "#/definitions/http_handler.GetIncomesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/payment": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Financial"
                ],
                "summary": "Добавляет оплату клиента фотографу",
                "parameters": [
                    {
                        "description": "Payload для добавления оплаты",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http_handler.AddPaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/photographers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photographers"
                ],
                "summary": "Возвращает список фотографов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Photographer"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photographers"
                ],
                "summary": "Создаёт нового фотографа",
                "parameters": [
                    {
                        "description": "Payload для создания фотографа",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http_handler.CreatePhotographerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID созданного фотографа",
                        "schema": {
                            "$ref": "#/definitions/http_handler.CreatePhotographerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Client": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photographer_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Debt": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "client_id": {
                    "type": "integer"
                },
                "client_name": {
                    "type": "string"
                },
                "occurredAt": {
                    "type": "string"
                }
            }
        },
        "domain.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "client_id": {
                    "type": "integer"
                },
                "occurredAt": {
                    "type": "string"
                }
            }
        },
        "domain.Photographer": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "http_handler.AddDebtRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 500
                },
                "client_id": {
                    "type": "integer",
                    "example": 2
                },
                "photographer_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "http_handler.AddPaymentRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 500
                },
                "client_id": {
                    "type": "integer",
                    "example": 2
                },
                "photographer_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "http_handler.CreateClientRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Alice"
                },
                "photographer_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "http_handler.CreateClientResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "http_handler.CreatePhotographerRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Alice"
                }
            }
        },
        "http_handler.CreatePhotographerResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "http_handler.GetIncomesResponse": {
            "type": "object",
            "properties": {
                "payments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Payment"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 10000
                }
            }
        },
        "http_handler.UpdateClientRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Alice Updated"
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
