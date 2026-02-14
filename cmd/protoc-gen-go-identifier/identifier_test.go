package main_test

import (
	"testing"

	td "github.com/OldBigBuddha/protoc-gen-go-identifier/cmd/protoc-gen-go-identifier/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- StringID ---

func TestStringID_Constructor(t *testing.T) {
	id := td.AsStringID("test-123")
	require.NotNil(t, id)
	assert.Equal(t, "test-123", id.Unwrap())
}

func TestStringID_Unwrap_Nil(t *testing.T) {
	var id *td.StringID
	assert.Equal(t, "", id.Unwrap())
}

func TestStringID_Equal(t *testing.T) {
	tests := []struct {
		name string
		a    *td.StringID
		b    *td.StringID
		want bool
	}{
		{"same value", td.AsStringID("abc"), td.AsStringID("abc"), true},
		{"different value", td.AsStringID("abc"), td.AsStringID("xyz"), false},
		{"both nil", nil, nil, true},
		{"one nil", td.AsStringID("abc"), nil, false},
		{"empty string equal", td.AsStringID(""), td.AsStringID(""), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.Equal(tt.b))
		})
	}
}

func TestStringID_Clone(t *testing.T) {
	original := td.AsStringID("clone-me")
	cloned := original.Clone()

	assert.NotSame(t, original, cloned)
	assert.True(t, cloned.Equal(original))

	cloned.Id = "mutated"
	assert.NotEqual(t, "mutated", original.Unwrap())
}

// --- IntID ---

func TestIntID_Constructor(t *testing.T) {
	id := td.AsIntID(42)
	require.NotNil(t, id)
	assert.Equal(t, int64(42), id.Unwrap())
}

func TestIntID_Unwrap_Nil(t *testing.T) {
	var id *td.IntID
	assert.Equal(t, int64(0), id.Unwrap())
}

func TestIntID_Equal(t *testing.T) {
	tests := []struct {
		name string
		a    *td.IntID
		b    *td.IntID
		want bool
	}{
		{"same value", td.AsIntID(1), td.AsIntID(1), true},
		{"different value", td.AsIntID(1), td.AsIntID(2), false},
		{"both nil", nil, nil, true},
		{"one nil", td.AsIntID(1), nil, false},
		{"zero equal", td.AsIntID(0), td.AsIntID(0), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.Equal(tt.b))
		})
	}
}

func TestIntID_Clone(t *testing.T) {
	original := td.AsIntID(99)
	cloned := original.Clone()

	assert.NotSame(t, original, cloned)
	assert.True(t, cloned.Equal(original))
}

// --- BytesID ---

func TestBytesID_Constructor(t *testing.T) {
	id := td.AsBytesID([]byte{0xDE, 0xAD})
	require.NotNil(t, id)
	assert.Equal(t, []byte{0xDE, 0xAD}, id.Unwrap())
}

func TestBytesID_Unwrap_Nil(t *testing.T) {
	var id *td.BytesID
	assert.Nil(t, id.Unwrap())
}

func TestBytesID_Equal(t *testing.T) {
	tests := []struct {
		name string
		a    *td.BytesID
		b    *td.BytesID
		want bool
	}{
		{"same value", td.AsBytesID([]byte{1}), td.AsBytesID([]byte{1}), true},
		{"different value", td.AsBytesID([]byte{1}), td.AsBytesID([]byte{2}), false},
		{"both nil", nil, nil, true},
		{"one nil", td.AsBytesID([]byte{1}), nil, false},
		{"both empty", td.AsBytesID([]byte{}), td.AsBytesID([]byte{}), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.Equal(tt.b))
		})
	}
}

func TestBytesID_Clone(t *testing.T) {
	original := td.AsBytesID([]byte{0xCA, 0xFE})
	cloned := original.Clone()

	assert.NotSame(t, original, cloned)
	assert.True(t, cloned.Equal(original))
}

// --- PartialID (skip_constructor, skip_clone) ---

func TestPartialID_Unwrap(t *testing.T) {
	id := &td.PartialID{Id: "partial"}
	assert.Equal(t, "partial", id.Unwrap())
}

func TestPartialID_Unwrap_Nil(t *testing.T) {
	var id *td.PartialID
	assert.Equal(t, "", id.Unwrap())
}

func TestPartialID_Equal(t *testing.T) {
	a := &td.PartialID{Id: "x"}
	b := &td.PartialID{Id: "x"}
	assert.True(t, a.Equal(b))

	c := &td.PartialID{Id: "y"}
	assert.False(t, a.Equal(c))
}

