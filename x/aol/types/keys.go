package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "aol"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	OwnerKey  = "Owner-value-"
	TopicKey  = "Topic-value-"
	WriterKey = "Writer-value-"
	RecordKey = "Record-value-"

	compositeKeySeparator = "/"
)

type CompositeKey interface {
	Marshal() string
	Unmarshal(key string) error
}

func MarshalBytes(k CompositeKey) []byte {
	return []byte(k.Marshal())
}

func MustUnmarshalCompositeKey(k CompositeKey, key string) {
	if err := k.Unmarshal(key); err != nil {
		panic(fmt.Sprintf("failed to unmarshal composite key: %v", err))
	}
}

var (
	_ CompositeKey = &OwnerCompositeKey{}
	_ CompositeKey = &TopicCompositeKey{}
	_ CompositeKey = &WriterCompositeKey{}
	_ CompositeKey = &RecordCompositeKey{}
)

type OwnerCompositeKey struct {
	OwnerAddress string
}

func (k OwnerCompositeKey) Marshal() string {
	return k.OwnerAddress
}

func (k *OwnerCompositeKey) Unmarshal(key string) error {
	k.OwnerAddress = key
	return nil
}

type TopicCompositeKey struct {
	OwnerAddress string
	TopicName    string
}

func (k TopicCompositeKey) Marshal() string {
	return fmt.Sprintf("%s%s%s", k.OwnerAddress, compositeKeySeparator, k.TopicName)
}

func (k *TopicCompositeKey) Unmarshal(key string) error {
	sp := strings.Split(key, compositeKeySeparator)
	if len(sp) != 2 {
		return errors.New("invalid format for TopicCompositeKey")
	}

	if _, err := sdk.AccAddressFromBech32(sp[0]); err != nil {
		return fmt.Errorf("failed to decode owner address: %w", err)
	}

	k.OwnerAddress = sp[0]
	k.TopicName = sp[1]
	return nil
}

type WriterCompositeKey struct {
	OwnerAddress  string
	TopicName     string
	WriterAddress string
}

func (k WriterCompositeKey) Marshal() string {
	return fmt.Sprintf(
		"%s%s%s%s%s",
		k.OwnerAddress,
		compositeKeySeparator,
		k.TopicName,
		compositeKeySeparator,
		k.WriterAddress,
	)
}

func (k *WriterCompositeKey) Unmarshal(key string) error {
	sp := strings.Split(key, compositeKeySeparator)
	if len(sp) != 3 {
		return errors.New("invalid format for WriterCompositeKey")
	}

	if _, err := sdk.AccAddressFromBech32(sp[0]); err != nil {
		return fmt.Errorf("failed to decode owner address: %w", err)
	}
	if _, err := sdk.AccAddressFromBech32(sp[2]); err != nil {
		return fmt.Errorf("failed to decode writer address: %w", err)
	}

	k.OwnerAddress = sp[0]
	k.TopicName = sp[1]
	k.WriterAddress = sp[2]
	return nil
}

type RecordCompositeKey struct {
	OwnerAddress string
	TopicName    string
	Offset       uint64
}

func (k RecordCompositeKey) Marshal() string {
	return fmt.Sprintf(
		"%s%s%s%s%d",
		k.OwnerAddress,
		compositeKeySeparator,
		k.TopicName,
		compositeKeySeparator,
		k.Offset,
	)
}

func (k *RecordCompositeKey) Unmarshal(key string) error {
	sp := strings.Split(key, compositeKeySeparator)
	if len(sp) != 3 {
		return errors.New("invalid format for RecordCompositeKey")
	}

	if _, err := sdk.AccAddressFromBech32(sp[0]); err != nil {
		return fmt.Errorf("failed to decode owner address: %w", err)
	}
	offset, err := strconv.ParseUint(sp[2], 10, 64)
	if err != nil {
		return err
	}

	k.OwnerAddress = sp[0]
	k.TopicName = sp[1]
	k.Offset = offset
	return nil
}
