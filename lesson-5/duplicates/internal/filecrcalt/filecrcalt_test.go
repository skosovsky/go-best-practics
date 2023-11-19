package filecrcalt

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_fileMD5(t *testing.T) {
	appFs := afero.NewOsFs()
	assert.Equal(t, "889b519803b9de3e338330eb365d0884", FileMD5("/Users/skosovsky/Downloads/Test/Test3/Test4/2.csv", appFs))
}
