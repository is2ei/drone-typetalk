# drone-typetalk

[![Join the chat at https://gitter.im/is2ei/drone-typetalk](https://badges.gitter.im/is2ei/drone-typetalk.svg)](https://gitter.im/is2ei/drone-typetalk?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Drone plugin for sending Typetalk notifications.

| Badge | Description |
| ------------- | ------------- |
| [![Build Status](https://travis-ci.org/is2ei/drone-slack.svg?branch=master)][travis] | Travis Build Status |
| [![Build Status](https://cloud.drone.io/api/badges/is2ei/drone-typetalk/status.svg)][drone] | Drone Build Status |
| [![Go Report Card](https://goreportcard.com/badge/github.com/is2ei/drone-typetalk)][goreport] | Go Report |
| [![Join the chat at https://gitter.im/is2ei/drone-typetalk](https://badges.gitter.im/is2ei/drone-typetalk.svg)][gitter] | Gitter chat |

[travis]: https://travis-ci.org/is2ei/drone-slack
[drone]: https://cloud.drone.io/is2ei/drone-typetalk
[goreport]: https://goreportcard.com/report/github.com/is2ei/drone-typetalk
[gitter]: https://gitter.im/is2ei/drone-typetalk?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge

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

Example configuration with a custom message template:

```
steps:
- name: typetalk
  image: is2ei/typetalk
  settings:
    typetalk_token: xxxxxxxx
    topic_id: 12345
  template: >
    {{#success build.status}}
      build {{build.number}} succeeded. Good job.
    {{else}}
      build {{build.number}} failed. Fix me please.
    {{/success}}
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

## Template Reference

`drone.isDrone`  
Boolean value, indicates the runtime environment is Drone. [DRONE](https://docs.drone.io/reference/environ/drone/)  

`drone.hostname`  
String literal, provides the Drone agent hostname. [DRONE_MACHINE](https://docs.drone.io/reference/environ/drone-machine/)  

`drone.remoteURL`  
[DRONE_REMOTE_URL](https://docs.drone.io/reference/environ/drone-remote-url/)  

`system.host`  
String literal, provides the Drone server hostname. [DRONE_SYSTEM_HOST](https://docs.drone.io/reference/environ/drone-system-host/)  

`system.hostname`  
String literal, provides the Drone server hostname. [DRONE_SYSTEM_HOSTNAME](https://docs.drone.io/reference/environ/drone-system-hostname/)  

`system.version`  
String literal, provides the Drone server version. [DRONE_SYSTEM_VERSION](https://docs.drone.io/reference/environ/drone-system-version/)  

`runner.host`  
String literal, provides the Drone agent hostname. [DRONE_RUNNER_HOST](https://docs.drone.io/reference/environ/drone-runner-host/)  

`runner.hostname`  
String literal, provides the Drone agent hostname. [DRONE_RUNNER_HOSTNAME](https://docs.drone.io/reference/environ/drone-runner-hostname/)  

`runner.platform`  
String literal, provides the Drone agent os and architecture. [DRONE_RUNNER_PLATFORM](https://docs.drone.io/reference/environ/drone-runner-platform/)  

`runner.label`  
[DRONE_RUNNER_LABEL](https://docs.drone.io/reference/environ/drone-runner-label/)  

`git.HTTPURL`  
String literal, provides the repository git+http url. [DRONE_GIT_HTTP_URL](https://docs.drone.io/reference/environ/drone-git-http-url/)  

`git.SSHURL`  
String literal, provides the repository git+ssh url. [DRONE_GIT_SSH_URL](https://docs.drone.io/reference/environ/drone-git-ssh-url/)  

`repo.fullName`  
String literal, provides the full name of the repository. [DRONE_REPO](https://docs.drone.io/reference/environ/drone-repo/)  

`repo.owner`  
repository owner DRONE_REPO_OWNER  

`repo.name`  
String literal, provides the repository name. [DRONE_REPO_NAME](https://docs.drone.io/reference/environ/drone-repo-name/)  

`repo.branch`  
String literal, provides the default repository branch (e.g. master). [DRONE_REPO_BRANCH](https://docs.drone.io/reference/environ/drone-repo-branch/)  

`repo.link`  
String literal, provides the repository http link. [DRONE_REPO_LINK](https://docs.drone.io/reference/environ/drone-repo-link/)  

`repo.nameSpace`  
String literal, provides the repository namespace (e.g. account owner) [DRONE_REPO_NAMESPACE](https://docs.drone.io/reference/environ/drone-repo-namespace/)  

`repo.private`  
Boolean value, indicates the repository is public or private. [DRONE_REPO_PRIVATE](https://docs.drone.io/reference/environ/drone-repo-private/)  

`repo.SCM`  
String literal, provides the repository version control system. [DRONE_REPO_SCM](String literal, provides the repository version control system.)  

`build.branch`  
String literal, provides the branch for the current build. [DRONE_BRANCH](https://docs.drone.io/reference/environ/drone-branch/)  

`build.created`  
Unix timestamp, provides the date and time when the build was created in the system. [DRONE_BUILD_CREATED](https://docs.drone.io/reference/environ/drone-build-created/)  

`build.event`  
build event type enumeration, one of `push`, `pull_request`, `tag`, `deployment`  

`build.number`  
Integer value, provides the current build number. [DRONE_BUILD_NUMBER](https://docs.drone.io/reference/environ/drone-build-number/)  

`build.started`  
Unix timestamp, provides the date and time when the build was started. [DRONE_BUILD_STARTED](https://docs.drone.io/reference/environ/drone-build-started/)  

`build.status`  
build status type enumeration, either `success` or `failure` DRONE_BUILD_STATUS  

`build.link`  
DRONE_BUILD_LINK  

`build.pullRequest`  
Integer value, provides the pull request number for the current build. This value is only set if the build event is of type pull request. [DRONE_PULL_REQUEST](https://docs.drone.io/reference/environ/drone-pull-request/)  

`build.sourceBranch`  
String literal, provides the source branch for a pull request. [DRONE_SOURCE_BRANCH](https://docs.drone.io/reference/environ/drone-source-branch/)  

`build.targetBranch`  
String literal, provides the target branch for a pull request. [DRONE_TARGET_BRANCH](https://docs.drone.io/reference/environ/drone-target-branch/)  

`build.tag`  
String literal, provides the tag name for the current build. This value is only set if the build event is of type tag. [DRONE_TAG](https://docs.drone.io/reference/environ/drone-tag/)  

`commit.commit`  
String literal, provides the commit sha for the current build. [DRONE_COMMIT](https://docs.drone.io/reference/environ/drone-commit/)  

`commit.message`  
String literal, provides the commit message for the current build. [DRONE_COMMIT_MESSAGE](https://docs.drone.io/reference/environ/drone-commit-message/)  

`commit.after`  
String literal, provides the commit sha for the current build. [DRONE_COMMIT_AFTER](https://docs.drone.io/reference/environ/drone-commit-after/)  

`commit.author`  
String literal, provides the author username for the current commit. [DRONE_COMMIT_AUTHOR](https://docs.drone.io/reference/environ/drone-commit-author/)  

`commit.authorAvatar`  
String literal, provides the author avatar for the current commit. [DRONE_COMMIT_AUTHOR_AVATAR](https://docs.drone.io/reference/environ/drone-commit-author-avatar/)  

`commit.authorEmail`  
String literal, provides the author email for the current commit. [DRONE_COMMIT_AUTHOR_EMAIL](https://docs.drone.io/reference/environ/drone-commit-author-email/)  

`commit.authorName`  
String literal, provides the author name for the current commit. [DRONE_COMMIT_AUTHOR_NAME](https://docs.drone.io/reference/environ/drone-commit-author-name/)  

`commit.before`  
String literal, provides the parent commit sha for the current build. [DRONE_COMMIT_BEFORE](https://docs.drone.io/reference/environ/drone-commit-before/)  

`commit.branch`  
String literal, provides the branch for the current build. [DRONE_COMMIT_BRANCH](https://docs.drone.io/reference/environ/drone-commit-branch/)  

`commit.link`  
String literal, provides the http link to the current commit in the remote source code management system (e.g. GitHub). [DRONE_COMMIT_LINK](https://docs.drone.io/reference/environ/drone-commit-link/)

`commit.SHA`  
String literal, provides the commit sha for the current build. [DRONE_COMMIT_SHA](https://docs.drone.io/reference/environ/drone-commit-sha/)  

`commit.ref`  
String literal, provides the reference for the current build. [DRONE_COMMIT_REF](https://docs.drone.io/reference/environ/drone-commit-ref/)  
