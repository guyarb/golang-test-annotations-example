package foo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_foo_failed(t *testing.T) {
	require.Equal(t, 0, 3)
}

func Test_goo_succeeded(t *testing.T) {
	require.Equal(t, 0, 0)
}
