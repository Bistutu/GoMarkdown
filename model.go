package gomarkdown

// TaskModel task model
type TaskModel struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// DefinitionList definition list
type DefinitionList struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
}
