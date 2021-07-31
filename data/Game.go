package data

type Game struct {
	Id      int     `json:"id"`
	Ruleset Ruleset `json:"ruleset"`
	Timeout int     `json:"timeout"`
}

type Ruleset struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}