package client

import (
	"path/filepath"
	"testing"

	. "github.com/Serjlee/devicedetector/parser"
	"gotest.tools/assert"
)

func TestMediaPlayerParse(t *testing.T) {
	ps := NewMediaPlayer(filepath.Join(dir, FixtureFileMediaPlayer))
	var list []*ClientFixture
	err := ReadYamlFileFromDisk(`fixtures/mediaplayer.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		assert.DeepEqual(t, item.ClientMatchResult, r)
	}
}
