package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	expected = `ok`
)

func TestRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/containers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	v := make(map[string]string)
	v["db"] = "test"
	h := handleHttp(jLog("/dev/null"), v)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
