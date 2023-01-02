package transpose

import (
	"eel/eel"
	"encoding/json"

	"github.com/broothie/qst"
)

type Stats struct {
	Count int32
	Size  int32
	Time  int32
}

type QueryResponse struct {
	Status   string      `json:"status"`
	ResStats Stats       `json:"stats"`
	Results  interface{} `json:"results"`
}

func ExecuteQuery(query string, cfg *eel.EelConfig) (*QueryResponse, error) {
	res, err := qst.Post(
		"https://api.transpose.io/sql",
		qst.Header("X-API-KEY", cfg.TransposeToken),
		qst.Header("Content-Type", "application/json"),
		qst.BodyJSON(
			map[string]interface{}{
				"query": query,
			},
		),
	)

	if err != nil {
		return nil, err
	}

	var r QueryResponse

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
