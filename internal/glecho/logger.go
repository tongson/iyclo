package glecho

import (
	"github.com/rs/zerolog"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
	"strings"
)

const (
	LOGGER_TYPE = "logger{api}"
)

type loggerT struct {
	logger *zerolog.Logger
}

func loggerInit(logger zerolog.Logger) *loggerT {
	return &loggerT{
		&logger,
	}
}

func loggerCheck(L *lua.LState) *loggerT {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*loggerT); ok {
		return v
	} else {
		return nil
	}
}

func loggerPush(L *lua.LState, event *zerolog.Event) int {
	msg := L.CheckString(2)
	stubs := L.CheckTable(3)

	gostubs := make(map[string]interface{})
	err := gluamapper.Map(stubs, &gostubs)

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	for str, val := range gostubs {
		event.Interface(strings.ToLower(str), val)
	}
	event.Msg(msg)
	return 0
}

var loggerMethods = map[string]lua.LGFunction{
	"info": func(L *lua.LState) int {
		logger := loggerCheck(L)
		if logger == nil {
			L.Push(lua.LNil)
			L.Push(lua.LString("info: Invalid logger."))
			return 2
		}
		event := logger.logger.Info()
		return loggerPush(L, event)
	},
	"debug": func(L *lua.LState) int {
		logger := loggerCheck(L)
		if logger == nil {
			L.Push(lua.LNil)
			L.Push(lua.LString("info: Invalid logger."))
			return 2
		}

		event := logger.logger.Debug()
		return loggerPush(L, event)
	},
	"warn": func(L *lua.LState) int {
		logger := loggerCheck(L)
		if logger == nil {
			L.Push(lua.LNil)
			L.Push(lua.LString("info: Invalid logger."))
			return 2
		}
		event := logger.logger.Warn()
		return loggerPush(L, event)
	},
	"error": func(L *lua.LState) int {
		logger := loggerCheck(L)
		if logger == nil {
			L.Push(lua.LNil)
			L.Push(lua.LString("info: Invalid logger."))
			return 2
		}
		event := logger.logger.Error()
		return loggerPush(L, event)
	},
}

func Logger(L *lua.LState, jl zerolog.Logger) {
	loggerNew := func(L *lua.LState) int {
		ud := L.NewUserData()
		ud.Value = loggerInit(jl)
		L.SetMetatable(ud, L.GetTypeMetatable(LOGGER_TYPE))
		L.Push(ud)
		return 1
	}
	var loggerExports = map[string]lua.LGFunction{
		"new": loggerNew,
	}
	mt := L.NewTypeMetatable(LOGGER_TYPE)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), loggerMethods))
	L.SetGlobal("_ZEROLOG", L.SetFuncs(mt, loggerExports))
	newLogger := L.GetField(L.GetField(L.Get(lua.GlobalsIndex), "_ZEROLOG"), "new").(*lua.LFunction)
	L.Push(newLogger)
	L.Call(0, 1)
	L.SetGlobal("L", L.Get(L.GetTop()))
}
