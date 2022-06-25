package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/russross/blackfriday/v2"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func TestBlackfridayParser(t *testing.T) {
	f, err := os.Open("testdata/typical.md")
	require.NoError(t, err)

	raw, err := io.ReadAll(f)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	md := blackfriday.New()
	x := md.Parse(raw)
	fmt.Println(x.String())
}

func TestGoldmarkParser(t *testing.T) {
	f, err := os.Open("testdata/typical.md")
	require.NoError(t, err)

	raw, err := io.ReadAll(f)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	r := text.NewReader(raw)
	p := goldmark.DefaultParser()
	n := p.Parse(r)
	n.Dump(raw, 0)
}
