package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type (

	// Repo contains repository information.
	Repo struct {
		Owner string
		Name  string
	}

	// Build contains build information.
	Build struct {
		Status string
	}
)

func main() {
	token := os.Getenv("PLUGIN_TYPETALK_TOKEN")
	topicID := os.Getenv("PLUGIN_TOPIC_ID")
	template := os.Getenv("PLUGIN_TEMPLATE")

	repo := &Repo{
		Name: os.Getenv("DRONE_REPO_NAME"),
	}

	build := &Build{
		Status: os.Getenv("DRONE_BUILD_STATUS"),
	}

	endPoint := fmt.Sprintf("https://typetalk.com/api/v1/topics/%s?typetalkToken=%s", topicID, token)

	var message string

	if template == "" {
		message = fmt.Sprintf("[%s] %s",
			repo.Name,
			build.Status,
		)
	}

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
