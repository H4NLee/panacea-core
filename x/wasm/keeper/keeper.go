package keeper

import (
	"encoding/binary"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/medibloc/panacea-core/x/wasm/types"

	wasm "github.com/CosmWasm/go-cosmwasm"
)

// GasMultiplier is how many cosmwasm gas points = 1 SDK gas point.
// SDK reference costs can be found here: https://github.com/cosmos/cosmos-sdk/blob/02c6c9fafd58da88550ab4d7d494724a477c8a68/store/types/gas.go#L153-L164
const GasMultiplier = 100

// MaxGas for a contract is 900 million (enforced in the CosmWasm Rust implementation)
const MaxGas = 900_000_000

// InstanceCost is how much SDK gas we charge each time we load a WASM instance.
// Creating a new instance is costly, and this helps put a recursion limit to contracts calling contracts.
const InstanceCost uint64 = 40_000

// CompileCost is how much SDK gas we charge *per byte* for compiling WASM code.
const CompileCost uint64 = 2

type Keeper interface {
	Codec() *codec.Codec
	//TODO: other funcs
}

// wasmKeeper will have a reference to Wasmer with its own data directory.
type wasmKeeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	cdc           *codec.Codec
	accountKeeper auth.AccountKeeper
	bankKeeper    bank.Keeper

	wasmer       wasm.Wasmer
	queryPlugins QueryPlugins
	messenger    MessageHandler

	// queryGasLimit is the max wasm gas that can be spent on executing a query with a contract
	queryGasLimit uint64
}

func NewKeeper(
	storeKey sdk.StoreKey, cdc *codec.Codec, accountKeeper auth.AccountKeeper, bankKeeper bank.Keeper,
	stakingKeeper staking.Keeper, router sdk.Router, homeDir string, wasmConfig types.WasmConfig,
	supportedFeatures string, customEncoders *MessageEncoders, customPlugins *QueryPlugins,
) Keeper {
	wasmer, err := wasm.NewWasmer(filepath.Join(homeDir, "wasm"), supportedFeatures, wasmConfig.CacheSize)
	if err != nil {
		panic(err)
	}

	keeper := wasmKeeper{
		storeKey:      storeKey,
		cdc:           cdc,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		wasmer:        *wasmer,
		messenger:     NewMessageHandler(router, customEncoders),
		queryGasLimit: wasmConfig.SmartQueryGasLimit,
	}
	keeper.queryPlugins = DefaultQueryPlugins(bankKeeper, stakingKeeper, keeper).Merge(customPlugins)
	return keeper
}

func (k wasmKeeper) Codec() *codec.Codec {
	return k.cdc
}

// Create uploads and compiles a WASM contract, returning a short identifier for the contract.
func (k wasmKeeper) Create(ctx sdk.Context, creator sdk.AccAddress, wasmCode []byte, source string, builder string) (uint64, error) {
	wasmCode, err := uncompress(wasmCode)
	if err != nil {
		return 0, types.ErrCreateFailed(err)
	}
	ctx.GasMeter().ConsumeGas(CompileCost*uint64(len(wasmCode)), "Compiling WASM Bytecode")

	codeHash, err := k.wasmer.Create(wasmCode)
	if err != nil {
		return 0, types.ErrCreateFailed(err)
	}

	store := ctx.KVStore(k.storeKey)
	codeID := k.autoIncrementID(ctx, types.KeyLastCodeID)
	codeInfo := types.NewCodeInfo(codeHash, creator, source, builder)
	store.Set(CodeKey(codeID), k.cdc.MustMarshalBinaryBare(codeInfo))

	return codeID, nil
}

func (k wasmKeeper) Instantiate

func (k wasmKeeper) autoIncrementID(ctx sdk.Context, lastIDKey []byte) uint64 {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(lastIDKey)
	id := uint64(1)
	if bz != nil {
		id = binary.BigEndian.Uint64(bz)
	}

	store.Set(lastIDKey, sdk.Uint64ToBigEndian(id+1))
	return id
}
