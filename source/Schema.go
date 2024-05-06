package goform

type Schema struct {
	Map map[string]any `json:"map"`
}

func NewSchema(themap map[string]any) {

	var schema Schema

	schema.Map = make(map[string]any)

	for entry := range themap {
		schema.Map[entry] = themap[entry]
	}

	return schema

}
