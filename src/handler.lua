return {
	get_containers = function(id)
		E:response():write_header(202)
		E:response():write("ok")
		E:response():flush()
	end,
}
