package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/russross/blackfriday/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/yuin/goldmark"
)

func TestBlackfridayParser(t *testing.T) {
	f, err := os.Open("testdata/typical.yaml")
	require.NoError(t, err)

	raw, err := io.ReadAll(f)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	md := blackfriday.New()
	x := md.Parse(raw)
	fmt.Println(x.String())
}

func Test(t *testing.T) {
	assert.True(t, true)
}
