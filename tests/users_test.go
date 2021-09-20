package tests

import (
	"fmt"
	dto "go-api-integration-tests/dto"
	"go-api-integration-tests/lib"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//testify suite
type UsersTestSuite struct {
	suite.Suite
	ApiClient   *resty.Client
	EndpointUrl string
}

// it runs after suite starts running
func (suite *UsersTestSuite) SetupSuite() {
	//instanciate rest client
	suite.ApiClient = resty.New()
	suite.EndpointUrl = lib.GetEnvVariable("BASE_URL") + "/users"
	//suite.EndpointUrl = "https://jsonplaceholder.typicode.com/users"
}

// test runner (needed to run testify with go testing lib)
func TestConceptListTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}

func (suite *UsersTestSuite) TestSmoke() {
	suite.T().Run("Should success to get all users", func(t *testing.T) {
		//response dto
		users := []dto.Users{}
		//resty request
		resp, err := suite.ApiClient.
			R().
			SetResult(&users).
			SetError(&users).
			Get(suite.EndpointUrl)

		//assert success request
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
		assert.NotNil(t, users)

		//assert response content
		for _, user := range users {
			assert.NotEmpty(t, user.Id)
			assert.NotEmpty(t, user.Name)
			assert.NotEmpty(t, user.Email)
			assert.NotEmpty(t, user.Username)
			assert.NotNil(t, user.Address)
			assert.NotEmpty(t, user.Company, "Company is not empty")
		}
	})

}

func (suite *UsersTestSuite) TestRegression() {
	suite.T().Run("Should success to get a user", func(t *testing.T) {
		//response dto
		user := dto.Users{}
		userId := fmt.Sprintf("%s%d", "/", lib.GetRandomNumber(1, 10))
		//resty request
		resp, err := suite.ApiClient.
			R().
			SetResult(&user).
			SetError(&user).
			Get(suite.EndpointUrl + userId)

		//assert success request
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
		assert.NotNil(t, user)

		//assert response content
		assert.NotEmpty(t, user.Id)
		assert.NotEmpty(t, user.Name)
		assert.NotEmpty(t, user.Email)
		assert.NotEmpty(t, user.Username)
		assert.NotNil(t, user.Address)
		assert.NotEmpty(t, user.Company)
	})
}
