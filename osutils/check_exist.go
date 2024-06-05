package osutils

import "os"

// FileExist function to check whether a configuration file exists in the desired directory
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
