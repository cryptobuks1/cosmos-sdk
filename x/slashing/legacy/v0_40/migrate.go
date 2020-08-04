package v040

// DONTCOVER
// nolint

import (
	"sort"

	v039slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v0_39"
)

// Migrate accepts exported x/slashing genesis state from v0.39 and migrates it
// to v0.40 x/slashing genesis state. The migration includes:
//
// - Chaning SigningInfos and MissedBlocks from map to array.
func Migrate(oldGenState v039slashing.GenesisState) GenesisState {
	// Note that the two following `for` loop over a map's keys, so are not
	// deterministic.
	var si = make([]SigningInfo, 0, len(oldGenState.SigningInfos))
	for address, signingInfo := range oldGenState.SigningInfos {
		si = append(si, SigningInfo{
			Address:              address,
			ValidatorSigningInfo: signingInfo,
		})
	}
	var mb = make([]ValidatorMissedBlocks, 0, len(oldGenState.MissedBlocks))
	for address, missedBlocks := range oldGenState.MissedBlocks {
		mb = append(mb, ValidatorMissedBlocks{
			Address:      address,
			MissedBlocks: missedBlocks,
		})
	}

	// We sort these two arrays by address, so that we get determinstic states.
	sort.Slice(si, func(i, j int) bool { return si[i].Address < si[j].Address })
	sort.Slice(mb, func(i, j int) bool { return mb[i].Address < mb[j].Address })

	return GenesisState{
		Params:       oldGenState.Params,
		SigningInfos: si,
		MissedBlocks: mb,
	}
}