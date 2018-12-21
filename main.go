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

func postMessage(apiEndPoint string, p *PostMessageParam) {
	raw, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(apiEndPoint, "application/json", bytes.NewReader(raw))
	if resp != nil {
		resp.Body.Close()
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	apiEndPoint := fmt.Sprintf("https://typetalk.com/api/v1/topics/%s?typetalkToken=%s",
		os.Getenv("PLUGIN_TYPETALK_TOKEN"),
		os.Getenv("PLUGIN_TOPIC_ID"),
	)

	repo := &Repo{
		Owner: os.Getenv("DRONE_REPO_OWNER"),
		Name:  os.Getenv("DRONE_REPO_NAME"),
		Link:  os.Getenv("DRONE_REPO_LINK"),
	}

	build := &Build{
		Status: os.Getenv("DRONE_BUILD_STATUS"),
	}

	var message string

	template := os.Getenv("PLUGIN_TEMPLATE")
	if template == "" {
		message = buildDefaultMessage(repo, build)
	}

	p := &PostMessageParam{
		Message:      message,
		ShowLinkMeta: false,
	}

	postMessage(apiEndPoint, p)
}
