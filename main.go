package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	droneTemplate "github.com/drone/drone-template-lib/template"
)

const (
	baseURL = "https://typetalk.com"
)

type (

	// Drone contains drone information.
	Drone struct {
		IsDrone   string
		Hostname  string
		RemoteURL string
	}

	// System contains drone server information.
	System struct {
		Host     string
		Hostname string
		Version  string
	}

	// Runner contains drone agent information.
	Runner struct {
		Host     string
		Hostname string
		Platform string
		Label    string
	}

	// Git contains git information.
	Git struct {
		HTTPURL string
		SSHURL  string
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
		Branch       string
		Created      string
		Event        string
		Number       string
		Started      string
		Status       string
		Link         string
		PullRequest  string
		SourceBranch string
		TargetBranch string
		Tag          string
	}

	// Commit contains current commit information.
	Commit struct {
		Commit       string
		After        string
		Author       string
		AuthorAvatar string
		AuthorEmail  string
		AuthorName   string
		Before       string
		Branch       string
		Link         string
		SHA          string
		Ref          string
		Message      string
	}

	// Env contains environment variables value.
	Env struct {
		Drone  *Drone
		System *System
		Runner *Runner
		Git    *Git
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
		IsDrone:   os.Getenv("DRONE"),
		Hostname:  os.Getenv("DRONE_MACHINE"),
		RemoteURL: os.Getenv("DRONE_REMOTE_URL"),
	}

	system := &System{
		Host:     os.Getenv("DRONE_SYSTEM_HOST"),
		Hostname: os.Getenv("DRONE_SYSTEM_HOSTNAME"),
		Version:  os.Getenv("DRONE_SYSTEM_VERSION"),
	}

	runner := &Runner{
		Host:     os.Getenv("DRONE_RUNNER_HOST"),
		Hostname: os.Getenv("DRONE_RUNNER_HOSTNAME"),
		Platform: os.Getenv("DRONE_RUNNER_PLATFORM"),
		Label:    os.Getenv("DRONE_RUNNER_LABEL"),
	}

	git := &Git{
		HTTPURL: os.Getenv("DRONE_GIT_HTTP_URL"),
		SSHURL:  os.Getenv("DRONE_GIT_SSH_URL"),
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
		Created:      os.Getenv("DRONE_BUILD_CREATED"),
		Event:        os.Getenv("DRONE_BUILD_EVENT"),
		Number:       os.Getenv("DRONE_BUILD_NUMBER"),
		Started:      os.Getenv("DRONE_BUILD_STARTED"),
		Status:       os.Getenv("DRONE_BUILD_STATUS"),
		Branch:       os.Getenv("DRONE_BRANCH"),
		Link:         os.Getenv("DRONE_BUILD_LINK"),
		PullRequest:  os.Getenv("DRONE_PULL_REQUEST"),
		SourceBranch: os.Getenv("DRONE_SOURCE_BRANCH"),
		TargetBranch: os.Getenv("DRONE_TARGET_BRANCH"),
		Tag:          os.Getenv("DRONE_TAG"),
	}

	commit := &Commit{
		Commit:       os.Getenv("DRONE_COMMIT"),
		Message:      os.Getenv("DRONE_COMMIT_MESSAGE"),
		After:        os.Getenv("DRONE_COMMIT_AFTER"),
		Author:       os.Getenv("DRONE_COMMIT_AUTHOR"),
		AuthorAvatar: os.Getenv("DRONE_COMMIT_AUTHOR_AVATAR"),
		AuthorEmail:  os.Getenv("DRONE_COMMIT_AUTHOR_EMAIL"),
		AuthorName:   os.Getenv("DRONE_COMMIT_AUTHOR_NAME"),
		Before:       os.Getenv("DRONE_COMMIT_BEFORE"),
		Branch:       os.Getenv("DRONE_COMMIT_BRANCH"),
		Link:         os.Getenv("DRONE_COMMIT_LINK"),
		SHA:          os.Getenv("DRONE_COMMIT_SHA"),
		Ref:          os.Getenv("DRONE_COMMIT_REF"),
	}

	env := &Env{
		Drone:  drone,
		System: system,
		Runner: runner,
		Repo:   repo,
		Git:    git,
		Build:  build,
		Commit: commit,
	}

	var message string

	t := os.Getenv("PLUGIN_TEMPLATE")
	tRaw := os.Getenv("PLUGIN_TEMPLATE_RAW")
	if t == "" && tRaw == "" {
		message = buildDefaultMessage(repo, build)
	} else if t != "" {
		message, _ = droneTemplate.RenderTrim(t, env)
	} else {
		tmpl, _ := template.New("message").Parse(tRaw)
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
