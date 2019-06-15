package logger

import "github.com/fatih/color"

func Info(format string, a ...interface{}) {
	color.White(format, a...)
}

func Success(format string, a ...interface{}) {
	color.Green(format, a...)
}

func Warn(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

func Error(format string, a ...interface{}) {
	color.Red(format, a...)
}

func Notice(format string, a ...interface{}) {
	d := color.New(color.FgBlue, color.Bold)
	d.Printf(format, a...)
}
