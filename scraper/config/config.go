package config

type UserConfig = struct {
	Output string
	All    bool
	Single bool
	Yes    bool
}

var Config UserConfig = UserConfig{}
