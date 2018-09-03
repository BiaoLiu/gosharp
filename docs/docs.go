// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-09-03 17:22:02.2632678 +0800 CST m=+0.074803301

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "gosharp API",
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
        "version": "1.0"
    },
    "host": "localhost:8200",
    "basePath": "/",
    "paths": {
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录注册模块"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/serializers.UserResponse"
                        }
                    }
                }
            }
        },
        "/test": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试模块"
                ],
                "summary": "测试",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializers.UserResponse": {
            "type": "object",
            "properties": {
                "created_time": {
                    "type": "string",
                    "example": "创建时间"
                },
                "email": {
                    "type": "string",
                    "example": "邮箱"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "mobile": {
                    "type": "string",
                    "example": "手机号码"
                },
                "username": {
                    "type": "string",
                    "example": "用户名"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
