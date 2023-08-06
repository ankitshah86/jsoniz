package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/ankitshah86/jsoniz/internal/config"
)

// ValidateJson validates the json string
func ValidateJson(str string) bool {
	var js map[string]interface{}

	err := json.Unmarshal([]byte(str), &js)
	if err != nil {
		fmt.Println("could not parse json :", err)
		return false
	}

	//check for duplicate keys recursively

	if config.IsJsonDuplicateKeyCheckEnabled() {
		//TODO: implement duplicate key check
		fmt.Println("duplicate key check is not implemented yet")
	}

	return true
}
