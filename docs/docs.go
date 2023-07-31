// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/activities": {
            "get": {
                "description": "get activities",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activities"
                ],
                "summary": "Show all activities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/activity.Activity"
                            }
                        }
                    }
                }
            }
        },
        "/activities/id/{id}": {
            "get": {
                "description": "Get activity by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activities"
                ],
                "summary": "Show all activity by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Activity ID",
                        "name": "id",
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
                                "$ref": "#/definitions/activity.Activity"
                            }
                        }
                    }
                }
            }
        },
        "/activities/name/{name}": {
            "get": {
                "description": "Get activities by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activities"
                ],
                "summary": "Show all activities by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Activity Name",
                        "name": "name",
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
                                "$ref": "#/definitions/activity.Activity"
                            }
                        }
                    }
                }
            }
        },
        "/activities/start/session/id/{id}": {
            "post": {
                "description": "Stop Activity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activities"
                ],
                "summary": "Stop activity timer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Activity ID",
                        "name": "id",
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
                                "$ref": "#/definitions/activity.ExerciseSession"
                            }
                        }
                    }
                }
            }
        },
        "/activities/user/exercises/user/{user_id}": {
            "post": {
                "description": "Get exercise session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activities"
                ],
                "summary": "Get user exercise session",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Activity ID",
                        "name": "id",
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
                                "$ref": "#/definitions/activity.ExerciseSession"
                            }
                        }
                    }
                }
            }
        },
        "/activities/user/session/total/stats/{user_id}": {
            "post": {
                "description": "Get user exercise total data for durations and calories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activities"
                ],
                "summary": "Get user exercise data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Activity ID",
                        "name": "id",
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
                                "$ref": "#/definitions/activity.ExerciseSession"
                            }
                        }
                    }
                }
            }
        },
        "/activities/user/session/total/user/{user_id}": {
            "post": {
                "description": "Get user exercise total data for durations and calories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activities"
                ],
                "summary": "Get user exercise data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Activity ID",
                        "name": "id",
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
                                "$ref": "#/definitions/activity.ExerciseSession"
                            }
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get the user's info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.UserSession"
                        }
                    }
                }
            }
        },
        "/users/sign-in": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Sign in a user",
                "parameters": [
                    {
                        "description": "The user's email and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.signInRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.SuccessResponse"
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "contains the session id in bearer format"
                            }
                        }
                    }
                }
            }
        },
        "/users/sign-out": {
            "get": {
                "description": "Delete current user session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "Delete user session",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sessionId string",
                        "name": "sessionId",
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
                                "$ref": "#/definitions/auth.UserSession"
                            }
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
                    "users"
                ],
                "summary": "Sign out a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/users/sign-up": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Sign up a user",
                "parameters": [
                    {
                        "description": "The user's first name, last name, email, and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.userRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.signUpSuccessResponse"
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "contains the session id in bearer format"
                            }
                        }
                    }
                }
            }
        },
        "/users/user/info": {
            "get": {
                "description": "Get info about user session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "Get user session",
                "parameters": [
                    {
                        "type": "string",
                        "description": "session string",
                        "name": "session",
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
                                "$ref": "#/definitions/auth.UserSession"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "activity.Activity": {
            "type": "object",
            "properties": {
                "calories_per_hour": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "duration_minutes": {
                    "type": "number"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "name": {
                    "type": "string"
                },
                "total_calories": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "$ref": "#/definitions/sql.NullString"
                }
            }
        },
        "activity.ExerciseSession": {
            "type": "object",
            "properties": {
                "activity_id": {
                    "type": "integer"
                },
                "calories_burned": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "duration_hours": {
                    "type": "integer"
                },
                "duration_minutes": {
                    "type": "integer"
                },
                "duration_seconds": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": ""
                },
                "session_name": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "auth.UserSession": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "sql.NullString": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if String is not NULL",
                    "type": "boolean"
                }
            }
        },
        "user.SuccessResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "user.signInRequestBody": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "user.signUpSuccessResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "user.userRequestBody": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 48,
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "StayHealthy Swagger Documentation",
	Description:      "Alpha server built with Go and Chi",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
