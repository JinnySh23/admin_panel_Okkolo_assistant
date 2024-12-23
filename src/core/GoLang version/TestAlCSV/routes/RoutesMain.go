// ------------------------------------
// RR IT 2024
//
// ------------------------------------

//
// ----------------------------------------------------------------------------------
//
// 								Routes (Path)
//
// ----------------------------------------------------------------------------------
//

package routes

import (
	// Internal project packages
	"rr/TestAlCSV/config"
	"rr/TestAlCSV/modules/rr_csv"
	"rr/TestAlCSV/modules/rr_debug"

	// Third-party libraries
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	// System Packages
	"net/http"
	"strconv"
)

// ----------------------------------------------
//
//	Structures
//
// ----------------------------------------------
type Admin_LoginJSON struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Dataset_NewDataJSON struct {
	NewQuestion string `json:"new_question"`
	NewAnswer   string `json:"new_answer"`
}

// ----------------------------------------------
//
// 				Root requests
//
// ----------------------------------------------
//HTML-path

// /adm-panel
func Handler_LordPanel(c *gin.Context) {
	// We check the session, if there is one, we transfer it to the page, if not, to the login.
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id != nil {
		c.HTML(http.StatusOK, "lord_panel.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

// /login
func Handler_Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	}

	if c.Request.Method == "POST" {
		json_data := new(Admin_LoginJSON)
		err := c.ShouldBindJSON(&json_data)

		// Checking whether the JSON has arrived or the hat
		if err != nil {
			rr_debug.PrintLOG("RoutesMain.go", "Handler_Login", "c.ShouldBindJSON", "Неверные данные в запросе", err.Error())
			if config.CONFIG_IS_DEBUG {
				Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message+" Error: "+err.Error())
			} else {
				Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message)
			}
			return
		}

		if json_data.Login == config.CONFIG_DEFAULT_LOGIN && json_data.Password == config.CONFIG_DEFAULT_PASSWORD {
			session := sessions.Default(c)
			session.Set("session_user_id", strconv.Itoa(-1))
			session.Save()
			Answer_OK(c)
		} else {
			Answer_Unauthorized(c, ANSWER_INVALID_CREDENTIALS().Code, ANSWER_INVALID_CREDENTIALS().Message)
		}
		return
	}
}

// GET /api/dataset
func Handler_API_Dataset_GetData(c *gin.Context) {
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id != nil {
		// Reading data from a CSV file
		data, err := rr_csv.ReadCSVData(config.DATASET_FILE_PATH)
		if err != nil {
			Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message)
			return
		}
		Answer_SendObject(c, data)
	} else {
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message)
		return
	}
}

// POST /api/dataset
func Handler_API_Dataset_AddData(c *gin.Context) {
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id != nil {

		json_data := new(Dataset_NewDataJSON)
		err_bind := c.ShouldBindJSON(&json_data)

		if err_bind != nil {
			if config.CONFIG_IS_DEBUG {
				Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message+" Error: "+err_bind.Error())
			} else {
				Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message)
			}
			return
		}

		// Checking for required fields
		if (json_data.NewQuestion == "") && (json_data.NewAnswer == "") {
			Answer_BadRequest(c, ANSWER_EMPTY_FIELDS().Code, ANSWER_EMPTY_FIELDS().Message)
			return
		}

		// Reading data from a CSV file
		err := rr_csv.AppendToCSV(config.DATASET_FILE_PATH, json_data.NewQuestion, json_data.NewAnswer)
		if err != nil {
			Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message)
			return
		}
		Answer_OK(c)
	} else {
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message)
		return
	}
}

// DELETE /api/dataset
func Handler_API_Dataset_DeleteData(c *gin.Context) {
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id != nil {

		// Getting the index from the URL parameter
		index_str := c.Query("index")
		if index_str == "" {
			Answer_BadRequest(c, ANSWER_INVALID_URL_PARAMETER().Code, ANSWER_INVALID_URL_PARAMETER().Message)
			return
		}

		// Converting the index to a number
		index, err := strconv.Atoi(index_str)
		if err != nil {
			Answer_BadRequest(c, ANSWER_INVALID_STRING_TO_INT_CONVERSION().Code, ANSWER_INVALID_STRING_TO_INT_CONVERSION().Message)
			return
		}

		err = rr_csv.RemoveRowFromCSV(config.DATASET_FILE_PATH, index)
		if err != nil {
			Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message)
			return
		}

		Answer_OK(c)
	} else {
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message)
		return
	}
}

// /logout
func Handler_Logout(c *gin.Context) {
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id == nil {
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message)
		return
	} else {
		session.Clear()
		Answer_OK(c)
		return
	}
}
