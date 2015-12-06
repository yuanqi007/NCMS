package commonTools

import (
	"os"
	"path/filepath"
)

func GetCurrentPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
