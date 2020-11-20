package testenv

import (
	"context"
	"testing"

	"github.com/bonedaddy/bdsm/utils"
	"github.com/stretchr/testify/require"
)

func TestTestenv(t *testing.T) {
	testenv, err := NewBlockchain(context.Background())
	require.NoError(t, err)
	testenv.SendETH(testenv.Auth.From, utils.OneEthInWei)
}
