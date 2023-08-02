package helpers

import (
	"encoding/json"
	"fmt"
	"strings"
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
	if err := checkDuplicateKeys(str); err != nil {
		fmt.Println("duplicate keys found")
		return false
	}

	return true
}

// checkDuplicateKeys checks for duplicate keys recursively
// returns false if duplicate keys are found
func checkDuplicateKeys(jsonString string) error {
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	return checkKeys(decoder, make(map[string]bool))
}

func checkKeys(decoder *json.Decoder, keys map[string]bool) error {
	t, err := decoder.Token()
	if err != nil {
		return err
	}

	// Verify that we have an object (denoted by {)
	if delim, ok := t.(json.Delim); !ok || string(delim) != "{" {
		return fmt.Errorf("expected start of object, got %v", t)
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		key, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected a key, got %v", t)
		}

		if _, duplicate := keys[key]; duplicate {
			fmt.Printf("Duplicate key found: %s\n", key)
			return fmt.Errorf("duplicate key found: %s", key)
		}

		keys[key] = true

		t, err = decoder.Token()
		if err != nil {
			return err
		}

		if delim, ok := t.(json.Delim); ok && string(delim) == "{" {
			// We have a nested object, recursively check keys
			if err := checkKeys(decoder, make(map[string]bool)); err != nil {
				return err
			}
		}
	}

	t, err = decoder.Token() // Consume closing delimiter
	if err != nil {
		return err
	}

	if delim, ok := t.(json.Delim); !ok || string(delim) != "}" {
		return fmt.Errorf("expected end of object, got %v", t)
	}

	return nil
}
