package menu

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadMenuFromJSON(path string) ([]MenuItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read menu file %s: %w", path, err)
	}
	var items []MenuItem
	err = json.Unmarshal(data, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return items, nil
}
