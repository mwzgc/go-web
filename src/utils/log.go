package utils

import "github.com/kataras/golog"

// Log ...
var Log *golog.Logger = nil

// SetLog ...
func SetLog(log *golog.Logger) {
	Log = log
}
