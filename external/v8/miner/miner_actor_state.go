package miner

import (
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

type State struct {
	// Information not related to sectors.
	Info                       cid.Cid
	PreCommitDeposits          abi.TokenAmount // Total funds locked as PreCommitDeposits
	LockedFunds                abi.TokenAmount // Total rewards and added funds locked in vesting table
	VestingFunds               cid.Cid         // VestingFunds (Vesting Funds schedule for the miner).
	FeeDebt                    abi.TokenAmount // Absolute value of debt this miner owes from unpaid fees
	InitialPledge              abi.TokenAmount // Sum of initial pledge requirements of all active sectors
	PreCommittedSectors        cid.Cid         // Map, HAMT[SectorNumber]SectorPreCommitOnChainInfo
	PreCommittedSectorsCleanUp cid.Cid         // BitFieldQueue (AMT[Epoch]*BitField)
	AllocatedSectors           cid.Cid         // BitField
	Sectors                    cid.Cid         // Array, AMT[SectorNumber]SectorOnChainInfo (sparse)
	ProvingPeriodStart         abi.ChainEpoch
	CurrentDeadline            uint64
	Deadlines                  cid.Cid
	EarlyTerminations          bitfield.BitField
	DeadlineCronActive         bool
}
