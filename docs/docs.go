// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/index": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "首页"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/createUser": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "repassword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\", \"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\", \"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "所有用户",
                "responses": {
                    "200": {
                        "description": "code\", \"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户 ID",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\", \"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
