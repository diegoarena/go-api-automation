# README #
this repo contains example of integration tests developed in golang with testify & resty 

### Requirements ###
- golang 1.5 
- testify  `go get -u github.com/stretchr/testify`
- resty    `go get -u github.com/go-resty/resty/v2`
- docker

### How to run test ###
Build and run the docker container  
```bash
$ docker-compose up --build -d
```
Run the tests
```bash
$ docker exec -ti go-api-automation-example gotestsum --format testname -- -run TestConceptListTestSuite/TestRegression ./... -count=1
EMPTY dto
EMPTY lib
PASS tests.TestConceptListTestSuite/TestRegression/Should_success_to_get_a_user (0.55s)
PASS tests.TestConceptListTestSuite/TestRegression (0.55s)
PASS tests.TestConceptListTestSuite (0.55s)
PASS tests

DONE 3 tests in 0.993s
```

### Reporting the results ###
Soon..