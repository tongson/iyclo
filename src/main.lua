local tuple = require("tuple")
local handler = require("handler")
local request = tuple(E:request():method(), E:request():uri())
local signature = tostring(request)
-- Decision Table
-- (method, path)
local entrypoints = {
	["(GET, /api/v1/containers)"] = "get_containers"
}
local matched = entrypoints[signature]
if matched then
	return handler[matched]()
else
	E:response():write_header(404)
	E:response():flush()
end
