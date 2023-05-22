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

func ImportState() *[]Order {
	var orders []Order
	return &orders
}

func isStateFileEmpty() bool {
	stateFile, _ := os.Open("state.json")
	defer stateFile.Close()

	bytes, _ := io.ReadAll(stateFile)
	content := string(bytes)

	if content == "" || content == "{}" {
		return true
	}
	return false
}
