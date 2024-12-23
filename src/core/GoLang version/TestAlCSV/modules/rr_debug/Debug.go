// ------------------------------------
// RR IT 2024
//
// ------------------------------------

//
// ----------------------------------------------------------------------------------
//
// 				Debug (Debugging function module, main file)
//
// ----------------------------------------------------------------------------------
//

package rr_debug

import (
	// Internal project packages
	"rr/TestAlCSV/config"

	// System Packages
	"encoding/json"
	"fmt"
	"log"
)

const (
	Reset     = "\033[0m"
	RedBg     = "\033[41m"
	GreenBg   = "\033[42m"
	YellowBg  = "\033[43m"
	BlueBg    = "\033[44m"
	WhiteText = "\033[97m"
	// Add other colors as needed.
)

func ColorBoxText(text string, bgColor string, textColor string) string {
	return fmt.Sprintf("%s%s%s%s%s", bgColor, textColor, " "+text+" ", Reset, WhiteText)
}

func PrintLOG(file_name string, parent_function string, category string, error_type string, error_message string) {
	if config.CONFIG_PRINT_LOG {
		logPrefix := fmt.Sprintf("[FILE: %s]:[PARENT FUNCTION: %s]:[%s]: ", file_name, parent_function, category)

		var coloredErrorType string
		coloredErrorType = ColorBoxText(error_type, RedBg, WhiteText)

		if error_message == "" {
			fmt.Println(coloredErrorType + " " + logPrefix + error_type)
		} else {
			fmt.Println(coloredErrorType + " " + logPrefix + error_type + " | Message: " + error_message)
		}
	}
}

// Printing an object in the console is beautiful
func PrintObject(object interface{}) error {

	log.Println("---------------------------------")

	b, err := json.MarshalIndent(object, "", "  ")

	if err != nil {
		return err
	}

	log.Println(string(b))

	log.Println("---------------------------------")

	return nil

}
