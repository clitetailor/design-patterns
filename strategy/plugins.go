package main

func NewAddField(fieldName string, fieldValue interface{}) *AddField {
	return &AddField{
		field: fieldName,
		value: fieldValue,
	}
}

type AddField struct {
	field string
	value interface{}
}

func (plugin *AddField) Process(event map[string]interface{}) map[string]interface{} {
	event[plugin.field] = plugin.value

	return event
}

func NewDropField(fieldName string) *DropField {
	return &DropField{
		field: fieldName,
	}
}

type DropField struct {
	field string
}

func (plugin *DropField) Process(event map[string]interface{}) map[string]interface{} {
	delete(event, plugin.field)

	return event
}
