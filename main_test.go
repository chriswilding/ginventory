package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://example/", nil)
	w := httptest.NewRecorder()

	Index(w, req)

	res := w.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.Equal(t, body, []uint8("hello\n"))
}
