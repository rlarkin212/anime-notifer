package util

import (
	"encoding/json"
	"net/http"
)

func UnmarshallResponseBody(res *http.Response, target interface{}) error {
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
