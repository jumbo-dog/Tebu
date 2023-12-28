package dialog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	internalModels "tebu-discord/internal/models"
)

func GetDialog(title string) []internalModels.Dialog {
	absPath, err := filepath.Abs("../../internal/json_files/" + title + ".json")
	if err != nil {
		fmt.Println("Error getting absolute file path:", err)
	}

	fileContent, err := os.ReadFile(absPath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil
	}

	var dialogs []internalModels.Dialog

	err = json.Unmarshal(fileContent, &dialogs)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	return dialogs
}
