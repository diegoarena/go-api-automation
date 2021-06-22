# README #
this repo contains example of integration tests developed in golang with testify & resty 

### Requirements ###
- golang 1.5 
- testify  `go get -u github.com/stretchr/testify`
- resty    `go get -u github.com/go-resty/resty/v2`

### How to run test ###
Inside of project run:
```bash
$ go test -v ./...
?       go-api-integration-tests        [no test files]
?       go-api-integration-tests/dto    [no test files]
?       go-api-integration-tests/dto/accounts   [no test files]
?       go-api-integration-tests/dto/commons    [no test files]
?       go-api-integration-tests/lib    [no test files]
=== RUN   TestConceptListTestSuite
=== RUN   TestConceptListTestSuite/TestGetUser
=== RUN   TestConceptListTestSuite/TestGetUser/Should_success_to_get_a_user
=== RUN   TestConceptListTestSuite/TestGetUsers
=== RUN   TestConceptListTestSuite/TestGetUsers/Should_success_to_get_all_users
--- PASS: TestConceptListTestSuite (0.85s)
    --- PASS: TestConceptListTestSuite/TestGetUser (0.67s)
        --- PASS: TestConceptListTestSuite/TestGetUser/Should_success_to_get_a_user (0.67s)
    --- PASS: TestConceptListTestSuite/TestGetUsers (0.19s)
        --- PASS: TestConceptListTestSuite/TestGetUsers/Should_success_to_get_all_users (0.19s)
PASS
ok      go-api-integration-tests/tests  0.859s
```

### Reporting the results ###
Soon..