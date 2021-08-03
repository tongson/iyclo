package main

const versionStringG = "0.1.0 (Chill Hazelnut)"
const defaultLogG = "/var/log/iyclo.json"
const defaultDbG = "/var/lib/iyclo.db"
const defaultSocketG = "/var/run/iyclo.sock"

type actionT struct {
	version bool
	log     string
	db      string
	socket  string
}
