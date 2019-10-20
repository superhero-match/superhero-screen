package es

import (
	"context"
	"fmt"
)

// Ping the Elasticsearch server to make sure that ES is running.
func(es *ES) Ping() error {
	_, _, err := es.Client.Ping(fmt.Sprintf("http://%s:%s", es.Host, es.Port)).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
