package application

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandlerSuccessCase(t *testing.T) {
	jsonRequest := `{"expression": "3+(8*3)"}`
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(jsonRequest))
	w := httptest.NewRecorder()
	CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := "result: 27.000000"

	if  string(data) != expected {
		t.Errorf("Expected 27 but got %v", string(data))
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}
}
func TestCalcHandlerBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("wrong status code")
	}
}