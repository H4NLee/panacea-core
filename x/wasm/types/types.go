package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/common"
)

const defaultLRUCacheSize = uint64(0)
const defaultQueryGasLimit = uint64(3000000)

// Model is a struct that holds a KV pair
type Model struct {
	// hex-encode key to read it better (this is often ascii)
	Key common.HexBytes `json:"key"`
	// base64-encode raw value
	Value []byte `json:"val"`
}

func (m Model) ValidateBasic() error {
	if len(m.Key) == 0 {
		return ErrEmpty("key")
	}
	return nil
}

// CodeInfo is data for the uploaded contract WASM code
type CodeInfo struct {
	CodeHash []byte         `json:"code_hash"`
	Creator  sdk.AccAddress `json:"creator"`
	Source   string         `json:"source"`
	Builder  string         `json:"builder"`
}

// NewCodeInfo fills a new CodeInfo struct
func NewCodeInfo(codeHash []byte, creator sdk.AccAddress, source string, builder string) CodeInfo {
	return CodeInfo{
		CodeHash: codeHash,
		Creator:  creator,
		Source:   source,
		Builder:  builder,
	}
}

func (c CodeInfo) ValidateBasic() error {
	if len(c.CodeHash) == 0 {
		return ErrEmpty("code hash")
	}
	if err := sdk.VerifyAddressFormat(c.Creator); err != nil {
		return err
	}
	if err := validateSourceURL(c.Source); err != nil {
		return err
	}
	if err := validateBuilder(c.Builder); err != nil {
		return err
	}
	return nil
}

// WasmConfig is the extra config required for wasm
type WasmConfig struct {
	SmartQueryGasLimit uint64 `mapstructure:"query_gas_limit"`
	CacheSize          uint64 `mapstructure:"lru_size"`
}

// DefaultWasmConfig returns the default settings for WasmConfig
func DefaultWasmConfig() WasmConfig {
	return WasmConfig{
		SmartQueryGasLimit: defaultQueryGasLimit,
		CacheSize:          defaultLRUCacheSize,
	}
}
