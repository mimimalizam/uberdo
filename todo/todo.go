package todo

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	// Printf("%s\n", b) is more effirient; it doesn't create a copy.

	return nil
}
