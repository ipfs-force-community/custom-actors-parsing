package reward

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/reward"
	"github.com/filecoin-project/go-state-types/builtin/v8/util/smoothing"
)

type State struct {
	CumsumBaseline          reward.Spacetime
	CumsumRealized          reward.Spacetime
	EffectiveNetworkTime    abi.ChainEpoch
	EffectiveBaselinePower  abi.StoragePower
	ThisEpochReward         abi.TokenAmount
	ThisEpochRewardSmoothed smoothing.FilterEstimate
	ThisEpochBaselinePower  abi.StoragePower
	Epoch                   abi.ChainEpoch
	TotalStoragePowerReward abi.TokenAmount
	SimpleTotal             abi.TokenAmount
	BaselineTotal           abi.TokenAmount
}
