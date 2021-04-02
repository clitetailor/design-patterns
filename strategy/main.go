package main

import "fmt"

func main() {
	processor := NewProcessor()

	processor.AddPlugin(NewAddField("message", "test.com -> 127.0.0.1"))
	processor.AddPlugin(NewAddField("tags", []string{"json_parse_failed"}))
	processor.AddPlugin(NewDropField("log_type"))

	event := processor.Process(map[string]interface{}{
		"log_type": "dns",
	})

	fmt.Println(event)
}
