// Package http API.
// @title webgocast
// @version 1.0
// @description webg论坛
// @termsOfService https://github.com/swaggo/swag

// @contact.name dailingling
// @contact.email dailingling

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
package http

import (
	_ "github.com/dll02/webgo/app/http/swagger"
)
