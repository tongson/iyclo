package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	expected = `test`
)

func TestRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/endpoint", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleHttp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
