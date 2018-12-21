package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	baseURL = "https://typetalk.com"
)

type (

	// Repo contains repository information.
	Repo struct {
		Owner    string
		FullName string
		Name     string
		Branch   string
		Link     string
	}

	// Build contains build information.
	Build struct {
		Status string
	}

	// PostMessageRequestParam contains parameters for POST reqeust
	PostMessageRequestParam struct {
		Message      string `json:"message"`
		ShowLinkMeta bool   `json:"showLinkMeta,omitempty"`
	}
)

func buildDefaultMessage(repo *Repo, build *Build) string {
	return fmt.Sprintf("[[%s/%s](%s):%s] %s",
		repo.Owner,
		repo.Name,
		repo.Link,
		repo.Branch,
		build.Status,
	)
}

// PostMessage posts a message to Typetalk
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/post-message/
func PostMessage(baseURL, topicID, token string, p *PostMessageRequestParam) (*http.Response, error) {

	apiEndPoint := fmt.Sprintf("%s/api/v1/topics/%s?typetalkToken=%s",
		baseURL,
		topicID,
		token,
	)

	raw, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(apiEndPoint, "application/json", bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}

	if resp != nil {
		resp.Body.Close()
	}

	return resp, err
}

func main() {

	repo := &Repo{
		Owner:    os.Getenv("DRONE_REPO_OWNER"),
		FullName: os.Getenv("DRONE_REPO"),
		Name:     os.Getenv("DRONE_REPO_NAME"),
		Branch:   os.Getenv("DRONE_REPO_BRANCH"),
		Link:     os.Getenv("DRONE_REPO_LINK"),
	}

	build := &Build{
		Status: os.Getenv("DRONE_BUILD_STATUS"),
	}

	var message string

	template := os.Getenv("PLUGIN_TEMPLATE")
	if template == "" {
		message = buildDefaultMessage(repo, build)
	}

	p := &PostMessageRequestParam{
		Message:      message,
		ShowLinkMeta: false,
	}

	PostMessage(
		baseURL,
		os.Getenv("PLUGIN_TOPIC_ID"),
		os.Getenv("PLUGIN_TYPETALK_TOKEN"),
		p,
	)
}
