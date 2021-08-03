package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	vars := make(map[string]string)
	vars["db"] = "/tmp"
	jl := jLog("/dev/null")
	e := echo.New()
	expected := `ok`

	req := httptest.NewRequest(http.MethodGet, "/api/v1/containers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handleHttp(jl, vars)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
