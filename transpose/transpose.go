package transpose

import (
	"eel/eel"
	"encoding/json"
	"errors"

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
		qst.Header("X-Api-Key", cfg.TransposeToken),
		qst.BodyJSON(
			map[string]interface{}{
				"sql": query,
			},
		),
	)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("non-200 status code returned")
	}

	var r QueryResponse

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
