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
		Link  string
	}

	// Build contains build information.
	Build struct {
		Status string
	}

	// PostMessageParam contains parameters for POST reqeust
	//
	// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/post-message/
	PostMessageParam struct {
		Message      string `json:"message"`
		ShowLinkMeta bool   `json:"showLinkMeta,omitempty"`
	}
)

func buildDefaultMessage(repo *Repo, build *Build) string {
	return fmt.Sprintf("[[%s/%s](%s)] %s",
		repo.Owner,
		repo.Name,
		repo.Link,
		build.Status,
	)
}

func postMessage(message string, showLinkMeta bool) {

}

func main() {
	token := os.Getenv("PLUGIN_TYPETALK_TOKEN")
	topicID := os.Getenv("PLUGIN_TOPIC_ID")
	template := os.Getenv("PLUGIN_TEMPLATE")

	repo := &Repo{
		Owner: os.Getenv("DRONE_REPO_OWNER"),
		Name:  os.Getenv("DRONE_REPO_NAME"),
		Link:  os.Getenv("DRONE_REPO_LINK"),
	}

	build := &Build{
		Status: os.Getenv("DRONE_BUILD_STATUS"),
	}

	endPoint := fmt.Sprintf("https://typetalk.com/api/v1/topics/%s?typetalkToken=%s", topicID, token)

	var message string

	if template == "" {
		message = buildDefaultMessage(repo, build)
	}

	p := &PostMessageParam{
		Message:      message,
		ShowLinkMeta: false,
	}

	raw, err := json.Marshal(p)
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
