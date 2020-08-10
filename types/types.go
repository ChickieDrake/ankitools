package types

type Note struct {
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
