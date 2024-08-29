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
        "/bookings": {
            "post": {
                "description": "Create a new booking for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Create a new booking",
                "parameters": [
                    {
                        "description": "Booking information",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateBookingReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateBookingResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/bookings/user/{user_id}": {
            "get": {
                "description": "Retrieve all bookings for a specific user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get bookings by user ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetBookingByUserIdResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/bookings/{booking_id}": {
            "get": {
                "description": "Retrieve a specific booking by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get a booking by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking ID",
                        "name": "booking_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetBookingResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing booking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Update a booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking ID",
                        "name": "booking_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated booking information",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateBookingReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateBookingResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing booking",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Delete a booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking ID",
                        "name": "booking_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.DeleteBookingResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/hotels": {
            "get": {
                "description": "Retrieve a list of all hotels",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Get all hotels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetAllHotelsST"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new hotel with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Create a new hotel",
                "parameters": [
                    {
                        "description": "Hotel information",
                        "name": "hotel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateHotelReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateHotelResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/hotels/{hotel_id}": {
            "get": {
                "description": "Retrieve a specific hotel by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Get a hotel by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel ID",
                        "name": "hotel_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetHotelByIdResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing hotel's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Update an existing hotel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel ID",
                        "name": "hotel_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated hotel information",
                        "name": "hotel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateHotelReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateHotelResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a hotel by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Delete a hotel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel ID",
                        "name": "hotel_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.DeleteHotelResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/hotels/{hotel_id}/rooms/availability": {
            "get": {
                "description": "Retrieve the availability of rooms for a specific hotel",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Get hotel rooms availability",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel ID",
                        "name": "hotel_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.HotelRoomsAvailabilityResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Authenticate a user and return a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Create a new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateUserResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/users/verify": {
            "post": {
                "description": "Verify a user's account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Verify user",
                "parameters": [
                    {
                        "description": "User verification data",
                        "name": "verificationData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.VerifyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.VerifyResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/users/{user_id}": {
            "get": {
                "description": "Retrieve a user's information by their ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetUserByIdResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing user's information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateUserResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.DeleteUserResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.Booking": {
            "type": "object",
            "properties": {
                "booking_id": {
                    "type": "string"
                },
                "check_in_date": {
                    "type": "string"
                },
                "check_out_date": {
                    "type": "string"
                },
                "hotel_id": {
                    "type": "string"
                },
                "room_type": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.CUD": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "handlers.CUDUser": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "handlers.CreateBookingReq": {
            "type": "object",
            "properties": {
                "check_in_date": {
                    "type": "string"
                },
                "check_out_date": {
                    "type": "string"
                },
                "hotel_id": {
                    "type": "string"
                },
                "room_type": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.CreateBookingResp": {
            "type": "object",
            "properties": {
                "check_in_date": {
                    "type": "string"
                },
                "check_out_date": {
                    "type": "string"
                },
                "hotel_id": {
                    "type": "string"
                },
                "room_type": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.CreateHotelReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "handlers.CreateHotelResp": {
            "type": "object",
            "properties": {
                "hotel": {
                    "$ref": "#/definitions/handlers.Hotel"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.CreateUserReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.CreateUserResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.DeleteBookingResp": {
            "type": "object",
            "properties": {
                "booking_id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.DeleteHotelResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.DeleteUserResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.GetAllHotelsST": {
            "type": "object",
            "properties": {
                "hotels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.Hotel"
                    }
                },
                "hotels_count": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.GetBookingByUserIdResp": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.Booking"
                    }
                }
            }
        },
        "handlers.GetBookingResp": {
            "type": "object",
            "properties": {
                "booking_id": {
                    "type": "string"
                },
                "check_in_date": {
                    "type": "string"
                },
                "check_out_date": {
                    "type": "string"
                },
                "hotel_id": {
                    "type": "string"
                },
                "room_type": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.GetHotelByIdResp": {
            "type": "object",
            "properties": {
                "hotel": {
                    "$ref": "#/definitions/handlers.Hotel"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.GetUserByIdResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "user": {
                    "$ref": "#/definitions/handlers.User"
                }
            }
        },
        "handlers.Hotel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "cud": {
                    "$ref": "#/definitions/handlers.CUD"
                },
                "hotel_id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "room": {
                    "$ref": "#/definitions/handlers.Room"
                },
                "rooms_count": {
                    "type": "integer"
                }
            }
        },
        "handlers.HotelRoomsAvailabilityResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.Room"
                    }
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.LoginResp": {
            "type": "object",
            "properties": {
                "expires_in": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "handlers.Room": {
            "type": "object",
            "properties": {
                "availability": {
                    "type": "boolean"
                },
                "hotel_id": {
                    "type": "string"
                },
                "price_per_night": {
                    "type": "number"
                },
                "room_id": {
                    "type": "string"
                },
                "room_type": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateBookingReq": {
            "type": "object",
            "properties": {
                "booking_id": {
                    "type": "string"
                },
                "check_in_date": {
                    "type": "string"
                },
                "check_out_date": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                }
            }
        },
        "handlers.UpdateBookingResp": {
            "type": "object",
            "properties": {
                "booking_id": {
                    "type": "string"
                },
                "check_in_date": {
                    "type": "string"
                },
                "check_out_date": {
                    "type": "string"
                },
                "hotel_id": {
                    "type": "string"
                },
                "room_type": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateHotelReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "hotel_id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "handlers.UpdateHotelResp": {
            "type": "object",
            "properties": {
                "hotel": {
                    "$ref": "#/definitions/handlers.Hotel"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "handlers.UpdateUserReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateUserResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "user": {
                    "$ref": "#/definitions/handlers.User"
                }
            }
        },
        "handlers.User": {
            "type": "object",
            "properties": {
                "cud": {
                    "$ref": "#/definitions/handlers.CUDUser"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.VerifyReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handlers.VerifyResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authourization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "",
	Description:      "Api-gateway service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
