package main

import (
	"embed"
	"net"
	"net/http"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	zerolog "github.com/rs/zerolog"
	ll "github.com/tongson/LadyLua"
	glecho "github.com/tongson/iyclo/internal/glecho"
	lua "github.com/yuin/gopher-lua"
)

//go:embed src/*
var luaSrc embed.FS

func mainHttp(l net.Listener, jl zerolog.Logger, v map[string]string) {
	e := echo.New()
	e.Listener = l
	server := new(http.Server)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/*", handleHttp(jl, v))
	e.Logger.Fatal(e.StartServer(server))
}

func handleHttp(jl zerolog.Logger, v map[string]string) echo.HandlerFunc {
	return func(c echo.Context) error {
		L := lua.NewState()
		defer L.Close()
		L.SetMx(1024)
		glecho.Context(L, c)       // _G.E (Echo context)
		glecho.Logger(L, jl)       // _G.L (Logger)
		glecho.Variables(L, v)     // _G.V (Configuration variables)
		ll.GlobalGo(L, "exec")     // _G.exec
		ll.GlobalGo(L, "os")       // _G.os
		ll.GlobalGo(L, "fs")       // _G.fs
		ll.GlobalGo(L, "extend")   // _G.extend
		ll.PreloadGo(L, "json")    // require("json")
		ll.PreloadGo(L, "ulid")    // require("ulid")
		ll.PreloadGo(L, "bitcask") // require("bitcask")
		ll.PreloadGo(L, "crypto")  // require("crypto")
		ll.Preload(L)              // Embedded plain Lua modules
		return ll.Main(L, ll.ReadFile(luaSrc, "src/main.lua"))
	}
}
