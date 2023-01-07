package config

import (
	"errors"
	"os"
	"strings"

	"github.com/rrojan/gin-skeleton/utils"
	"gopkg.in/yaml.v3"
)

// LoadDataFromConfig loads returns all data from a ".yml" config
func LoadDataFromConfig(fileName string) map[string]interface{} {
	if !strings.Contains(fileName, ".yml") {
		fileName += ".yml"
	}
	file, err := os.ReadFile(fileName)
	utils.PanicIf(err)

	fileData := make(map[string]interface{})
	if err = yaml.Unmarshal(file, &fileData); err != nil {
		panic(errors.New("Cannot unmarshal config file :: " + fileName))
	}
	return fileData
}

// LoadKeyFromConfig loads returns single key data from a ".yml" config
func LoadKeyFromConfig(fileName string, key string) interface{} {
	fileData := LoadDataFromConfig(fileName)
	return fileData[key]
}
