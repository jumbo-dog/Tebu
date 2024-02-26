package models

type Dialog struct {
	ButtonLabel []string `json:"button_label"`
	DialogText  []string `json:"dialog_text"`
}

type Enemy struct {
	Health      int
	Name        string
	Description string
}
