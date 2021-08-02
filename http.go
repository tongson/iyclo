package main

import (
	"embed"
	"net"
	"net/http"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	ll "github.com/tongson/LadyLua"
	glecho "github.com/tongson/iyclo/internal/glecho"
	lua "github.com/yuin/gopher-lua"
)

//go:embed src/*
var luaSrc embed.FS

func mainHttp(l net.Listener) {
	e := echo.New()
	e.Listener = l
	server := new(http.Server)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/*", handleHttp)
	e.Logger.Fatal(e.StartServer(server))
}

func handleHttp(c echo.Context) error {
	L := lua.NewState()
	defer L.Close()
	L.SetMx(256)
	glecho.Load(L, c) // Loads global E
	glecho.Logger(L, jlG) // Loads global L
	ll.Preload(L)     // Load embedded Lua modules
	return ll.Main(L, ll.ReadFile(luaSrc, "src/main.lua"))
}
