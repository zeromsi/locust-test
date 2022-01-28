package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/locust-test/block"
)
var Chain *block.BlockChain
//New returns echo object
func New() *echo.Echo {
	InitEnvironmentVariables()
	initChain()
	echoInstance := echo.New()
	// Configuring Middleware Logger
	echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Skipping logging for health checking api
		Skipper: func(c echo.Context) bool {
			if c.Request().RequestURI == "/health" {
				return true
			}
			return false
		},
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	echoInstance.Use(middleware.Recover())
	return echoInstance
}

func initChain(){
	Chain = block.InitBlockChain()
}