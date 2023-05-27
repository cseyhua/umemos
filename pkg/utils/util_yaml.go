package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func YamlToStruct(file string, s interface{}) error {
	// 读取文件
	yamlConfig, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlConfig, s)
	if err != nil {
		return err
	}
	return nil
}
