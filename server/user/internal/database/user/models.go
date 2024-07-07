package database

type Margin = struct {
	Left  int32 `json:"left"`
	Right int32 `json:"right"`
}

type UserConfig = struct {
	Font     string `json:"font"`
	FontSize int32  `json:"fontSize"`
	Justify  bool   `json:"justify"`
	Margin   Margin `json:"margin"`
}
