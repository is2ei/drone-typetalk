package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	baseURL = "https://typetalk.com"
)

type (

	// Drone contains drone information.
	Drone struct {
		IsDrone string
		Branch  string
	}

	// Repo contains repository information.
	Repo struct {
		FullName   string
		Owner      string
		Name       string
		Branch     string
		Link       string
		HTTPURL    string
		NameSpace  string
		Private    string
		Visibility string
		SCM        string
	}

	// Build contains build information.
	Build struct {
		Number string
		Status string
		Event  string
		Link   string
	}

	// Commit contains current commit information.
	Commit struct {
		SHA     string
		Ref     string
		Branch  string
		Author  string
		Pull    string
		Message string
	}

	// Env contains environment variables value.
	Env struct {
		Drone  *Drone
		Repo   *Repo
		Build  *Build
		Commit *Commit
	}

	// PostMessageRequestParam contains parameters for POST reqeust
	PostMessageRequestParam struct {
		Message      string `json:"message"`
		ShowLinkMeta bool   `json:"showLinkMeta,omitempty"`
	}
)

func buildDefaultMessage(repo *Repo, build *Build) string {
	return fmt.Sprintf("[[%s/%s](%s):%s] [Build#%s: %s](%s)",
		repo.Owner,
		repo.Name,
		repo.Link,
		repo.Branch,
		build.Number,
		build.Status,
		build.Link,
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

	drone := &Drone{
		IsDrone: os.Getenv("DRONE"),
		Branch:  os.Getenv("DRONE_BRANCH"),
	}

	repo := &Repo{
		FullName:   os.Getenv("DRONE_REPO"),
		Owner:      os.Getenv("DRONE_REPO_OWNER"),
		Name:       os.Getenv("DRONE_REPO_NAME"),
		Branch:     os.Getenv("DRONE_REPO_BRANCH"),
		Link:       os.Getenv("DRONE_REPO_LINK"),
		NameSpace:  os.Getenv("DRONE_REPO_NAMESPACE"),
		Private:    os.Getenv("DRONE_REPO_PRIVATE"),
		Visibility: os.Getenv("DRONE_REPO_VISIBILITY"),
		SCM:        os.Getenv("DRONE_REPO_SCM"),
	}

	build := &Build{
		Number: os.Getenv("DRONE_BUILD_NUMBER"),
		Status: os.Getenv("DRONE_BUILD_STATUS"),
		Event:  os.Getenv("DRONE_BUILD_EVENT"),
		Link:   os.Getenv("DRONE_BUILD_LINK"),
	}

	commit := &Commit{
		SHA:     os.Getenv("DRONE_COMMIT_SHA"),
		Ref:     os.Getenv("DRONE_COMMIT_REF"),
		Branch:  os.Getenv("DRONE_COMMIT_BRANCH"),
		Author:  os.Getenv("DRONE_COMMIT_AUTHOR"),
		Pull:    os.Getenv("DRONE_PULL_REQUEST"),
		Message: os.Getenv("DRONE_COMMIT_MESSAGE"),
	}

	env := &Env{
		Drone:  drone,
		Repo:   repo,
		Build:  build,
		Commit: commit,
	}

	var message string

	t := os.Getenv("PLUGIN_TEMPLATE")
	if t == "" {
		message = buildDefaultMessage(repo, build)
	} else {
		tmpl, _ := template.New("message").Parse(t)
		var b bytes.Buffer
		tmpl.Execute(&b, env)
		message = b.String()
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
