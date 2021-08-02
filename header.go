package main

const versionString = "0.1.0 (Chill Hazelnut)"
const defaultLog = "/var/log/iyclo.json"
const defaultDb = "/var/lib/iyclo.bitcask"
const defaultSocket = "/var/run/iyclo.socket"

type actionT struct {
	version bool
	log     string
	db      string
	socket  string
}
