package hpmib

import (
	"errors"
)

var (
	// ErrNoResultsReturned is the error returned when no results are returned when querying the MIB for a specific OID.
	ErrNoResultsReturned = errors.New("querying MIB returned no results")
	// ErrExpectedInteger is the error returned when the value at a specific OID was expected to be an integer but another
	// type was returned instead.
	ErrExpectedInteger = errors.New("expected value to be an integer")
	// ErrExpectedOctetString is the error returned when the value at a specific OID was expected to be an octet string but
	// another type was returned instead.
	ErrExpectedOctetString = errors.New("expected value to be a octet string")
)
