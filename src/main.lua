local crypto = require("crypto")
local reqid = crypto.fast_random()
local handler = require("handler")
local request = {}
request[1] = E:request():method()
request[2] = E:request():uri()
local signature = table.concat(request, "%")
-- Decision Table
-- (method, path)
local entrypoints = {
	["GET%/api/v1/containers"] = "get_containers"
}
local matched = entrypoints[signature]
if matched then
	L:info(reqid, { sig = signature, fn = matched })
	return handler[matched](reqid)
else
	E:response():write_header(404)
	E:response():flush()
end
