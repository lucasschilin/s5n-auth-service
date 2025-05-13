package root_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lucasschilin/schily-users-api/internal/root"
)

func TestGETRootRequest(t *testing.T) {
	// Arrange
	statusExpected := http.StatusOK
	bodyExpected := `{"detail":"Schily Users API healthed and online 🟢"}`

	// Act
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(root.RootHandler)

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(w, r)

	// Assert
	if w.Code != statusExpected {
		t.Errorf("Expected %v, got %v", statusExpected, w.Code)
	}

	bodyGot := strings.TrimSpace(w.Body.String())
	if bodyGot != bodyExpected {
		t.Errorf("Expected %v, got %v", bodyExpected, bodyGot)
	}

}
