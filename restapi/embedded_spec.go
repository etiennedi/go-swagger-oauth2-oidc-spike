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
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "spikeyfoo",
    "version": "1.0.0"
  },
  "paths": {
    "/foo": {
      "get": {
        "tags": [
          "foo"
        ],
        "responses": {
          "200": {
            "description": "some-endpoint",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "basic": {
      "description": "don't ever use this",
      "type": "basic"
    },
    "oidc": {
      "description": "OIDC (OpenConnect ID - based on OAuth2)",
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://foo"
    }
  },
  "security": [
    {},
    {
      "basic": null
    },
    {
      "oidc": [
        "foohoo"
      ]
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "spikeyfoo",
    "version": "1.0.0"
  },
  "paths": {
    "/foo": {
      "get": {
        "tags": [
          "foo"
        ],
        "responses": {
          "200": {
            "description": "some-endpoint",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "basic": {
      "description": "don't ever use this",
      "type": "basic"
    },
    "oidc": {
      "description": "OIDC (OpenConnect ID - based on OAuth2)",
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://foo"
    }
  },
  "security": [
    {},
    {
      "basic": []
    },
    {
      "oidc": [
        "foohoo"
      ]
    }
  ]
}`))
}
