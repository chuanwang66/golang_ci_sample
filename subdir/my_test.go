package subdir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlabla (t *testing.T) {
	t.Log("blabla in tlog")

	v := "hehe"
	assert.Equal(t, "hehe", v)
}
