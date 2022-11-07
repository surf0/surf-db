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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sh/records/bonus/{mapname}/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "surfheaven"
                ],
                "summary": "get all records for all bonuses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "map name",
                        "name": "mapname",
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
                                "$ref": "#/definitions/models.Record"
                            }
                        }
                    }
                }
            }
        },
        "/sh/records/bonus/{mapname}/{bonus}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "surfheaven"
                ],
                "summary": "get all records for specific bonus",
                "parameters": [
                    {
                        "type": "string",
                        "description": "map name",
                        "name": "mapname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "bonus e.g. 1, 2, ...",
                        "name": "bonus",
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
                                "$ref": "#/definitions/models.Record"
                            }
                        }
                    }
                }
            }
        },
        "/sh/records/map/{mapname}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "surfheaven"
                ],
                "summary": "get all records for map by map name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "map name",
                        "name": "mapname",
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
                                "$ref": "#/definitions/models.Record"
                            }
                        }
                    }
                }
            }
        },
        "/sh/records/stage/{mapname}/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "surfheaven"
                ],
                "summary": "get all records for all stages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "map name",
                        "name": "mapname",
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
                                "$ref": "#/definitions/models.Record"
                            }
                        }
                    }
                }
            }
        },
        "/sh/records/stage/{mapname}/{stage}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "surfheaven"
                ],
                "summary": "get all records for specific stage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "map name",
                        "name": "mapname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "stage e.g. 1, 2, ...",
                        "name": "stage",
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
                                "$ref": "#/definitions/models.Record"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Record": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "improvement": {
                    "type": "string"
                },
                "map_name": {
                    "type": "string"
                },
                "player_id": {
                    "type": "string"
                },
                "player_name": {
                    "type": "string"
                },
                "server": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "track": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "surfheaven.eu CS:GO surf servers",
            "name": "surfheaven"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Surf DB API",
	Description:      "CS:GO & CS:S surf record and map data. Github: https://github.com/surf0/surf-db",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
