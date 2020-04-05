package HttpHandlers

import (
	"encoding/json"
	"net/http"
)

const JSON_HEADER = "application/json"
const PLAINTEXT_HEADER = "text/plain"
const CONTENT_TYPE_HEADER = "Content-Type"

func WriteJsonResponse(w http.ResponseWriter, v interface{}) error {
	w.Header().Set(CONTENT_TYPE_HEADER, JSON_HEADER)
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func WritePlaintextResponse(w http.ResponseWriter, body []byte) {
	w.Header().Set(CONTENT_TYPE_HEADER, PLAINTEXT_HEADER)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func WriteBadRequestResponse(w http.ResponseWriter, err string) {
	w.Header().Set(CONTENT_TYPE_HEADER, PLAINTEXT_HEADER)
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(err))
}
