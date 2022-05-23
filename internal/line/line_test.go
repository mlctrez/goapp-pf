package line

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParts(t *testing.T) {
	req := require.New(t)
	req.True(true)

	var parts []string

	parts = Parts(120, "Abcdefg")
	req.Equal([]string{"Abcdefg"}, parts)

	parts = Parts(5, "Something else")
	req.Equal([]string{"Something", "else"}, parts)

	parts = Parts(20, "Something else that is much longer and requires splitting.")
	req.Equal([]string{"Something else that is", "much longer and requires", "splitting."}, parts)

}
