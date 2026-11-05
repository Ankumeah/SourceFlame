package git

type File struct {
	Name string `json:"name"`
	Perm uint   `json:"perm"`
	Dir  bool   `json:"dir"`
}

type Ref struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
}
