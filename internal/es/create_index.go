package es

import (
	"context"
	"fmt"
)

// CreateIndex crates Elasticsearch index if index does not exist yet.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!! DO NOT CALL THIS METHOD THIS IS ONLY FOR LOCAL DEVELOPMENT PURPOSES !!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func(es *ES) CreateIndex() error {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := es.Client.IndexExists(es.Index).Do(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		mapping := `
		{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings":{
				"properties":{
					"superhero_id":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"email":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"name":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"superhero_name":{
						"type":"text",
						"store": true,
						"fielddata": true
					}, 
					"main_profile_pic_url":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"gender":{
						"type":"integer"
					},
					"looking_for_gender":{
						"type":"integer"
					},
					"age":{
						"type":"integer"
					},
					"looking_for_age_min":{
						"type":"integer"
					},
					"looking_for_age_max":{
						"type":"integer"
					},
					"looking_for_distance_max":{
						"type":"integer"
					},
					"distance_unit":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"location":{
						"type":"geo_point"
					},
					"birthday":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"country":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"city":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"superpower":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"account_type":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"is_deleted":{
						"type":"boolean"
					},
					"deleted_at":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"is_blocked":{
						"type":"boolean"
					},
					"blocked_at":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"updated_at":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"created_at":{
						"type":"text",
						"store": true,
						"fielddata": true
					}
				}
			}
		}
		`

		createIndex, err := es.Client.CreateIndex(es.Index).Body(mapping).Do(context.Background())
		if err != nil {
			return err
		}

		if !createIndex.Acknowledged {
			fmt.Println("create index is not acknowledged")
		}

		return nil
	}

	fmt.Println("index already exists")

	return nil
}
