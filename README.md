# README #
this repo contains example of integration tests developed in golang with testify & resty 

### Requirements ###
- golang 1.5 
- testify  `go get -u github.com/stretchr/testify`
- resty    `go get -u github.com/go-resty/resty/v2`
- docker

### How to run test ###
Inside of project run:
```bash
$ docker-compose up --build
Building tests
[+] Building 2.9s (6/6) FINISHED
 => [internal] load build definition from dockerfile                                                                                         0.1s 
 => => transferring dockerfile: 215B                                                                                                         0.0s 
 => [internal] load .dockerignore                                                                                                            0.1s 
 => => transferring context: 2B                                                                                                              0.0s 
 => [internal] load metadata for docker.io/library/golang:1.15-alpine                                                                        2.6s 
 => [1/2] FROM docker.io/library/golang:1.15-alpine@sha256:b58c367d52e46cdedc25ec9cd74cadb14ad65e8db75b25e5ec117cdb227aa264                  0.0s 
 => CACHED [2/2] RUN set -ex;     apk update;     apk add --no-cache git                                                                     0.0s 
 => exporting to image                                                                                                                       0.1s 
 => => exporting layers                                                                                                                      0.0s 
 => => writing image sha256:b76ffc5f7a9e5a02c48cb7b598cc40d588b5a787fbcab5381e07fe920ef9562a                                                 0.0s 
 => => naming to docker.io/library/go-api-automation-example                                                                                 0.0s 

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
Recreating go-api-automation-example ... done
Attaching to go-api-automation-example
go-api-automation-example | go: downloading github.com/stretchr/testify v1.7.0
go-api-automation-example | go: downloading github.com/go-resty/resty/v2 v2.6.0
go-api-automation-example | go: downloading github.com/joho/godotenv v1.3.0
go-api-automation-example | go: downloading github.com/pmezard/go-difflib v1.0.0
go-api-automation-example | go: downloading github.com/davecgh/go-spew v1.1.0
go-api-automation-example | go: downloading gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
go-api-automation-example | go: downloading golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
go-api-automation-example | ?           go-api-integration-tests/dto    [no test files]
go-api-automation-example | ?           go-api-integration-tests/lib    [no test files]
go-api-automation-example | === RUN   TestConceptListTestSuite
go-api-automation-example | === RUN   TestConceptListTestSuite/TestGetUser
go-api-automation-example | === RUN   TestConceptListTestSuite/TestGetUsers
go-api-automation-example | === RUN   TestConceptListTestSuite/TestGetUsers/Should_success_to_get_all_users
go-api-automation-example | --- PASS: TestConceptListTestSuite (3.76s)
go-api-automation-example |     --- PASS: TestConceptListTestSuite/TestGetUser (3.45s)
go-api-automation-example |         --- PASS: TestConceptListTestSuite/TestGetUser/Should_success_to_get_a_user (3.45s)
go-api-automation-example |     --- PASS: TestConceptListTestSuite/TestGetUsers (0.31s)
go-api-automation-example |         --- PASS: TestConceptListTestSuite/TestGetUsers/Should_success_to_get_all_users (0.31s)
go-api-automation-example | PASS
go-api-automation-example | ok          go-api-integration-tests/tests  3.774s
```

### Reporting the results ###
Soon..