
package HttpHandlers

import (
	"net/http"
	"time"
)

type TimeHandler struct{}

func(s *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result, err := time.Now().MarshalText()
	if err != nil {
		WriteBadRequestResponse(w, "How the fuck")
	}
	WritePlaintextResponse(w, result)
}