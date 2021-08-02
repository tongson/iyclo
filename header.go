package main

const versionString = "0.1.0 (Chill Hazelnut)"
const defaultLog = "/var/log/iyclo.json"
const defaultDb = "/var/lib/iyclo.bitcask"

type actionT struct {
	version bool
	log     string
	db      string
}
