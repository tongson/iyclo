package glecho

import (
	"github.com/yuin/gopher-lua"
)

func Variables(L *lua.LState, vars map[string]string) {
	t := L.NewTable()
	for k, v := range vars {
		L.SetField(t, k, lua.LString(v))
	}
	L.SetGlobal("V", t)
}
