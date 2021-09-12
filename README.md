# README #
this repo contains example of integration tests developed in golang with testify & resty 

### Requirements ###
- golang 1.5 
- testify  `go get -u github.com/stretchr/testify`
- resty    `go get -u github.com/go-resty/resty/v2`
- docker

### How to run test ###
Build de the docker image 
```bash
$ docker-compose build
```
Inside of project run:
```bash
$ docker run -v ${pwd}:/test --rm=false go-api-automation-example gotestsum --format testname -- -run TestConceptListTestSuite/TestRegression ./...
EMPTY dto
EMPTY lib
PASS tests.TestConceptListTestSuite/TestRegression/Should_success_to_get_a_user (0.55s)
PASS tests.TestConceptListTestSuite/TestRegression (0.55s)
PASS tests.TestConceptListTestSuite (0.55s)
PASS tests

DONE 3 tests in 3.659s
```

### Reporting the results ###
Soon..