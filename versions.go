package versioner

import (
	"fmt"
	"strconv"
)

// Managed versioning strategy
type Versioned[V VersioningStrategy] struct {
	Version V
}

func (v Versioned[V]) String() string {
	return v.Version.Version()
}

func compareStrings(a, b string) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

type anyint interface {
	~int | ~uint64
}

func compareInts[I anyint](a, b I) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

// The required interface of a versioning strategy,
// allowing string representations of versions, as well
// as comparison between types
type VersioningStrategy interface {
	Version() string
	compare(other VersioningStrategy) int
}

type IntVersion uint64

var _ VersioningStrategy = (*IntVersion)(nil)

// compare implements VersioningStrategy.
func (v IntVersion) compare(other VersioningStrategy) int {
	panic("Not implemented")
}

func (v IntVersion) Version() string {
	return strconv.FormatUint(uint64(v), 10)
}

type SemanticVersionOpts struct {
	// The number of version components to use in the semantic
	// version. By default, three parts are used
	NumParts uint64

	// A cosmetic prefix (ignored in version calculations)
	Prefix string

	// If a suffix is present, it will be treated as after the
	// version without the suffix, and before the next version
	// without the suffix. If multiple suffixes are present for
	// the same version, they are sorted alphabetically
	Suffix string
}

type SemanticVersion interface {
	VersioningStrategy
	Opts() SemanticVersionOpts
	Parts() []IntVersion
}

type semanticVersion struct {
	opts  SemanticVersionOpts
	parts []IntVersion
}

func NewSemanticVersion(opts SemanticVersionOpts, parts ...IntVersion) SemanticVersion {
	if opts.NumParts == 0 {
		panic("NumParts must be > 0")
	}

	if opts.NumParts != uint64(len(parts)) {
		panic(fmt.Sprintf("Expected %d parts, but got %d (%v)",
			opts.NumParts, len(parts), parts))
	}

	return semanticVersion{
		opts:  opts,
		parts: parts,
	}
}

func (sv semanticVersion) Opts() SemanticVersionOpts {
	return sv.opts
}

func (sv semanticVersion) Parts() []IntVersion {
	parts := make([]IntVersion, sv.opts.NumParts)

	copy(parts, sv.parts)

	return parts
}

// Version implements SemanticVersion.
func (sv semanticVersion) Version() string {
	panic("unimplemented")
}

// compare implements SemanticVersion.
func (sv semanticVersion) compare(other VersioningStrategy) int {
	panic("unimplemented")
}
