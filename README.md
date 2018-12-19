# drone-typetalk

Drone plugin for sending Typetalk notifications.

| Badge | Description |
| ------------- | ------------- |
| [![Build Status](https://travis-ci.org/is2ei/drone-slack.svg?branch=master)][travis] | Travis Build Status |
| [![Go Report Card](https://goreportcard.com/badge/github.com/is2ei/drone-typetalk)][goreport] | Go Report |

[travis]: https://travis-ci.org/is2ei/drone-slack
[goreport]: https://goreportcard.com/report/github.com/is2ei/drone-typetalk

## Testing the plugin

```
docker run --rm \
  -e TYPETALK_TOKEN=xxxxxxxx \
  -e TOPIC_ID=12345 \
  -e MESSAGE="hello world" \
  is2ei/drone-typetalk
```