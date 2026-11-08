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

type Commit struct {
	Hash    string `json:"hash"`
	Author  Author `json:"author"`
	Message string `json:"message"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
