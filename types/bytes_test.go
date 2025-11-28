package types

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestUnitBytesMarshal(t *testing.T) {
	testCases := []struct {
		in     UnitBytes
		expect string
	}{
		{
			in:     UnitBytes(128),
			expect: "\"128b\"",
		},
		{
			in:     UnitBytes(1024),
			expect: "\"1kib\"",
		},
		{
			in:     UnitBytes(1024 * 1024),
			expect: "\"1mib\"",
		},
		{
			in:     UnitBytes(1024 * 1024 * 1024),
			expect: "\"1gib\"",
		},
		{
			in:     UnitBytes(1024 * 1024 * 1024 * 1024),
			expect: "\"1tib\"",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.expect, func(t *testing.T) {
			v, err := tC.in.MarshalJSON()
			assert.NilError(t, err)
			assert.Equal(t, tC.expect, string(v))
		})
	}
}

func TestDecodeUnitBytes(t *testing.T) {
	testCases := []struct {
		in     string
		expect UnitBytes
	}{
		{
			in:     "128b",
			expect: 128,
		},
		{
			in:     "128kib",
			expect: 128 * (1 << 10),
		},
		{
			in:     "128mib",
			expect: 128 * (1 << 20),
		},
		{
			in:     "128gib",
			expect: 128 * (1 << 30),
		},
		{
			in:     "128tib",
			expect: 128 * (1 << 40),
		},
	}
	for _, tC := range testCases {
		t.Run(string(tC.in), func(t *testing.T) {
			var got UnitBytes
			err := got.DecodeMapstructure(tC.in)
			assert.NilError(t, err)
			assert.Equal(t, tC.expect, got)
		})
	}
}
