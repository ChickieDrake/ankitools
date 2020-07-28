package convert

import "encoding/json"

type request_body struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
}

var marshal_function = json.Marshal

func ConvertToRequestMessage(action string) (message string, err error) {
	s := &request_body{Action: action, Version: 6}
	a, err := marshal_function(s)
	if err != nil {
		return
	}
	return string(a), err
}
