package testenv

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeployDateTime(t *testing.T) {
	testenv, err := NewBlockchain(context.Background())
	require.NoError(t, err)
	_, _, err = DeployDateTime(testenv)
	require.NoError(t, err)
}
