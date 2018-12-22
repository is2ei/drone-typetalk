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

## Parameters Reference

`typetalk_token`  
[Typetalk Token](https://developer.nulab-inc.com/docs/typetalk/#tttoken)

## Template Reference

`.Build.Number`  
Integer value, provides the current build number. [DRONE_BUILD_NUMBER](https://docs.drone.io/reference/environ/drone-build-number/)