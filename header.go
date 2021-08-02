package main

const versionString = "0.1.0 (Chill Hazelnut)"
const defaultLog = "/var/log/iyclo.json"

type actionT struct {
	version bool
	log     string
}
