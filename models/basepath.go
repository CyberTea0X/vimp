package models

import (
	"path"
	"path/filepath"
	"runtime"
)

var _, b, _, _ = runtime.Caller(0)
var BasePath = path.Join(filepath.Dir(b), "../.")
