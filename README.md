### Open Source Project Health Monitor

Indicator of project health (i.e. is this project dead?)

#### Features

- [ ] Accept Github repo and pull down metrics 
    - [ ] Stars
    - [ ] Forks
    - [ ] Followers
    - [ ] Commit stats
    - [ ] Number of contributors
    - [ ] Issue stats
- [ ] Calculate aggregate "health" or "activity" score
- [ ] Persist scores (database)
- [ ] Generate "badge" with score

#### Admin/Ops

- [ ] Setup Github app credentials
- [ ] Setup Heroku app
- [ ] Travis CI
- [ ] Automatically list "production" builds as releases
- [ ] Good test coverage

#### Setup

1. Get code: `git clone https://github.com/140proof/OSS-health.git`
1. Get dependencies: `go get -t`
1. Build: `go build`
1. Run: `./OSS-health`

#### Resources

* [Github API](https://developer.github.com/v3/)
* [Github Go Package](https://github.com/google/go-github)
* [Go OAuth2 Package](https://github.com/golang/oauth2)
