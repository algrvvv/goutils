package jsonutils

import (
	"fmt"
	"github.com/algrvvv/goutils/osutils"
	"os"
)

// CheckConfigAndWriteLayout function for creating a json file and filling it with default
func CheckConfigAndWriteLayout(jsonLayout, path string) error {
	if osutils.FileExist(path) {
		return nil
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
	}(file)

	_, err = file.WriteString(jsonLayout)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
