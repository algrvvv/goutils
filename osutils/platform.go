package osutils

import "runtime"

// GetPlatform function to get OS platform
func GetPlatform() string {
	return runtime.GOOS
}
