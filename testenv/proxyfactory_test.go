package testenv

import (
	"context"
	"math/big"
	"testing"

	"github.com/bonedaddy/bdsm/bindings/testproxyfactory"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestProxyFactory(t *testing.T) {
	testenv, err := NewBlockchain(context.Background())
	require.NoError(t, err)
	_, tx, contract, err := testproxyfactory.DeployTestproxyfactory(testenv.Auth, testenv)
	require.NoError(t, err)
	_, err = testenv.DoWaitDeployed(tx)
	require.NoError(t, err)
	tx, err = contract.DeployProxyContract(testenv.Auth, testenv.Auth.From, nil)
	require.NoError(t, err)
	err = testenv.DoWaitMined(tx)
	require.NoError(t, err)
	num, err := contract.NumProxies(nil)
	require.NoError(t, err)
	require.Equal(t, int64(1), num.Int64())

	addr, err := contract.Proxies(nil, big.NewInt(0))
	require.NoError(t, err)
	require.NotEqual(t, common.HexToAddress(""), addr)
}
