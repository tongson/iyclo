package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/rs/zerolog"
)

var t_variables map[string]string
var t_logger zerolog.Logger

func init() {
	t_variables = make(map[string]string)
	t_variables["db"] = "/tmp"
	t_logger = jLog("/dev/null")
}

func TestRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/containers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handleHttp(t_logger, t_variables)
	if assert.NoError(t, h(c)) {
		expected := `ok`
		assert.Equal(t, http.StatusAccepted, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/X", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handleHttp(t_logger, t_variables)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
