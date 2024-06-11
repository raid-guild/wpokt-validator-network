package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dan13ram/wpokt-oracle/ethereum/autogen"
	"github.com/dan13ram/wpokt-oracle/ethereum/util"
)

type WarpISMContract interface {
	ValidatorCount(opts *bind.CallOpts) (*big.Int, error)
	Eip712Domain(opts *bind.CallOpts) (util.DomainData, error)
}

type WarpISMContractImpl struct {
	contract *autogen.WarpISM
	address  common.Address
}

func (x *WarpISMContractImpl) Address() common.Address {
	return x.address
}

func (x *WarpISMContractImpl) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	return x.contract.ValidatorCount(opts)
}

func (x *WarpISMContractImpl) Eip712Domain(opts *bind.CallOpts) (util.DomainData, error) {
	return x.contract.Eip712Domain(opts)
}

func NewWarpISMContract(address common.Address, client *ethclient.Client) (WarpISMContract, error) {
	contract, err := autogen.NewWarpISM(address, client)
	if err != nil {
		return nil, err
	}

	return &WarpISMContractImpl{contract: contract, address: address}, nil
}