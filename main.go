package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("TYPETALK_TOKEN")
	topicID := os.Getenv("TOPIC_ID")
	message := os.Getenv("MESSAGE")

	endPoint := fmt.Sprintf("https://typetalk.com/api/v1/topics/%s?typetalkToken=%s", topicID, token)

	msg := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	raw, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(endPoint, "application/json", bytes.NewReader(raw))
	if resp != nil {
		resp.Body.Close()
	}

	if err != nil {
		log.Fatalln(err)
	}
}
