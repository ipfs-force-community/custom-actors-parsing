package market

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs-force-community/custom-actors-parsing/external"
)

var Methods = map[uint64]external.MethodMeta{
	1:  {"Constructor", *new(func(value *abi.EmptyValue) *abi.EmptyValue)},
	2:  {"AddBalance", *new(func(*address.Address) *abi.EmptyValue)},
	3:  {"WithdrawBalance", *new(func(*WithdrawBalanceParams) *WithdrawBalanceReturn)},
	4:  {"PublishStorageDeals", *new(func(*PublishStorageDealsParams) *PublishStorageDealsReturn)},
	5:  {"VerifyDealsForActivation", *new(func(*VerifyDealsForActivationParams) *VerifyDealsForActivationReturn)},
	6:  {"ActivateDeals", *new(func(*ActivateDealsParams) *abi.EmptyValue)},
	7:  {"OnMinerSectorsTerminate", *new(func(*OnMinerSectorsTerminateParams) *abi.EmptyValue)},
	8:  {"ComputeDataCommitment", *new(func(*ComputeDataCommitmentParams) *ComputeDataCommitmentReturn)},
	9:  {"CronTick", *new(func(*abi.EmptyValue) *abi.EmptyValue)},
	42: {"CheckHealth", *new(func(*DealHealthyParams) *abi.EmptyValue)},
}
