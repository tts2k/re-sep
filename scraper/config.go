package main

type Config = struct {
	Output string
	All    bool
	Single bool
	Yes    bool
}

var config Config = Config{}
