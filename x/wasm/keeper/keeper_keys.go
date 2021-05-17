package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	KeyDelimiter = []byte{0x00}

	CodeKeyPrefix = []byte{0x11} // {Prefix}{CodeID}
)

func CodeKey(codeID uint64) []byte {
	return bytes.Join([][]byte{
		CodeKeyPrefix,
		sdk.Uint64ToBigEndian(codeID),
	}, KeyDelimiter)
}
