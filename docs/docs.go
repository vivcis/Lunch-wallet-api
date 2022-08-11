// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "name": "Lunch-wallet Team API Support",
            "url": "http://www.swagger.io/support",
            "email": "info@lunchwallet.com"
        },
        "license": {
            "name": "BSD",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/blockedusers": {
            "get": {
                "description": "Admin gets to see all blocked users with this endpoint. It is an authorized route to only ADMIN",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "This endpoint enables admin to see all blocked users",
                "responses": {
                    "200": {
                        "description": "blocked users successfully gotten",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.FoodBeneficiary"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/createtimetable": {
            "post": {
                "description": "creates meal by collecting fields in models.Food in a form data. Note: \"images\" is a file to be uploaded in jpeg or png. \"name\" is the name of the meal, \"type\" is either brunch or dinner, \"weekday\" can be ignored but it is either monday - sunday, \"kitchen\" is either uno, edo-tech park, etc. \"year\", \"month\" and \"day\" are numbers. It is an authorized route to only ADMIN",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Food"
                ],
                "summary": "Admin creates meal",
                "parameters": [
                    {
                        "description": "images, type, name, kitchen, year, month, day",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Food"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/numberblocked": {
            "get": {
                "description": "Admin gets to see how manuy beneficiaries blocked. It is an authorized route to only ADMIN",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Gets number of blocked benefeciary",
                "responses": {
                    "200": {
                        "description": "successfully gotten",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/benefactor/allfood": {
            "get": {
                "description": "This should be used to get all the food in the database meant for today. This should be used instead of GetBrunch and GetDinner seperately for scalability purpose when rendering on the Beneficiary dashboard. Frontend can seperate dinner and brunch. It is an authorized route to only foodBeneficiary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Food"
                ],
                "summary": "Gets all the food in the database using the date of the present day",
                "responses": {
                    "200": {
                        "description": "Food successfully gotten",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Food"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/benefactor/brunch": {
            "get": {
                "description": "Gets all the brunch in the database meant for today. GetAllFood should be used instead for scalability purpose when rendering on the Beneficiary dashboard. Frontend can filter brunch and dinner It is an authorized route to only foodBeneficiary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Food"
                ],
                "summary": "Gets all the brunch in the database using the date of the present day",
                "responses": {
                    "200": {
                        "description": "Brunch found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Food"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/benefactor/dinner": {
            "get": {
                "description": "Gets all the dinner in the database meant for today. GetAllFood should be used instead for scalability purpose when rendering on the Beneficiary dashboard. Frontend can filter brunch and dinner It is an authorized route to only foodBeneficiary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Food"
                ],
                "summary": "Gets all the dinner in the database using the date of the present day",
                "responses": {
                    "200": {
                        "description": "Dinner found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Food"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/beneficiarylogout": {
            "post": {
                "description": "Log out a kitchen staff",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Logout User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/adminlogin": {
            "post": {
                "description": "Allows Admin to log in in order to use app dashboard. Admin must be active before he or she can log in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login Admin",
                "parameters": [
                    {
                        "description": "email, password",
                        "name": "kitchenStaff",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "login successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/benefactorlogin": {
            "post": {
                "description": "Allows Food Beneficiary to log in in order to use app dashboard. Beneficiary must be active before he or she can log in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login Food Beneficiary",
                "parameters": [
                    {
                        "description": "email, password",
                        "name": "kitchenStaff",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "login successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/beneficiarysignup": {
            "post": {
                "description": "creates a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Add user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FoodBeneficiary"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/beneficiaryverifyemail/{token}": {
            "patch": {
                "description": "verifies a food beneficiary email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Verify Email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token string",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/kitchenstafflogin": {
            "post": {
                "description": "Allows Kitchen staff to log in in order to use app dashboard. Staff must be active before he or she can log in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login Kitchen Staff",
                "parameters": [
                    {
                        "description": "email, password",
                        "name": "kitchenStaff",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "login successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/kitchenstaffsignup": {
            "post": {
                "description": "creates a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Add user",
                        "name": "staff",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.KitchenStaff"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/kitchenstaffverifyemail/{token}": {
            "patch": {
                "description": "verifies a kitchen staff email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Verify Email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token string",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/notifications": {
            "get": {
                "description": "Returns all notifications in the database and their dates to be rendered as will by the frontend",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Notifies users whenever there is a change worthy of notification",
                "responses": {
                    "200": {
                        "description": "notification successfully loaded",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Notification"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Food": {
            "type": "object",
            "properties": {
                "adminName": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "day": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Image"
                    }
                },
                "kitchen": {
                    "type": "string"
                },
                "month": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "weekday": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.FoodBeneficiary": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "location",
                "stack"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_block": {
                    "type": "boolean"
                },
                "location": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "stack": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Image": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "product_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.KitchenStaff": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "location"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_block": {
                    "type": "boolean"
                },
                "location": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Notification": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "month": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.UserLogin": {
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
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Lunch Wallet Swagger API",
	Description:      "This is a lunch wallet server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
