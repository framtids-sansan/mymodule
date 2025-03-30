package keeper

import (
	"context"
	"github.com/framtids-sansan/mymodule/extramint/types"
)

package keeper

import (
"context"
"fmt"

"cosmossdk.io/collections"
"cosmossdk.io/log"

"github.com/cosmos/cosmos-sdk/types"
"github.com/cosmos/cosmos-sdk/types/query"
"github.com/framtids-sansan/mymodule/types"
)

// InitGenesis initializes the module's state from a given genesis state.
func (k BaseKeeper) InitGenesis(ctx context.Context, genState *types.GenesisState) {
	// Set any parameters for the module from the genesis state
	if err := k.SetParams(ctx, genState.Params); err != nil {
		k.logger.Error("failed to set params", "error", err)
		panic(err)
	}

	// Initialize other module state (if applicable)
	// For this example, no other state is needed to be initialized
	// You can add any other logic related to initializing your module's state here

	k.logger.Debug("Module genesis initialized", "genesis_state", genState)
}

// ExportGenesis returns the module's genesis state.
func (k BaseKeeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	// Export any parameters related to the module
	params := k.GetParams(ctx)

	// Export other state (if applicable)
	// In this case, no specific state needs to be exported, so we'll use empty values
	genesisState := types.NewGenesisState(
		params,
	// For this module, we may not have any specific state to export
	// If any state like balances or configurations need exporting, handle

