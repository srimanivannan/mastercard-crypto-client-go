package utils

import (
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
)

func PrettyPrintJSON(jsonInput string, message string) string {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(jsonInput), &obj)
	if err != nil {
		return ""
	}

	formatter := colorjson.NewFormatter()
	formatter.Indent = 2

	coloredResult, _ := formatter.Marshal(obj)
	fmt.Println(message)
	fmt.Println(string(coloredResult))
	return string(coloredResult)
}
