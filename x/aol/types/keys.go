package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
)

type CompositeKey interface {
	Marshal() string
	Unmarshal(key string) error
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
	return fmt.Sprintf("%s/%s", k.OwnerAddress, k.TopicName)
}

func (k *TopicCompositeKey) Unmarshal(key string) error {
	sp := strings.Split(key, "/")
	if len(sp) != 2 {
		return errors.New("invalid format for genesis topic key")
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
	return fmt.Sprintf("%s/%s/%s", k.OwnerAddress, k.TopicName, k.WriterAddress)
}

func (k *WriterCompositeKey) Unmarshal(key string) error {
	sp := strings.Split(key, "/")
	if len(sp) != 3 {
		return errors.New("invalid format for genesis writer key")
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
	return fmt.Sprintf("%s/%s/%d", k.OwnerAddress, k.TopicName, k.Offset)
}

func (k *RecordCompositeKey) Unmarshal(key string) error {
	sp := strings.Split(key, "/")
	if len(sp) != 3 {
		return errors.New("invalid format for genesis record key")
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
