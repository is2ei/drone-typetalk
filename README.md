# drone-typetalk

Drone plugin for sending Typetalk notifications.

| Badge | Description |
| ------------- | ------------- |
| [![Build Status](https://travis-ci.org/is2ei/drone-slack.svg?branch=master)][travis] | Travis Build Status |
| [![Build Status](https://cloud.drone.io/api/badges/is2ei/drone-typetalk/status.svg)][drone] | Drone Build Status |
| [![Go Report Card](https://goreportcard.com/badge/github.com/is2ei/drone-typetalk)][goreport] | Go Report |

[travis]: https://travis-ci.org/is2ei/drone-slack
[drone]: https://cloud.drone.io/is2ei/drone-typetalk
[goreport]: https://goreportcard.com/report/github.com/is2ei/drone-typetalk

## Usage

The Typetalk plugin posts build status messages to your channel. The below configuration demonstrates  
simple usage.

```
steps:
- name: typetalk
  image: is2ei/typetalk
  settings:
    typetalk_token: xxxxxxxx
    topic_id: 12345
```

Example configuration with webhook sourced from a secret:

```
steps:
- name: typetalk
  image: is2ei/typetalk
  settings:
    typetalk_token:
      from_secret: typetalk_token
    topic_id:
      from_secret: topic_id
```

Example configuration for success and failure messages:

```
steps:
- name: typetalk
  image: is2ei/typetalk
  settings:
    typetalk_token: xxxxxxxx
    topic_id: 12345
  when:
    status: [ success, failure ]
```

## Parameters Reference

`typetalk_token`  
[Typetalk Token](https://developer.nulab-inc.com/docs/typetalk/#tttoken)

`topic_id`  
Topic ID to send notification.  

`template`  
Overwrite the default message temnplate.  
It uses `github.com/drone/drone-template-lib/template`.  

`template_raw`  
Overwrite the default message temnplate.  
It uses `text/template` If both `template` and `template_raw` values are set, `template` value will override the message template.  

## Template Raw Reference

`.Drone.IsDrone`  
Boolean value, indicates the runtime environment is Drone. [DRONE](https://docs.drone.io/reference/environ/drone/)  

`.Drone.Hostname`  
String literal, provides the Drone agent hostname. [DRONE_MACHINE](https://docs.drone.io/reference/environ/drone-machine/)  

`.Drone.Hostname`  
[DRONE_REMOTE_URL](https://docs.drone.io/reference/environ/drone-remote-url/)  

`.System.Host`  
String literal, provides the Drone server hostname. [DRONE_SYSTEM_HOST](https://docs.drone.io/reference/environ/drone-system-host/)  

`.System.Hostname`  
String literal, provides the Drone server hostname. [DRONE_SYSTEM_HOSTNAME](https://docs.drone.io/reference/environ/drone-system-hostname/)  

`.System.Version`  
String literal, provides the Drone server version. [DRONE_SYSTEM_VERSION](https://docs.drone.io/reference/environ/drone-system-version/)  

`.Runner.Host`  
String literal, provides the Drone agent hostname. [DRONE_RUNNER_HOST](https://docs.drone.io/reference/environ/drone-runner-host/)  

`.Runner.Hostname`  
String literal, provides the Drone agent hostname. [DRONE_RUNNER_HOSTNAME](https://docs.drone.io/reference/environ/drone-runner-hostname/)  

`.Runner.Platform`  
String literal, provides the Drone agent os and architecture. [DRONE_RUNNER_PLATFORM](https://docs.drone.io/reference/environ/drone-runner-platform/)  

`.Runner.Label`  
[DRONE_RUNNER_LABEL](https://docs.drone.io/reference/environ/drone-runner-label/)  

`.Git.HTTPURL`  
String literal, provides the repository git+http url. [DRONE_GIT_HTTP_URL](https://docs.drone.io/reference/environ/drone-git-http-url/)  

`.Git.SSHURL`  
String literal, provides the repository git+ssh url. [DRONE_GIT_SSH_URL](https://docs.drone.io/reference/environ/drone-git-ssh-url/)  

`.Repo.FullName`  
String literal, provides the full name of the repository. [DRONE_REPO](https://docs.drone.io/reference/environ/drone-repo/)  

`.Repo.Owner`  
repository owner DRONE_REPO_OWNER  

`.Repo.Name`  
String literal, provides the repository name. [DRONE_REPO_NAME](https://docs.drone.io/reference/environ/drone-repo-name/)  

`.Repo.Branch`  
String literal, provides the default repository branch (e.g. master). [DRONE_REPO_BRANCH](https://docs.drone.io/reference/environ/drone-repo-branch/)  

`.Repo.Link`  
String literal, provides the repository http link. [DRONE_REPO_LINK](https://docs.drone.io/reference/environ/drone-repo-link/)  

`.Repo.NameSpace`  
String literal, provides the repository namespace (e.g. account owner) [DRONE_REPO_NAMESPACE](https://docs.drone.io/reference/environ/drone-repo-namespace/)  

`.Repo.Private`  
Boolean value, indicates the repository is public or private. [DRONE_REPO_PRIVATE](https://docs.drone.io/reference/environ/drone-repo-private/)  

`.Repo.SCM`  
String literal, provides the repository version control system. [DRONE_REPO_SCM](String literal, provides the repository version control system.)  

`.Build.Branch`  
String literal, provides the branch for the current build. [DRONE_BRANCH](https://docs.drone.io/reference/environ/drone-branch/)  

`.Build.Created`  
Unix timestamp, provides the date and time when the build was created in the system. [DRONE_BUILD_CREATED](https://docs.drone.io/reference/environ/drone-build-created/)  

`.Build.Event`  
build event type enumeration, one of `push`, `pull_request`, `tag`, `deployment`  

`.Build.Number`  
Integer value, provides the current build number. [DRONE_BUILD_NUMBER](https://docs.drone.io/reference/environ/drone-build-number/)  

`.Build.Started`  
Unix timestamp, provides the date and time when the build was started. [DRONE_BUILD_STARTED](https://docs.drone.io/reference/environ/drone-build-started/)  

`.Build.Status`  
build status type enumeration, either `success` or `failure` DRONE_BUILD_STATUS  

`.Build.Link`  
DRONE_BUILD_LINK  

`.Build.PullRequest`  
Integer value, provides the pull request number for the current build. This value is only set if the build event is of type pull request. [DRONE_PULL_REQUEST](https://docs.drone.io/reference/environ/drone-pull-request/)  

`.Build.SourceBranch`  
String literal, provides the source branch for a pull request. [DRONE_SOURCE_BRANCH](https://docs.drone.io/reference/environ/drone-source-branch/)  

`.Build.TargetBranch`  
String literal, provides the target branch for a pull request. [DRONE_TARGET_BRANCH](https://docs.drone.io/reference/environ/drone-target-branch/)  

`.Build.Tag`  
String literal, provides the tag name for the current build. This value is only set if the build event is of type tag. [DRONE_TAG](https://docs.drone.io/reference/environ/drone-tag/)  

`.Commit.Commit`  
String literal, provides the commit sha for the current build. [DRONE_COMMIT](https://docs.drone.io/reference/environ/drone-commit/)  
