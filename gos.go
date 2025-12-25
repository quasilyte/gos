package gos

import "runtime"

const (
	IsWindows = runtime.GOOS == "windows"
	IsLinux   = runtime.GOOS == "linux"
	IsMacos   = runtime.GOOS == "darwin"
	IsWeb     = runtime.GOOS == "js"
	IsAndroid = runtime.GOOS == "android"
)
