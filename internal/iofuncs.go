package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	FROM    string `json:"from"`
	TO      string `json:"to"`
	MESSAGE string `json:"message"`
	Key     string
	Fan     string
}

func MarshallData(Message Message) ([]byte, error) {
	jsonData, err := json.Marshal(Message)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func UnmarshallData(MessageData []byte) (Message, error) {
	var Data Message
	err := json.Unmarshal(MessageData, &Data)
	if err != nil {
		return Message{}, err
	}
	return Data, nil
}

func GetInput() string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return ""
	}

	return scanner.Text()
}
