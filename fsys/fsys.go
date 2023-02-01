package fsys

import (
	"embed"
)

//go:embed regexes/*
var FileSystem embed.FS
