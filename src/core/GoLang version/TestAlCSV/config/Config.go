// ------------------------------------
// RR IT 2024
//
// ------------------------------------

package config

const (

	// ===========================================================
	// 							ROUTER CONFIG
	// ===========================================================

	// Printing debugging to the terminal
	CONFIG_PRINT_LOG = true

	// Basic real URL
	CONFIG_URL_BASE = "YOUR_SITE"

	// PORT
	CONFIG_RELEASE_SERVER_PORT          = "YOUR_HTTP_PORT (8083, 8084, etc)"
	CONFIG_DEBUG_SERVERLESS_SERVER_PORT = "YOUR_HTTP_PORT (8083, 8084, etc)"

	// Debugging level
	CONFIG_DEBUG_LEVEL = 1

	// Debugging mode
	CONFIG_IS_DEBUG = false

	// Using an internal server (for debugging)
	CONFIG_IS_DEBUG_SERVERLESS = true

	// The key for encrypting tokens (secret)
	CONFIG_SECRET = "YOUR_RANDOM_TOKEN"

	CONFIG_DEFAULT_LOGIN    = "YOUR_ADMIN_LOGIN_FOR_THE_SITE"
	CONFIG_DEFAULT_PASSWORD = "YOUR_ADMIN_PASSWORD_FOR_THE_SITE"

	//
	// File Storage
	//
	DATASET_FILE_PATH = "YOUR_DATASET_FILE_PATH"
)
