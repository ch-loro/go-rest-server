// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Golang Server API",
    "title": "sample",
    "license": {
      "name": "MIT"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8000",
  "basePath": "/v1",
  "paths": {
    "/gettest": {
      "get": {
        "tags": [
          "gettest"
        ],
        "summary": "for Get API test",
        "operationId": "gettest",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "any number",
            "name": "num",
            "in": "query"
          },
          {
            "type": "string",
            "description": "any string",
            "name": "str",
            "in": "query"
          },
          {
            "enum": [
              "foo",
              "bar"
            ],
            "type": "string",
            "description": "any enum",
            "name": "enum",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Obj"
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/posttest": {
      "get": {
        "tags": [
          "posttest"
        ],
        "summary": "for Post API test",
        "operationId": "posttest",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "any number",
            "name": "num",
            "in": "query"
          },
          {
            "type": "string",
            "description": "any string",
            "name": "str",
            "in": "query"
          },
          {
            "enum": [
              "foo",
              "bar"
            ],
            "type": "string",
            "description": "any enum",
            "name": "enum",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "none"
          },
          "201": {
            "description": "over"
          },
          "400": {
            "description": "Unavailable values"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Obj": {
      "type": "object",
      "required": [
        "id",
        "name",
        "value"
      ],
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "enum": [
            "foo",
            "bar"
          ]
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Golang Server API",
    "title": "sample",
    "license": {
      "name": "MIT"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8000",
  "basePath": "/v1",
  "paths": {
    "/gettest": {
      "get": {
        "tags": [
          "gettest"
        ],
        "summary": "for Get API test",
        "operationId": "gettest",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "any number",
            "name": "num",
            "in": "query"
          },
          {
            "type": "string",
            "description": "any string",
            "name": "str",
            "in": "query"
          },
          {
            "enum": [
              "foo",
              "bar"
            ],
            "type": "string",
            "description": "any enum",
            "name": "enum",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Obj"
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/posttest": {
      "get": {
        "tags": [
          "posttest"
        ],
        "summary": "for Post API test",
        "operationId": "posttest",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "any number",
            "name": "num",
            "in": "query"
          },
          {
            "type": "string",
            "description": "any string",
            "name": "str",
            "in": "query"
          },
          {
            "enum": [
              "foo",
              "bar"
            ],
            "type": "string",
            "description": "any enum",
            "name": "enum",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "none"
          },
          "201": {
            "description": "over"
          },
          "400": {
            "description": "Unavailable values"
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Obj": {
      "type": "object",
      "required": [
        "id",
        "name",
        "value"
      ],
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "enum": [
            "foo",
            "bar"
          ]
        }
      }
    }
  }
}`))
}
