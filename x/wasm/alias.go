package wasm

import (
	"github.com/medibloc/panacea-core/x/wasm/keeper"
	"github.com/medibloc/panacea-core/x/wasm/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute
)

var (
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc
	NewKeeper     = keeper.NewKeeper
)

type (
	Keeper = keeper.Keeper

	//TODO: msgs
)
