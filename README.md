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

`.Repo.FullName`  
String literal, provides the full name of the repository. [DRONE_REPO](https://docs.drone.io/reference/environ/drone-repo/)  

`.Repo.Owner`  
repository owner DRONE_REPO_OWNER  

`.Repo.Name`  
repository name DRONE_REPO_NAME  

`.Build.Status`  
build status type enumeration, either `success` or `failure` DRONE_BUILD_STATUS  

`.Build.Event`  
build event type enumeration, one of `push`, `pull_request`, `tag`, `deployment`  

`.Build.Number`  
Integer value, provides the current build number. [DRONE_BUILD_NUMBER](https://docs.drone.io/reference/environ/drone-build-number/)  

