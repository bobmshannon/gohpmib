# gohpmib [![CircleCI](https://circleci.com/gh/bobmshannon/gohpmib.svg?style=svg)](https://circleci.com/gh/bobmshannon/gohpmib) [![GoDoc](https://godoc.org/github.com/bobmshannon/gohpmib?status.svg)](https://godoc.org/github.com/bobmshannon/gohpmib) [![Go Report Card](https://goreportcard.com/badge/github.com/bobmshannon/gohpmib)](https://goreportcard.com/report/github.com/bobmshannon/gohpmib)
A Go library for programmatically interacting with an HP Management Information Base (MIB) using SNMP. For documentation on how to use this library, see [godoc](https://godoc.org/github.com/bobmshannon/gohpmib).

## Compatibility

HP servers that have the HP health SNMP agent installed and running on the host operating system are currently supported.

| Generation    | Supported     |
| ------------- |:-------------:|
| `7`           | ✅            |
| `8`           | ✅            |
| `9`           | ✅            |
| `10 `         | Experimental  |

## Bug reports

If a bug is discovered, file a GitHub issue with the following information:

- Description of the problem
- Steps to reproduce the problem
- Version of HP health agents installed (if applicable)
- Linux distribution and kernel version
- Any other relevant information that will be useful for debugging and reproducing the problem

## Features and enhancements

Feature and enhancement requests including support for additional fields and OIDs should first be raised as a GitHub issue.

## Bugfixes

Minor and trivial bugfixes can be addressed via a PR without raising a GitHub issue.
