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
        "/v1/health/": {
            "get": {
                "description": "Test Health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Test",
                "responses": {
                    "200": {
                        "description": "Done",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/test/binder/body": {
            "post": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Test BodyBinder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "BodyBinder",
                "parameters": [
                    {
                        "description": "person data",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.personData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Done",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/test/user/{id}": {
            "get": {
                "description": "Test UserById",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "UserById",
                "parameters": [
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
                        "description": "Done",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/send-otp": {
            "post": {
                "description": "Test SendOtp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "SendOtp to user",
                "parameters": [
                    {
                        "description": "GetOtpRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GetOtpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Done",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.GetOtpRequest": {
            "type": "object",
            "required": [
                "mobileNumber"
            ],
            "properties": {
                "mobileNumber": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "handlers.personData": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "mobile_number"
            ],
            "properties": {
                "first_name": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 4
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 6
                },
                "mobile_number": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "resultCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                },
                "validationErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validation.ValidationError"
                    }
                }
            }
        },
        "validation.ValidationError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "property": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthBearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
