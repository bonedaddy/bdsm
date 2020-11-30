package testenv

import (
	"context"
	"math/big"
	"testing"

	"github.com/bonedaddy/bdsm/bindings/structassigntest"
	"github.com/bonedaddy/bdsm/bindings/testproxyfactory"
	"github.com/bonedaddy/bdsm/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestProxyFactory(t *testing.T) {
	// create an account for proxy deployment
	acct, _, err := utils.NewAccount()
	require.NoError(t, err)

	testenv, err := NewBlockchain(context.Background())
	require.NoError(t, err)

	require.NoError(t, testenv.SendETH(acct.From, utils.OneEthInWei))

	_, tx, contract, err := testproxyfactory.DeployTestproxyfactory(testenv.Auth, testenv)
	require.NoError(t, err)
	_, err = testenv.DoWaitDeployed(tx)
	require.NoError(t, err)
	tx, err = contract.DeployProxyContract(testenv.Auth, acct.From, nil)
	require.NoError(t, err)
	err = testenv.DoWaitMined(tx)
	require.NoError(t, err)
	num, err := contract.NumProxies(nil)
	require.NoError(t, err)
	require.Equal(t, int64(1), num.Int64())

	addr, err := contract.Proxies(nil, big.NewInt(0))
	require.NoError(t, err)
	require.NotEqual(t, common.HexToAddress(""), addr)

	scontract, err := structassigntest.NewStructassigntest(addr, testenv)
	require.NoError(t, err)

	tx, err = scontract.SetMany(testenv.Auth)
	require.NoError(t, err)
	testenv.DoWaitMined(tx)
}
