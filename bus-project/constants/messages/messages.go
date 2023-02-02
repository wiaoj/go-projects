package messages

import (
	"fmt"
	consoleColor "golang_projects/constants/consoleColor"
)

var FailedConnectingDatabase = ErrorMessage("Failed to connect to database")
var SuccessConnectingDatabase = SuccessMessage("Success to connect to database")

func SuccessMessage(message string) string {
	return fmt.Sprint(consoleColor.Green().Bold("SUCCESS: ").Apply(message))
}

func ErrorMessage(message string) string {
	return fmt.Sprint(consoleColor.RedLight().Bold("ERROR: ").Apply(message))
}

func WarningMessage(message string) string {
	return fmt.Sprint(consoleColor.YellowLight().Bold("WARNING: ").Apply(message))

}

// func getWarningMessage(message string) string {
// 	return fmt.Sprintf(`%s%s`, colorBlue, message)
// }

// func getSuccessMessage(message string) string {
// 	return consoleColor.Green(message)
// }

// var SuccesConnectingDatabase = getSuccessMessage()
