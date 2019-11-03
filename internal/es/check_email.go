package es

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/superhero-screen/internal/es/model"

	"gopkg.in/olivere/elastic.v7"
)

// CheckEmailExists checks if document of type superhero in superheros index exists
// and if so, if the superhero is blocked.
func (es *ES) CheckEmailExists(email string) (rsp *model.CheckEmailResponse, err error) {
	var result model.CheckEmailResponse

	fmt.Println(email)
	fmt.Println(es.Index)

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
				IsDeleted:    s.IsDeleted,
				IsBlocked:    s.IsBlocked,
				Superhero:    &s,
			}
		}
	}

	return &result, nil
}
