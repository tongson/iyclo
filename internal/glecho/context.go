package glecho

import (
	lua "github.com/yuin/gopher-lua"
	"github.com/labstack/echo/v4"
)

func Context(L *lua.LState, c echo.Context) {
	t := L.NewTable()

	lCtx := L.NewTypeMetatable("context")
	L.SetField(lCtx, "__index", L.SetFuncs(L.NewTable(), API_LContext))
	t.RawSetH(lua.LString("context"), lCtx)

	lResponse := L.NewTypeMetatable("response")
	L.SetField(lResponse, "__index", L.SetFuncs(L.NewTable(), API_LResponse))
	t.RawSetH(lua.LString("response"), lResponse)

	lCookie := L.NewTypeMetatable("cookie")
	L.SetField(lCookie, "__index", L.SetFuncs(L.NewTable(), API_LCookie))
	t.RawSetH(lua.LString("cookie"), lCookie)

	lUrl := L.NewTypeMetatable("url")
	L.SetField(lUrl, "__index", L.SetFuncs(L.NewTable(), API_LUrl))
	t.RawSetH(lua.LString("url"), lUrl)

	lRequest := L.NewTypeMetatable("request")
	L.SetField(lRequest, "__index", L.SetFuncs(L.NewTable(), API_LRequest))
	t.RawSetH(lua.LString("request"), lRequest)

	t.RawSetH(lua.LString("__version__"), lua.LString(VERSION))

	L.SetGlobal("_ECHO", L.SetFuncs(t, API))
	newFromUD := L.GetField(L.GetField(L.Get(lua.GlobalsIndex), "_ECHO"), "newFromUD").(*lua.LFunction)
	L.Push(newFromUD)
	cud := L.NewUserData()
	cud.Value = c
	L.Push(cud)
	L.Call(1, 1)
	L.SetGlobal("E", L.Get(1))
}
