package repo

import (
	_ "embed"
)

//go:embed index.html
var index []byte

//go:embed styles.css
var styles []byte

//go:embed script.js
var script []byte

//go:embed emails.html
var emails []byte

//go:embed emails_script.js
var emailsScript []byte
