package database

type Margin = struct {
	Left  int `json:"left"`
	Right int `json:"right"`
}

type UserConfig = struct {
	Font     string `json:"font"`
	FontSize int    `json:"fontSize"`
	Jusitfy  bool   `json:"justify"`
	Margin   Margin `json:"margin"`
}
