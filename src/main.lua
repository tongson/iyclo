local tuple = require("tuple")
local crypto = require("crypto")
local ID = crypto.fast_random()
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
	L:info(ID, { sig = signature, fn = matched })
	return handler[matched](ID)
else
	E:response():write_header(404)
	E:response():flush()
end
