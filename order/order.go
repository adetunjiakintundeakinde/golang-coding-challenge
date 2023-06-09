package order

import (
	"errors"
	"io"
	"os"
)

type Order struct {
	Id        string
	Recipient string
	Date      string
	Cost      string
}

func ResourceAdd(order Order, orders []Order) (*[]Order, error) {
	if !isStateFileEmpty() && len(orders) == 0 {
		return nil, errors.New("state file is not empty")
	}
	orders = append(orders, order)
	return &orders, nil
}

func isStateFileEmpty() bool {
	state := readStateFile()

	if state == "" || state == "{}" {
		return true
	}
	return false
}

func readStateFile() string {
	stateFile, _ := os.Open("state.json")
	defer stateFile.Close()

	bytes, _ := io.ReadAll(stateFile)
	content := string(bytes)

	return content
}

func ImportState() *[]Order {
	var orders []Order
	return &orders
}
