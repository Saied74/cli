package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ItemValidator is supplied by the user to validate user enterires.
type ItemValidator func(string) bool

//Item holds the values to be displayed for each line item.
type Item struct {
	Name      string
	Prompt    string
	Response  string
	Value     string
	Validator ItemValidator
}

//ItemResponse is sent back to the user though the channel.
type ItemResponse struct {
	Name  string
	Value string
}

//Items is the table of all the data to be displayed
type Items struct {
	OrderList   []string
	ItemList    map[string]*Item
	ActionLines []string
	sender      chan ItemResponse
}

//Command is called to start the CLI
func Command(t *Items) chan ItemResponse {
	t.sender = make(chan ItemResponse)
	go t.run()
	return t.sender
}

func (t *Items) run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		for i, item := range t.OrderList {
			fmt.Printf("%d.\t%s %s\n", i+1, t.ItemList[item].Prompt, t.ItemList[item].Value)
		}
		for _, action := range t.ActionLines {
			fmt.Println(action)
		}
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		switch input {
		case "c":
			r := ItemResponse{
				Name:  "Continue",
				Value: input,
			}
			t.sender <- r
			continue
		case "q":
			r := ItemResponse{
				Name:  "Quit",
				Value: input,
			}
			t.sender <- r
			continue
		default:
			n, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("You did not enter a number, you entered: ", input)
				continue
			}
			if n <= len(t.OrderList) {
				n--
				itemName := t.OrderList[n]
				fmt.Println(t.ItemList[itemName].Prompt)
				input, _ = reader.ReadString('\n')
				input = strings.TrimSuffix(input, "\n")
				if !t.ItemList[itemName].Validator(input) {
					fmt.Println("You did not enter a valid value, you entered: ", input)
					continue
				}
				t.ItemList[itemName].Value = input
				r := ItemResponse{
					Name:  itemName,
					Value: input,
				}
				t.sender <- r
			} else {
				fmt.Printf("Item number %d is too large, try again\n", n)
				continue
			}
		}
	}
}
