package utils

import (
	"os"
	"path/filepath"
)

func GetPath4Current() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir //strings.Replace(dir, "\\", "/", -1)
}
