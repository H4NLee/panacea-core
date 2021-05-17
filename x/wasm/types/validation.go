package types

import (
	"fmt"
	"net/url"
	"regexp"
)

const (
	MaxWasmSize = 500 * 1024

	// MaxLabelSize is the longest label that can be used when Instantiating a contract
	MaxLabelSize = 128

	// BuildTagRegexp is a docker image regexp.
	// We only support max 128 characters, with at least one organization name (subset of all legal names).
	//
	// Details from https://docs.docker.com/engine/reference/commandline/tag/#extended-description :
	//
	// An image name is made up of slash-separated name components (optionally prefixed by a registry hostname).
	// Name components may contain lowercase characters, digits and separators.
	// A separator is defined as a period, one or two underscores, or one or more dashes. A name component may not start or end with a separator.
	//
	// A tag name must be valid ASCII and may contain lowercase and uppercase letters, digits, underscores, periods and dashes.
	// A tag name may not start with a period or a dash and may contain a maximum of 128 characters.
	BuildTagRegexp = "^[a-z0-9][a-z0-9._-]*[a-z0-9](/[a-z0-9][a-z0-9._-]*[a-z0-9])+:[a-zA-Z0-9_][a-zA-Z0-9_.-]*$"

	MaxBuildTagSize = 128
)

func validateSourceURL(source string) error {
	if source != "" {
		u, err := url.Parse(source)
		if err != nil {
			return ErrInvalid(fmt.Errorf("not an url: %s", source))
		}
		if !u.IsAbs() {
			return ErrInvalid(fmt.Errorf("not an absolute url: %s", source))
		}
		if u.Scheme != "https" {
			return ErrInvalid(fmt.Errorf("must use https: %s", source))
		}
	}
	return nil
}

func validateBuilder(buildTag string) error {
	if len(buildTag) > MaxBuildTagSize {
		return ErrLimit(fmt.Errorf("build tag is longer than %d characters: %s", MaxBuildTagSize, buildTag))
	}

	if buildTag != "" {
		ok, err := regexp.MatchString(BuildTagRegexp, buildTag)
		if err != nil || !ok {
			return ErrInvalid(fmt.Errorf("invalid build tag: %s", buildTag))
		}
	}
	return nil
}

func validateWasmCode(s []byte) error {
	if len(s) == 0 {
		return ErrEmpty("wasm code")
	}
	if len(s) > MaxWasmSize {
		return ErrLimit(fmt.Errorf("wasm code is longer than %d bytes: %d", MaxWasmSize, len(s)))
	}
	return nil
}

func validateLabel(label string) error {
	if label == "" {
		return ErrEmpty("label")
	}
	if len(label) > MaxLabelSize {
		return ErrLimit(fmt.Errorf("label is longer than %d characters: %d", MaxLabelSize, len(label)))
	}
	return nil
}
