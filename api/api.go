package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rrojan/gin-skeleton/api/routes"
	"github.com/rrojan/gin-skeleton/config"
	"github.com/rrojan/gin-skeleton/utils"
)

// Every directory has an "App" type that represents it

// App for API
type App struct {
	Router *routes.Router
	Port   string
}

// NewApp creates a new instance of app
func NewApp() *App {
	r := routes.NewRouter()
	return &App{
		Router: r,
		Port:   fmt.Sprintf(":%d", config.LoadKeyFromConfig("api", "PORT").(int)),
	}
}

func (a App) Run() {
	apiConfig := config.LoadDataFromConfig("api")

	// Set trusted proxies for API
	a.Router.SetTrustedProxies(
		utils.SliceInterfaceToString(apiConfig["TRUSTED_PROXIES"].([]interface{})),
	)
	// Set app mode
	gin.SetMode(apiConfig["MODE"].(string))

	a.Router.Run(a.Port)
}
