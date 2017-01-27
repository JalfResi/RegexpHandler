package regexphandler

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestRegexpHandler(t *testing.T) {
	// Create a request to pass to our handler. We dont have any query parameters
	// for now, so we'll pass 'nil' as the third parameter
	req, err := http.NewRequest("GET", "/Ben", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()
	handler := RegexpHandler{}
	handler.HandleFunc(regexp.MustCompile("/[b|B]en"), func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
}
