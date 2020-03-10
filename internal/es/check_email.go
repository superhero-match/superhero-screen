/*
  Copyright (C) 2019 - 2020 MWSOFT
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
	"encoding/json"
	"fmt"

	"github.com/superhero-match/superhero-screen/internal/es/model"

	"gopkg.in/olivere/elastic.v7"
)

// CheckEmailExists checks if document of type superhero in superheros index exists
// and if so, if the superhero is blocked.
func (es *ES) CheckEmailExists(email string) (rsp *model.CheckEmailResponse, err error) {
	var result model.CheckEmailResponse

	q := elastic.NewTermQuery("email", email)

	fmt.Println()
	fmt.Printf("%+v", q)
	fmt.Println()

	searchResult, err := es.Client.Search().
		Index(es.Index).
		Query(q).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Printf("SearchResult: %+v", searchResult)

	fmt.Println()

	fmt.Println(searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			fmt.Printf("Hit: %+v", hit)
			var s model.Superhero

			err := json.Unmarshal(hit.Source, &s)
			if err != nil {
				return nil, err
			}

			fmt.Println()
			fmt.Printf("Superhero Unmarshalled: %+v", &s)
			fmt.Println()

			isRegistered := false

			if &s != nil {
				isRegistered = true
			}

			result = model.CheckEmailResponse{
				IsRegistered: isRegistered,
				Superhero:    &s,
			}
		}
	}

	return &result, nil
}
