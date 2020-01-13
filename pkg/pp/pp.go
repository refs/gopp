package pp

import (
	"encoding/json"
	"fmt"
)

// Print just dumps a map[string]interface{} so that a human can parse it
func Print(whatevs map[string]interface{}) {
	b, err := json.MarshalIndent(whatevs, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
