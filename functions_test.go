package tml

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSprintf(t *testing.T) {
	output := Sprintf("<red>hello</red>")
	assert.Equal(t, "\x1b[0m\x1b[31mhello\x1b[39m\x1b[0m", output)
}

func TestSprintfWithArgs(t *testing.T) {
	output := Sprintf("<red>%s</red>", "world")
	assert.Equal(t, "\x1b[0m\x1b[31mworld\x1b[39m\x1b[0m", output)
}

func TestFprintf(t *testing.T) {
	var buf bytes.Buffer
	err := Fprintf(&buf, "<red>hello</red>")
	require.NoError(t, err)
	assert.Equal(t, "\x1b[0m\x1b[31mhello\x1b[39m\x1b[0m", buf.String())
}

func TestFprintfWithArgs(t *testing.T) {
	var buf bytes.Buffer
	err := Fprintf(&buf, "<red>%s</red>", "world")
	require.NoError(t, err)
	assert.Equal(t, "\x1b[0m\x1b[31mworld\x1b[39m\x1b[0m", buf.String())
}

func TestFprintln(t *testing.T) {
	var buf bytes.Buffer
	Fprintln(&buf, "<red>hello</red>")
	assert.Equal(t, "\x1b[0m\x1b[31mhello\x1b[39m\x1b[0m\n", buf.String())
}

func TestFprintlnAppliesColorTags(t *testing.T) {
	var buf bytes.Buffer
	Fprintln(&buf, "<red>red text</red>")
	// Ensure ANSI codes are present, not the raw tag text
	assert.Contains(t, buf.String(), "\x1b[31m")
	assert.NotContains(t, buf.String(), "<red>")
}
