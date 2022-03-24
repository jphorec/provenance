package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	epochtypes "github.com/provenance-io/provenance/x/epoch/types"
	"github.com/provenance-io/provenance/x/reward/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	// distribute logic goes here, i.e record the number of shares claimable in that epoch and the total rewards pool
	// also unlock the module account?
	ctx.Logger().Info(fmt.Sprintf("In epoch end for %s %d", epochIdentifier, epochNumber))
	var rewardPrograms []types.RewardProgram
	// get all the rewards programs
	err := k.IterateRewardPrograms(ctx, func(rewardProgram types.RewardProgram) (stop bool) {
		rewardPrograms = append(rewardPrograms, rewardProgram)
		return false
	})
	if err != nil {
		return
	}

	for _, rewardProgram := range rewardPrograms {
		if rewardProgram.EpochId == epochIdentifier {
			// this is epoch that ended, and matches up with the reward program identifier
			// check if any of the events match with any of the reward program running

		}
	}

}

// ___________________________________________________________________________________________________

// Hooks wrapper struct for incentives keeper
type Hooks struct {
	k Keeper
}

var _ epochtypes.EpochHooks = Hooks{}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}