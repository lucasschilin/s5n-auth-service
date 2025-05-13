package health

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RootHandler)
	handler.ServeHTTP(rr, req)

	statusExpected := http.StatusOK
	statusGot := rr.Code
	if statusGot != statusExpected {
		t.Errorf("Expected %v, got %v", statusExpected, statusGot)
	}

	bodyExpected := `{"detail":"Schily Users API online 🟢"}`
	bodyGot := strings.TrimSpace(rr.Body.String())

	fmt.Printf("Body got (raw): %q\n", bodyGot)
	fmt.Printf("Body expected (raw): %q\n", bodyExpected)

	if bodyGot != bodyExpected {
		t.Errorf("Expected %v, got %v", bodyExpected, bodyGot)
	}

}
