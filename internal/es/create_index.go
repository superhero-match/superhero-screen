/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package es

import (
	"context"
	"fmt"
)

// CreateIndex crates Elasticsearch index if index does not exist yet.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!! DO NOT CALL THIS METHOD THIS IS ONLY FOR LOCAL DEVELOPMENT PURPOSES !!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (es *es) CreateIndex() error {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := es.Client.IndexExists(es.Index).Do(context.Background())
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("index already exists")
		return nil
	}

	mapping := `
		{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings":{
				"properties":{
					"superhero_id":{
						"type":"keyword",
						"store": true
					},
					"email":{
						"type":"keyword",
						"store": true
					},
					"name":{
						"type":"keyword",
						"store": true
					},
					"superhero_name":{
						"type":"keyword",
						"store": true
					}, 
					"main_profile_pic_url":{
						"type":"keyword",
						"store": true
					},
					"profile_pics":{
						"type": "nested"
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
						"type":"keyword",
						"store": true
					},
					"location":{
						"type":"geo_point"
					},
					"birthday":{
						"type":"date"
					},
					"country":{
						"type":"keyword",
						"store": true
					},
					"city":{
						"type":"keyword",
						"store": true
					},
					"superpower":{
						"type":"keyword",
						"store": true
					},
					"account_type":{
						"type":"keyword",
						"store": true
					},
					"created_at":{
						"type":"date"
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
