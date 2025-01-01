package elastic

import (
	"eattogether/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

type ElasticClient struct {
	Connection *elasticsearch.Client
}

// TODO подумать в строну инкапскуляции логики по методам. Возможно отдельная библиотечка???
func (e *ElasticClient) Search(index string, query string) (*models.ElasticResponse, error) {
	result, err := e.Connection.Search(
		e.Connection.Search.WithIndex(index),
		e.Connection.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		fmt.Printf("Error during es query: %v\n", err)
		return nil, err
	}

	defer result.Body.Close()

	var ElasticResponse models.ElasticResponse
	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Printf("Error while reading body: %v\n", err)
		return nil, err
	}

	err = json.Unmarshal(body, &ElasticResponse)
	if err != nil {
		fmt.Printf("Error while unmarshalling: %v\n", err)
		return nil, err
	}

	return &ElasticResponse, nil
}

// TODO добавить конфиг для гибкой конфигурации
func CreateElasticClient() (*ElasticClient, error) {
	connection, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Printf("Can't create to elastic: %v\n", err)
		return nil, err
	}

	return &ElasticClient{
		Connection: connection,
	}, nil
}
