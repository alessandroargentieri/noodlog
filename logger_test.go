package noodlog

import (
	"fmt"
	"os"
	"testing"
)

var defaultLogger = Logger{
	level:                infoLevel,
	logWriter:            os.Stdout,
	prettyPrint:          false,
	traceCaller:          false,
	traceCallerLevel:     5,
	obscureSensitiveData: false,
	sensitiveParams:      nil,
	colors:               false,
	colorMap:             colorMap,
}

var customLogger = Logger{
	level:                errorLevel,
	logWriter:            os.Stderr,
	prettyPrint:          true,
	traceCaller:          true,
	traceCallerLevel:     6,
	obscureSensitiveData: true,
	sensitiveParams:      []string{"password"},
	colors:               true,
	colorMap:             colorMap,
}

func toStr(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

func TestNewLogger(t *testing.T) {

	expected := toStr(defaultLogger)
	actual := toStr(*NewLogger())

	if actual != expected {
		t.Errorf("TestNewLogger failed: expected %s, got %s", expected, actual)
	}
}

func TestSetConfigsEmptyConfigs(t *testing.T) {
	expected := toStr(defaultLogger)
	actual := toStr(*NewLogger().SetConfigs(Configs{}))

	if actual != expected {
		t.Errorf("TestSetConfigsEmptyConfigs failed: expected %s, got %s", expected, actual)
	}
}

func TestSetConfigsFullConfigs(t *testing.T) {
	expected := toStr(customLogger)
	actual := toStr(*NewLogger().SetConfigs(Configs{
		LogLevel:             LevelError,
		LogWriter:            os.Stderr,
		JSONPrettyPrint:      Enable,
		TraceCaller:          Enable,
		SinglePointTracing:   Enable,
		Colors:               Enable,
		ObscureSensitiveData: Enable,
		SensitiveParams:      []string{"password"},
	}))

	if actual != expected {
		t.Errorf("TestSetConfigsFullConfigs failed: expected %s, got %s", expected, actual)
	}
}

/*func TestSetConfigsEmptyObject(t *testing.T) {

	errFormat := "TestSetConfigsEmptyObject failed: param %s expected %v, got %v"

	SetConfigs(Configs{})

	if logLevel != 3 {
		t.Errorf(errFormat, "logLevel", 5, logLevel)
	}
	if JSONPrettyPrint {
		t.Errorf(errFormat, "JSONPrettyPrint", false, JSONPrettyPrint)
	}
	if traceCallerEnabled {
		t.Errorf(errFormat, "traceCallerEnabled", false, traceCallerEnabled)
	}
	if traceCallerLevel != 5 {
		t.Errorf(errFormat, "traceCallerLevel", 5, traceCallerLevel)
	}
	if colorEnabled {
		t.Errorf(errFormat, "colorEnabled", false, colorEnabled)
	}
	if colorMap[traceLabel] != colorReset {
		t.Errorf(errFormat, "colorMap[traceLabel]", colorReset, colorMap[traceLabel])
	}
	if colorMap[debugLabel] != colorGreen {
		t.Errorf(errFormat, "colorMap[debugLabel]", colorGreen, colorMap[debugLabel])
	}
	if colorMap[infoLabel] != colorReset {
		t.Errorf(errFormat, "colorMap[infoLabel]", colorReset, colorMap[infoLabel])
	}
	if colorMap[warnLabel] != colorYellow {
		t.Errorf(errFormat, "colorMap[warnLabel]", colorYellow, colorMap[warnLabel])
	}
	if colorMap[errorLabel] != colorRed {
		t.Errorf(errFormat, "colorMap[errorLabel]", colorRed, colorMap[errorLabel])
	}
	if obscureSensitiveDataEnabled {
		t.Errorf(errFormat, "obscureSensitiveDataEnabled", false, obscureSensitiveDataEnabled)
	}
	if len(sensitiveParams) != 0 {
		t.Errorf(errFormat, "sensitiveParams", 0, len(sensitiveParams))
	}
}*/

/*func TestSetConfigsFullObject(t *testing.T) {

	sensitiveList := []string{"password", "age"}
	errFormat := "TestSetConfigsFullObject failed: param %s expected %v, got %v"

	SetConfigs(Configs{
		LogLevel:           LevelError,
		JSONPrettyPrint:    Enable,
		TraceCaller:        Enable,
		SinglePointTracing: Enable,
		Colors:             Enable,
		CustomColors: &CustomColors{
			Trace: Purple,
			Debug: Yellow,
			Info:  Red,
			Warn:  Blue,
			Error: Cyan,
		},
		ObscureSensitiveData: Enable,
		SensitiveParams:      sensitiveList,
	})

	if logLevel != 5 {
		t.Errorf(errFormat, "logLevel", 5, logLevel)
	}
	if !JSONPrettyPrint {
		t.Errorf(errFormat, "JSONPrettyPrint", true, JSONPrettyPrint)
	}
	if !traceCallerEnabled {
		t.Errorf(errFormat, "traceCallerEnabled", true, traceCallerEnabled)
	}
	if traceCallerLevel != 6 {
		t.Errorf(errFormat, "traceCallerLevel", 6, traceCallerLevel)
	}
	if !colorEnabled {
		t.Errorf(errFormat, "colorEnabled", true, colorEnabled)
	}
	if purpleCode := composeColor(colorPurple); colorMap[traceLabel] != purpleCode {
		t.Errorf(errFormat, "colorMap[traceLabel]", purpleCode, colorMap[traceLabel])
	}
	if yellowCode := composeColor(colorYellow); colorMap[debugLabel] != yellowCode {
		t.Errorf(errFormat, "colorMap[debugLabel]", yellowCode, colorMap[debugLabel])
	}
	if redCode := composeColor(colorRed); colorMap[infoLabel] != redCode {
		t.Errorf(errFormat, "colorMap[infoLabel]", redCode, colorMap[infoLabel])
	}
	if blueCode := composeColor(colorBlue); colorMap[warnLabel] != blueCode {
		t.Errorf(errFormat, "colorMap[warnLabel]", blueCode, colorMap[warnLabel])
	}
	if cyanCode := composeColor(colorCyan); colorMap[errorLabel] != cyanCode {
		t.Errorf(errFormat, "colorMap[errorLabel]", cyanCode, colorMap[errorLabel])
	}
	if !obscureSensitiveDataEnabled {
		t.Errorf(errFormat, "obscureSensitiveDataEnabled", true, obscureSensitiveDataEnabled)
	}
	if len(sensitiveParams) != 2 {
		t.Errorf(errFormat, "sensitiveParams", 2, len(sensitiveParams))
	}
}*/

/*func TestLogLevel(t *testing.T) {
	testMap := map[string]int{
		"trace":       1,
		"debug":       2,
		"info":        3,
		"warn":        4,
		"error":       5,
		"invalidName": 3,
	}

	for label, level := range testMap {
		LogLevel(label)
		if logLevel != level {
			t.Errorf("TestLogLevel failed: expected %d, got %d", level, logLevel)
		}
	}
}

func TestEnableDisableJSONPrettyPrint(t *testing.T) {
	errFormat := "TestEnableDisableJSONPrettyPrint failed: JSONPrettyPrint expected %t, got %t "
	EnableJSONPrettyPrint()
	if !JSONPrettyPrint {
		t.Errorf(errFormat, true, JSONPrettyPrint)
	}
	DisableJSONPrettyPrint()
	if JSONPrettyPrint {
		t.Errorf(errFormat, false, JSONPrettyPrint)
	}
}

var testsuite = []struct {
	in       interface{}
	expected string
}{
	{"This is a test", `"message":"This is a test"`},
	{42, `"message":42`},
	{13.75, `"message":13.75`},
	{false, `"message":false`},
}

func TestSimpleLogger(t *testing.T) {

	var b bytes.Buffer
	LogWriter(&b)                                             // we want to write log in memory
	SetConfigs(Configs{LogLevel: LevelInfo, Colors: Disable}) // we don't want any ANSI sequence in the strings

	for _, tt := range testsuite {
		b.Reset() // clears the buffer for the next log message
		Info(tt.in)
		if !strings.Contains(b.String(), tt.expected) {
			t.Errorf("Failed Test logger! Expected: %v, got: %v", tt.expected, b.String())
		}
	}
}*/

//TestEnableDisableObscureSensitiveData

//TestSetSensitiveParams

//TestSetComposeLog

//TestSetComposeMessage

//TestStringify

//TestAdaptMessage

//TestStrToObj

//TestObscureSensitiveData

//TestObscureParam
