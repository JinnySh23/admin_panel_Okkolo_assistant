// ------------------------------------
// RR IT 2024
//
// ------------------------------------

//
// ----------------------------------------------------------------------------------
//
// 						JSON Answers (Standard responses)
//
// ----------------------------------------------------------------------------------
//

package routes

import (
	// Internal project packages
	"rr/TestAlCSV/config"

	// Third-party libraries
	"github.com/gin-gonic/gin"

	// System Packages
	"net/http"
)

// Structure for engine responses
type EngineAnswer struct {
	Code    int
	Message string
}

// ----------------------------------
// STANDARD ANSWERS, version 11_2020
// ----------------------------------
func ANSWER_OK() EngineAnswer {
	return EngineAnswer{
		Code:    0,
		Message: "OK",
	}
}

func ANSWER_OBJECT_EXISTS() EngineAnswer {
	return EngineAnswer{
		Code:    1,
		Message: "Object exists",
	}
}

func ANSWER_OBJECT_NOT_FOUND() EngineAnswer {
	return EngineAnswer{
		Code:    2,
		Message: "Object not found",
	}
}

func ANSWER_INVALID_JSON() EngineAnswer {
	return EngineAnswer{
		Code:    3,
		Message: "Invalid JSON",
	}
}

func ANSWER_EMPTY_FIELDS() EngineAnswer {
	return EngineAnswer{
		Code:    4,
		Message: "Empty fields",
	}
}

func ANSWER_UNEXPECTED_ERROR() EngineAnswer {
	return EngineAnswer{
		Code:    5,
		Message: "Unexpected error",
	}
}

func ANSWER_INVALID_CREDENTIALS() EngineAnswer {
	return EngineAnswer{
		Code:    6,
		Message: "Invalid credentials",
	}
}

func ANSWER_LOGIN_REQUIRED() EngineAnswer {
	return EngineAnswer{
		Code:    7,
		Message: "Login required",
	}
}

func ANSWER_PERMISSION_DENIED() EngineAnswer {
	return EngineAnswer{
		Code:    8,
		Message: "Permission denied (no authority)",
	}
}

func ANSWER_FILE_ERROR_TOO_LARGE() EngineAnswer {
	return EngineAnswer{
		Code:    9,
		Message: "File too large",
	}
}

func ANSWER_FILE_ERROR_INVALID_TYPE() EngineAnswer {
	return EngineAnswer{
		Code:    10,
		Message: "Invalid file type",
	}
}

// ----------------------------------
// Custom responses
// ----------------------------------
func ANSWER_INVALID_SESSION() EngineAnswer {
	return EngineAnswer{
		Code:    500,
		Message: "Invalid session",
	}
}

// Error converting from JSON to string
func ANSWER_INVALID_JSON_TO_STRING_CONVERSION() EngineAnswer {
	return EngineAnswer{
		Code:    503,
		Message: "Invalid JSON to string conversion",
	}
}

// Error converting from string to JSON
func ANSWER_INVALID_STRING_TO_JSON_CONVERSION() EngineAnswer {
	return EngineAnswer{
		Code:    504,
		Message: "Invalid string to JSON conversion",
	}
}

// Error converting from string to fraction
func ANSWER_INVALID_STRING_TO_FLOAT_CONVERSION() EngineAnswer {
	return EngineAnswer{
		Code:    505,
		Message: "Invalid string to float conversion",
	}
}

// Error converting from string to date
func ANSWER_INVALID_STRING_TO_DATE_CONVERSION() EngineAnswer {
	return EngineAnswer{
		Code:    506,
		Message: "Invalid string to date conversion",
	}
}

// Error converting from string to number
func ANSWER_INVALID_STRING_TO_INT_CONVERSION() EngineAnswer {
	return EngineAnswer{
		Code:    507,
		Message: "Invalid string to int conversion",
	}
}

// URL parameter error
func ANSWER_INVALID_URL_PARAMETER() EngineAnswer {
	return EngineAnswer{
		Code:    508,
		Message: "Invalid URL parameter",
	}
}

// Wrong command
func ANSWER_INVALID_COMMAND() EngineAnswer {
	return EngineAnswer{
		Code:    510,
		Message: "Invalid command",
	}
}

// Error sending an external request
func ANSWER_SENDING_EXTERNAL_REQUEST_ERROR(error_message string) EngineAnswer {
	return EngineAnswer{
		Code:    511,
		Message: error_message,
	}
}

//
// Successful responses
//

// Successfully created a new object
func Answer_SendObjectID(c *gin.Context, new_object_id uint) {
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    ANSWER_OK().Code,
			"message": ANSWER_OK().Message,
		},
		"data": gin.H{
			"id": new_object_id,
		},
	})
}

// Give an object
func Answer_SendObject(c *gin.Context, object interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    ANSWER_OK().Code,
			"message": ANSWER_OK().Message,
		},
		"data": object,
	})
}

// Give a string
func Answer_SendString(c *gin.Context, str string) {
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    ANSWER_OK().Code,
			"message": ANSWER_OK().Message,
		},
		"data": str,
	})
}

// 200 - Successful request
func Answer_OK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    ANSWER_OK().Code,
			"message": ANSWER_OK().Message,
		},
		"data": nil,
	})
}

// Send the file
func Answer_File(c *gin.Context, filepath string) {
	if config.CONFIG_IS_DEBUG_SERVERLESS {
		// Send via an internal server
		// We remove / at the beginning
		// relative_path := filepath[:1]
		// log.Println(filepath[1:])
		c.File(filepath[1:])
	} else {
		// Send via NGINX X-Accel
		c.Writer.Header().Set("X-Accel-Redirect", filepath)
		c.String(http.StatusOK, "OK")
	}
}

//
// Responses with an error
//

// 403 Forbidden - prohibited (not authorized)
func Answer_Forbidden(c *gin.Context, error_code int, error_message string) {
	c.JSON(http.StatusForbidden, gin.H{
		"status": gin.H{
			"code":    error_code,
			"message": error_message,
		},
		"data": nil,
	})
}

// 404 Not Found - the object was not found
func Answer_NotFound(c *gin.Context, error_code int, error_message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": gin.H{
			"code":    error_code,
			"message": error_message,
		},
		"data": nil,
	})
}

// 400 Bad Request - request error
func Answer_BadRequest(c *gin.Context, error_code int, error_message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": gin.H{
			"code":    error_code,
			"message": error_message,
		},
		"data": nil,
	})
}

// 401 Unauth - incorrect authorization
func Answer_Unauthorized(c *gin.Context, error_code int, error_message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status": gin.H{
			"code":    error_code,
			"message": error_message,
		},
		"data": nil,
	})
}

// 429 Too Many Requests - multiple requests per unit of time
func Answer_TooManyRequests(c *gin.Context, error_code int, error_message string) {
	c.JSON(http.StatusTooManyRequests, gin.H{
		"status": gin.H{
			"code":    error_code,
			"message": error_message,
		},
		"data": nil,
	})
}

// 500 Internal Server Error - error on the server
func Answer_InternalServerError(c *gin.Context, error_code int, error_message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": gin.H{
			"code":    error_code,
			"message": error_message,
		},
		"data": nil,
	})
}
