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
$  docker-compose run tests gotestsum --format testname -- -run TestConceptListTestSuite/TestRegression ./...
Creating network "go-api-automation-example_default" with the default driver
Creating go-api-automation-example_tests_run ... done
go: downloading github.com/stretchr/testify v1.7.0
go: downloading github.com/go-resty/resty/v2 v2.6.0
go: downloading github.com/joho/godotenv v1.3.0
go: downloading github.com/davecgh/go-spew v1.1.0
go: downloading gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
EMPTY dto
EMPTY lib
PASS tests.TestConceptListTestSuite/TestRegression/Should_success_to_get_a_user (0.68s)
PASS tests.TestConceptListTestSuite/TestRegression (0.68s)
PASS tests.TestConceptListTestSuite (0.68s)
PASS tests

DONE 3 tests in 6.657s
```

### Reporting the results ###
Soon..