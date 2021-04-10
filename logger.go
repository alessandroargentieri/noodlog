package noodlog

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var logLevel int = infoLevel
var JSONPrettyPrint bool = false

var logLevels = map[string]int{
	traceLabel: traceLevel,
	debugLabel: debugLevel,
	infoLabel:  infoLevel,
	warnLabel:  warnLevel,
	errorLabel: errorLevel,
	panicLabel: panicLevel,
	fatalLabel: fatalLevel,
}

// LogLevel function sets the log level
func LogLevel(level string) {
	logLevel = logLevels[level]
	if logLevel == 0 {
		logLevel = infoLevel
	}
}

// EnableJSONPrettyPrint enables JSON pretty printing
func EnableJSONPrettyPrint() {
	JSONPrettyPrint = true
}

// DisableJSONPrettyPrint diables JSON pretty printing
func DisableJSONPrettyPrint() {
	JSONPrettyPrint = false
}

// Trace function prints a log with trace log level
func Trace(message ...interface{}) {
	printLog(traceLabel, message)
}

// Debug function prints a log with debug log level
func Debug(message ...interface{}) {
	printLog(debugLabel, message)
}

// Info function prints a log with info log level
func Info(message ...interface{}) {
	printLog(infoLabel, message)
}

// Warn function prints a log with warn log level
func Warn(message ...interface{}) {
	printLog(warnLabel, message)
}

// Error function prints a log with error log level
func Error(message ...interface{}) {
	printLog(errorLabel, message)
}

// Panic function prints a log with panic log level
func Panic(message ...interface{}) {
	panic(composeLog(panicLabel, message))
}

// Fatal function prints a log with fatal log level
func Fatal(message ...interface{}) {
	printLog(fatalLabel, message)
	os.Exit(1)
}

func printLog(label string, message []interface{}) {
	if logLevels[label] >= logLevel {
		fmt.Println(composeLog(label, message))
	}
}

func composeLog(level string, message []interface{}) string {
	caller := map[string]string{}
	if traceCallerEnabled {
		caller = traceCaller()
	}

	logMsg := record{
		File:     caller[file],
		Function: caller[function],
		Level:    level,
		Message:  composeMessage(message),
		Time:     strings.Split(time.Now().String(), "m")[0],
	}

	var jsn []byte
	if JSONPrettyPrint {
		jsn, _ = json.MarshalIndent(logMsg, "", "   ")
	} else {
		jsn, _ = json.Marshal(logMsg)
	}

	logRecord := string(jsn)
	if colorEnabled {
		logRecord = fmt.Sprintf("%s%s%s", colorMap[level], logRecord, colorReset)
	}

	return logRecord
}

func composeMessage(message []interface{}) interface{} {
	switch len(message) {
	case 0:
		return ""
	case 1:
		return adaptMessage(message[0])
	default:
		switch message[0].(type) {
		case string:
			msg0 := message[0].(string)
			if strings.Contains(msg0, "%") {
				return fmt.Sprintf(msg0, message[1:]...)
			}
		}
		return stringify(message)
	}
}

func stringify(message []interface{}) string {
	msg := ""
	for _, m := range message {
		if m != nil {
			msg = msg + fmt.Sprintf("%v", m)
		}
	}
	return msg
}

func adaptMessage(message interface{}) interface{} {
	// TODO
	return message
}
