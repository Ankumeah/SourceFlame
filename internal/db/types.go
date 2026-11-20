package db

type Repo struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type EmailSender struct {
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Verified bool    `json:"verified"`
}
