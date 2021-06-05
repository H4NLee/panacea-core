package types

import (
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MaxTopicLength       = 70
	MaxMonikerLength     = 70
	MaxDescriptionLength = 5000
	MaxRecordKeyLength   = 70
	MaxRecordValueLength = 5000
)

func validateTopic(topic string) error {
	if len(topic) > MaxTopicLength {
		return sdkerrors.Wrapf(ErrMessageTooLarge, "topic (%d > %d)", len(topic), MaxTopicLength)
	}

	// cannot be an empty string
	if !regexp.MustCompile("^[A-Za-z0-9._-]+$").MatchString(topic) {
		return sdkerrors.Wrapf(ErrInvalidTopic, "topic %s", topic)
	}

	return nil
}

func validateMoniker(moniker string) error {
	if len(moniker) > MaxMonikerLength {
		return sdkerrors.Wrapf(ErrMessageTooLarge, "moniker (%d > %d)", len(moniker), MaxMonikerLength)
	}

	// can be an empty string
	if !regexp.MustCompile("^[A-Za-z0-9._-]*$").MatchString(moniker) {
		return sdkerrors.Wrapf(ErrInvalidMoniker, "moniker %s", moniker)
	}

	return nil
}

func validateDescription(description string) error {
	if len(description) > MaxDescriptionLength {
		return sdkerrors.Wrapf(ErrMessageTooLarge, "description (%d > %d)", len(description), MaxDescriptionLength)
	}
	return nil
}

func validateRecordKey(key []byte) error {
	if len(key) > MaxRecordKeyLength {
		return sdkerrors.Wrapf(ErrMessageTooLarge, "key (%d > %d)", len(key), MaxRecordKeyLength)
	}
	return nil
}

func validateRecordValue(value []byte) error {
	if len(value) > MaxRecordValueLength {
		return sdkerrors.Wrapf(ErrMessageTooLarge, "value (%d > %d)", len(value), MaxRecordValueLength)
	}
	return nil
}
