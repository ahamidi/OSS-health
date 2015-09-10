![Project Health](https://oss-health.herokuapp.com/140proof/oss-health/badge.svg)

### Open Source Project Health Monitor

Indicator of project health (i.e. is this project dead?)

#### Features

- [ ] Accept Github repo and pull down metrics 
    - [x] Stars
    - [x] Forks
    - [x] Followers
    - [x] Commit stats
    - [x] Number of contributors
    - [ ] Issue stats
- [ ] Calculate aggregate "health" or "activity" score
- [ ] Persist scores (database)
- [x] Generate "badge" with score

#### Admin/Ops

- [x] Setup Github app credentials
- [x] Setup Heroku app
- [ ] Travis CI
- [ ] Automatically list "production" builds as releases
- [ ] Good test coverage

#### Setup

1. Get code: `git clone https://github.com/140proof/OSS-health.git`
1. Get dependencies: `go get -t`
1. Build: `go build`
1. Run: `./OSS-health`

#### Deploy (Heroku)

1. [Install Heroku Toolbelt](https://toolbelt.heroku.com/)
1. Configure Heroku: `heroku login`
1. Add Heroku App as Git Remote: `heroku git:remote -a oss-health`
1. Deploy: `git push heroku master`

*Notes*: In order to deploy a different branch you must use `git push heroku <other_branch>:master`

#### Resources

* [Github API](https://developer.github.com/v3/)
* [Github Go Package](https://github.com/google/go-github)
* [Go OAuth2 Package](https://github.com/golang/oauth2)
