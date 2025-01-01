package models

type ElasticHit struct {
	Source map[string]interface{} `json:"_source"`
}

type ElasticResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []ElasticHit `json:"hits"`
	} `json:"hits"`
}
