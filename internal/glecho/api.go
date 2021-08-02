package glecho

import (
	lua "github.com/yuin/gopher-lua"
	"github.com/labstack/echo/v4"
)

var VERSION string = "1.0.0"
var LNAME string = "echo"

var API = map[string]lua.LGFunction{
	"newFromUD": LContext_NewFromUD,
}

var API_LContext = map[string]lua.LGFunction{
	"is_tls":     LContext_IsTLS,
	"error":     LContext_HttpError,
	"no_content": LContext_NoContent,
	"string":    LContext_String,
	"html":      LContext_HTML,
	"blob":      LContext_Blob,
	"json":      LContext_Json,
	"jsonp":     LContext_Jsonp,
	"scheme":    LContext_Scheme,
	"host":      LContext_Host,
	"method":    LContext_Method,
	"param":     LContext_Param,
	"form_file":  LContext_FormFile,
	"form_value": LContext_FormValue,
	//"formParams":    LContext_FormParams,
	"multipart_form": LContext_MultipartForm,
	"get_cookie":     LContext_GetCookie,
	"get_cookies":    LContext_GetCookies,
	"response":      LContext_Response,
	"request":       LContext_Request,
	"real_ip":        LContext_RealIP,
}

var API_LRequest = map[string]lua.LGFunction{
	"read_body":   LRequest_ReadBody,
	"close_body":  LRequest_CloseBody,
	"host":       LRequest_Host,
	"method":     LRequest_Method,
	"url":        LRequest_URL,
	"uri":        LRequest_URI,
	"remoteaddr": LRequest_RemoteAddr,
	"get_header":  LRequest_GetHeader,
	"useragent":  LRequest_UserAgent,
}

var API_LResponse = map[string]lua.LGFunction{
	"set_header":   LResponse_SetHeader,
	"get_header":   LResponse_GetHeader,
	"del_header":   LResponse_DelHeader,
	"write_header": LResponse_WriteHeader,
	"write":       LResponse_Write,
	"flush":       LResponse_Flush,
	"committed":   LResponse_Committed,
	"size":        LResponse_Size,
	"status":      LResponse_Status,
}

var API_LMultipartForm = map[string]lua.LGFunction{
	"files":     LMultipartForm_Files,
	"remove_all": LMultipartForm_RemoveAll,
}

var API_LFormFile = map[string]lua.LGFunction{
	"read":     LFormFile_Read,
	"size":     LFormFile_Size,
	"filename": LFormFile_FileName,
}

var API_LCookie = map[string]lua.LGFunction{
	"table":    LCookie_Table,
	"domain":   LCookie_Domain,
	"secure":   LCookie_Secure,
	"httponly": LCookie_HttpOnly,
	"path":     LCookie_Path,
	"expires":  LCookie_Expires,
	"value":    LCookie_Value,
	"name":     LCookie_Name,
}

var API_LUrl = map[string]lua.LGFunction{
	"scheme":   LUrl_Scheme,
	"host":     LUrl_Host,
	"is_abs":    LUrl_IsAbs,
	"path":     LUrl_Path,
	"query":    LUrl_Query,
	"username": LUrl_Username,
	"password": LUrl_Password,
	"param":    LUrl_Param,
	"nopaque":  LUrl_Opaque,
}

func XLoader(L *lua.LState) int {
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

	L.SetFuncs(t, API)
	L.Push(t)
	return 1
}

func Load(L *lua.LState, c echo.Context) {
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
