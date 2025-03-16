package cengine

import (
	"bytes"
	"os"
	"path/filepath"
)

/*
Handle loading game data from data archives.
cengine does not preemptively load the entire game archive at once. only a few
parts are loaded on startup - namely, all Decks (part of a single collection),
Hands (optional), and the entrypoint file.
*/
func LoadData(path string) LoaderData {
	f, err := os.ReadFile(filepath.Join(".", path, "def.dkf"))
	if err != nil {
		panic(err)
	}
	r := bytes.Runes(f)
	d := newParser(r)
	return d.Parse()
}
