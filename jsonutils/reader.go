package jsonutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// GetValueByKey Obtaining data from the json file using key access
// through the point. Example: settings.parameter -> will get the data that is
// in settings and then in parameter. Will return either a value or a map with other data
func GetValueByKey(path, key string) (string, error) {
	splitKey := strings.Split(key, ".")

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error opening file:", err)
		}
	}(file)

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "", err
	}

	var result = data
	resultString := ""
	for i := 0; i < len(splitKey); i++ {
		value := result[splitKey[i]]
		if value != nil {
			switch v := value.(type) {
			case string:
				resultString = v
				break
			default:
				result = result[splitKey[i]].(map[string]interface{})
			}
		} else {
			return "", errors.New("key not found")
		}

	}

	return resultString, nil
}
