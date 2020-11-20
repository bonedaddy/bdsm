package testenv

import (
	"github.com/bonedaddy/bdsm/bindings/datetime"
	"github.com/ethereum/go-ethereum/common"
)

func DeployDateTime(testenv *Testenv) (common.Address, *datetime.Datetime, error) {
	_, tx, contract, err := datetime.DeployDatetime(testenv.Auth, testenv)
	if err != nil {
		return common.Address{}, nil, err
	}
	addr, err := testenv.DoWaitDeployed(tx, "datetime")

	return addr, contract, nil
}
