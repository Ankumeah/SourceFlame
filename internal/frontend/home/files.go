package home

import (
	_ "embed"
)

//go:embed index.html
var index []byte

//go:embed styles.css
var styles []byte

//go:embed script.js
var script []byte
