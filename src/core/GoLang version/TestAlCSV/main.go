// ------------------------------------
// RR IT 2024
//
// ------------------------------------

package main

import (

	// Internal project packages
	"rr/TestAlCSV/config"
	"rr/TestAlCSV/middleware"
	"rr/TestAlCSV/routes"

	// Third-party libraries
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	// System Packages
	"io"
	"net/http"
	"os"
)

//
// ----------------------------------------------------------------------------------
//
// 											MAIN
//
// ----------------------------------------------------------------------------------
//

func main() {

	// The LOG FILE, if we are not debugging
	if !config.CONFIG_IS_DEBUG {
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()

		// Logging to a file.
		f, _ := os.Create("gin_server.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// Creating a router for processing requests
	r := gin.Default()

	// Distribution of static for the debug version
	if config.CONFIG_IS_DEBUG_SERVERLESS {
		r.Static("/assets", "./assets") // For static in debugging mode
		// Uploading HTML
		r.LoadHTMLGlob("assets/html/*")
	} else {
		// Uploading HTML
		r.LoadHTMLGlob("static/assets/html/*")
	}

	// For sessions
	store := memstore.NewStore([]byte(config.CONFIG_SECRET))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 дней (Samara don't get out of the TV)
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	r.Use(sessions.Sessions("data", store))

	// CORS
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimiterMiddleware(rate.Limit(10), 5))

	//
	// 	   --------- Paths ---------
	// 	For the implementation of paths, see routes.go
	//

	//
	// Common paths
	//

	// r.GET("/", routes.Handler_Index)
	r.GET("/login", routes.Handler_Login)
	r.POST("/login", routes.Handler_Login)
	r.POST("/login/", routes.Handler_Login)
	r.GET("/logout", routes.Handler_Logout)
	r.GET("/lord-panel", routes.Handler_LordPanel)

	// API Group
	api := r.Group("/api")
	{
		//Dataset
		dataset := api.Group("/dataset")
		{
			dataset.GET("/", routes.Handler_API_Dataset_GetData)
			dataset.POST("/", routes.Handler_API_Dataset_AddData)
			dataset.DELETE("/", routes.Handler_API_Dataset_DeleteData)
		}
	}

	if config.CONFIG_IS_DEBUG_SERVERLESS {
		// Starting the server
		r.Run(":" + config.CONFIG_DEBUG_SERVERLESS_SERVER_PORT) // listen and serve on 0.0.0.0:PORT
	} else {
		// Starting the server
		r.Run(":" + config.CONFIG_RELEASE_SERVER_PORT) // listen and serve on 0.0.0.0:PORT
	}
}

//
// ----------------------------------------------------------------------------------
//
// 										/END OF	MAIN
//
// ----------------------------------------------------------------------------------
//
