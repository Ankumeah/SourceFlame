package git

type File struct {
	Name   string `json:"name"`
	Dir    bool   `json:"dir"`
	Submod bool   `json:"submod"`
}

type Ref struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
}
