package logx

import (
	"fmt"
	"log"
)

// Color an int signifying the ANSI escape code for a given color
// example: Black is 30 because to make black text is '\x1b[30mHello World!\x1b[0m'
type Color int

// Color codes added for easy reference. Not all are used yet :)
const (
	Black     Color = 30
	Red       Color = 31
	Green     Color = 32
	Yellow    Color = 33
	Blue      Color = 34
	Magenta   Color = 35
	Cyan      Color = 36
	White     Color = 37
	BlackBG   Color = 40
	RedBG     Color = 41
	GreenBG   Color = 42
	YellowBG  Color = 43
	BlueBG    Color = 44
	MagentaBG Color = 45
	CyanBG    Color = 46
	WhiteBG   Color = 47
)

type LogLevel int
type LogType struct {
	Level LogLevel
	Name  string
	Icon  string
	Color Color
}

const (
	Debug     LogLevel = 1
	Info      LogLevel = 2
	Notice    LogLevel = 3
	Warning   LogLevel = 4
	Error     LogLevel = 5
	Critical  LogLevel = 6
	Alert     LogLevel = 7
	Emergency LogLevel = 8
)

var logTypes = map[LogLevel]LogType{
	Debug:     {Debug, "DEBUG", "😎", Cyan},
	Info:      {Info, "INFO", "ℹ️", Green},
	Notice:    {Notice, "NOTICE", "📌", Yellow},
	Warning:   {Warning, "WARNING", "⚠️", Magenta},
	Error:     {Error, "ERROR", "💩", Red},
	Critical:  {Critical, "CRITICAL", "😱", Red},
	Alert:     {Alert, "ALERT", "☠️", Red},
	Emergency: {Emergency, "EMERGENCY", "☢️", RedBG},
}

// MakeColoredText makes colored string output for logs
func MakeColoredText(colorCode Color, str interface{}) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, str)
}

// MakeColoredTextFunc returns a function that makes strings with a preset color
func MakeColoredTextFunc(colorCode Color) func(str interface{}) string {
	return func(str interface{}) string {
		return MakeColoredText(colorCode, str)
	}
}

func MakeColoredTextWithBackgroundColor(color Color, background Color, str interface{}) string {
	coloredString := MakeColoredText(color, str)
	return MakeColoredText(background, coloredString)
}

var GREEN = MakeColoredTextFunc(Green)

// MakeLogPrefix formats a prefix for a log
func MakeLogPrefix(level LogLevel) string {
	lt := logTypes[level]
	return fmt.Sprintf(`%s  %s -`, lt.Icon, MakeColoredTextFunc(lt.Color)(string(logTypes[level].Name)))
}

func MakeLeveledLoggerContent(level LogLevel) func(str interface{}) string {
	return func(str interface{}) string {
		prefix := MakeLogPrefix(level)
		return fmt.Sprintf("%s %v", prefix, str)
	}
}

// LogWithLevel logs something with the correct log level formatting.
func LogWithLevel(level LogLevel, str interface{}) {
	log.Println(MakeLeveledLoggerContent(level)(str))
}
