package gofalcon

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "reflect"
    "testing"
)

// TestGoFalcon tests the GoFalcon framework
func TestGoFalcon(t *testing.T) {
    router := NewServer()
    router.Handle("GET", "/ping", func(c *Context) {
        c.JSON(http.StatusOK, M{"message": "pong"})
    })

    req, _ := http.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }

    var expected, actual M
    expected = M{"message": "pong"}
    if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("Expected body %v, got %v", expected, actual)
    }
}
