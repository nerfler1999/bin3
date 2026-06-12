package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешна")
	return nil

}

func CheckJson(name string) bool {
	isJson := strings.HasSuffix(name, ".json")
	return isJson
}
