package common

import (
	"fmt"
)

const (
	//INFO add comment
	INFO = 1
	//FAULT add comment
	FAULT = 2
	//ALERT add comment
	ALERT = 3
)

var (
	//Console unused
	Console = false
	logDesc = []string{
		0:     "None",
		INFO:  `[!]`,
		FAULT: `[x]`,
		ALERT: `[√]`,
	}
	logFile = "dark_eye.log"
)

//LogBuild add comment
func LogBuild(module string, logCt string, level int) string {
	return fmt.Sprintf("%s /%s/ %s",
		logDesc[level],
		module,
		logCt)
}
