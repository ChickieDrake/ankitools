package types

type NoteInfo struct {
	NoteID    int
	ModelName string
	Tags      *[]string
	Fields    *map[string]Field
	CardIDs   *[]int `json:"cards"`
}

type Field struct {
	Value string
	Order int
}

type NoteUpdate struct {
	NoteID int                `json:"id"`
	Fields *map[string]string `json:"fields"`
}
