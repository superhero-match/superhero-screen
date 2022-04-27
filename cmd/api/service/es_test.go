package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/superhero-match/superhero-screen/internal/es"
)

func TestService_CheckEmailExists(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer testServer.Close()

	response := `{
							  "took" : 982,
							  "timed_out" : false,
							  "_shards" : {
								"total" : 5,
								"successful" : 5,
								"skipped" : 0,
								"failed" : 0
							  },
							  "hits" : {
								"total" : {
								  "value" : 10000,
								  "relation" : "gte"
								},
								"max_score" : 1.0,
								"hits" : [
								  {
									"_index" : "superhero",
									"_type" : "_doc",
									"_id" : "2ds34f6w-43f5-2344-dsf4-kf9ekw9fke9w",
									"_score" : 1.0,
									"_source" : {
									  "superhero_id" : "123456789",
									  "email" : "test@test.com",
                                      "name" : "Test Tester 1",
                                      "superhero_name": "Unit Tester 1",
                                      "main_profile_pic_url": "https://www.test.com/1",
                                      "profile_pics": [{
                                        "id": 1,
                                        "superhero_id": "123456789",
                                        "url": "https://www.test.com/2",
                                        "position": 1 
                                      }],
                                      "gender": 1,
                                      "looking_for_gender": 2,
                                      "age": 36,
                                      "looking_for_age_min": 25,
                                      "looking_for_age_max": 45,
                                      "looking_for_distance_max": 50,
                                      "distance_unit": "km",
                                      "location": {
                                        "lat": 0.123456789,
                                        "lon": 0.123456789
                                      },
									  "birthday" : "1985-04-26T12:00:00",
									  "country" : "Test Country",
									  "city" : "Test City",
									  "superpower" : "Unit Testing",
									  "account_type" : "FREE",
									  "created_at" : "2022-04-26T12:00:00"
									}
								  }
								]
							}
					}`

	mockClient, err := es.MockElasticSearchClient(testServer.URL, response)
	assert.NoError(t, err)

	mockEs := es.New(mockClient, "superhero")

	mockService := New(mockEs)

	_, err = mockService.CheckEmailExists("test@test.com")
	assert.NoError(t, err)
}

func TestService_CreateIndex(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer testServer.Close()

	response := `{"status": 200}`

	mockClient, err := es.MockElasticSearchClient(testServer.URL, response)
	assert.NoError(t, err)

	mockEs := es.New(mockClient, "superhero")

	mockService := New(mockEs)

	err = mockService.CreateIndex()
	assert.NoError(t, err)
}

func TestService_DeleteIndex(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer testServer.Close()

	response := `{"status": 200}`

	mockClient, err := es.MockElasticSearchClient(testServer.URL, response)
	assert.NoError(t, err)

	mockEs := es.New(mockClient, "superhero")

	mockService := New(mockEs)

	err = mockService.DeleteIndex()
	assert.NoError(t, err)
}
