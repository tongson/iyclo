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
	L.SetMx(1024)
	glecho.Load(L, c)            // Loads _G.E (Echo context)
	glecho.LoadLogger(L, jlG)    // Loads _G.L (Logger)
	ll.LoadGlobalGo(L, "exec")   // _G.exec
	ll.LoadGlobalGo(L, "os")     // _G.os
	ll.LoadGlobalGo(L, "fs")     // _G.fs
	ll.LoadGlobalGo(L, "extend") // _G.extend
	ll.PreloadGo(L, "json")      // require("json")
	ll.PreloadGo(L, "ulid")      // require("ulid")
	ll.PreloadGo(L, "bitcask")   // require("bitcask")
	ll.PreloadGo(L, "crypto")    // require("crypto")
	ll.Preload(L)                // Load embedded Lua modules
	return ll.Main(L, ll.ReadFile(luaSrc, "src/main.lua"))
}
