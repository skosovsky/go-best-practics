package main

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fileMD5(t *testing.T) {
	assert.Equal(t, "889b519803b9de3e338330eb365d0884", fileMD5("/Users/skosovsky/Downloads/Test/Test3/Test4/2.csv"))
}

func Test_fileMD5Alt(t *testing.T) {
	var appFs = afero.NewOsFs()
	assert.Equal(t, "889b519803b9de3e338330eb365d0884", fileMD5Alt("/Users/skosovsky/Downloads/Test/Test3/Test4/2.csv", appFs))
}
