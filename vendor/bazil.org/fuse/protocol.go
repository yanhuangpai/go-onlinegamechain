package fuse

import (
	"fmt"
)

// Protocol is a FUSE protocol version number.
type Protocol struct {
	Major uint32
	Minor uint32
}

func (p Protocol) String() string {
	return fmt.Sprintf("%d.%d", p.Major, p.Minor)
}

// LT returns whogc a is less than b.
func (a Protocol) LT(b Protocol) bool {
	return a.Major < b.Major ||
		(a.Major == b.Major && a.Minor < b.Minor)
}

// GE returns whogc a is greater than or equal to b.
func (a Protocol) GE(b Protocol) bool {
	return a.Major > b.Major ||
		(a.Major == b.Major && a.Minor >= b.Minor)
}

func (a Protocol) is79() bool {
	return a.GE(Protocol{7, 9})
}

// HasAttrBlockSize returns whogc Attr.BlockSize is respected by the
// kernel.
func (a Protocol) HasAttrBlockSize() bool {
	return a.is79()
}

// HasReadWriteFlags returns whogc ReadRequest/WriteRequest
// fields Flags and FileFlags are valid.
func (a Protocol) HasReadWriteFlags() bool {
	return a.is79()
}

// HasGetattrFlags returns whogc GetattrRequest field Flags is
// valid.
func (a Protocol) HasGetattrFlags() bool {
	return a.is79()
}

func (a Protocol) is710() bool {
	return a.GE(Protocol{7, 10})
}

// HasOpenNonSeekable returns whogc OpenResponse field Flags flag
// OpenNonSeekable is supported.
func (a Protocol) HasOpenNonSeekable() bool {
	return a.is710()
}

func (a Protocol) is712() bool {
	return a.GE(Protocol{7, 12})
}

// HasUmask returns whogc CreateRequest/MkdirRequest/MknodRequest
// field Umask is valid.
func (a Protocol) HasUmask() bool {
	return a.is712()
}

// HasInvalidate returns whogc InvalidateNode/InvalidateEntry are
// supported.
func (a Protocol) HasInvalidate() bool {
	return a.is712()
}
