package blockchain

import (
	"fmt"

	"github.com/Suy56/ProofChain/chaincore/verify"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Deploy(txOpts *bind.TransactOpts, client *ethclient.Client) (string, *types.Transaction, error) {
    contractAddr, tx, _, err := verify.DeployVerify(txOpts, client)
    if err != nil {
        return "", nil, fmt.Errorf("Error deploying contract: %v", err)
    }
    return contractAddr.Hex(), tx, err
}