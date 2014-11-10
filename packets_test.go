package dst

import (
	"encoding/hex"
	"reflect"
	"strings"
	"testing"
)

var headerTests = []struct {
	hex string
	hdr header
}{
	{
		"00000000 00000000 00000000",
		header{},
	},
	{
		"00123456 A0008901 98765432",
		header{
			packetType: typeData,
			flags:      0,
			connID:     0x123456,
			sequenceNo: 0xA0008901,
			timestamp:  0x98765432,
		},
	},
	{
		"34340000 22334455 55667788",
		header{
			packetType: typeACK,
			flags:      flagsCookie,
			connID:     0x340000,
			sequenceNo: 0x22334455,
			timestamp:  0x55667788,
		},
	},
}

func TestEncodeHeaders(t *testing.T) {
	for i, tc := range headerTests {
		var actual [dstHeaderLen]byte
		tc.hdr.marshal(actual[:])
		expected, _ := hex.DecodeString(strings.Replace(tc.hex, " ", "", -1))

		if !reflect.DeepEqual(actual[:], expected) {
			t.Errorf("Encode %d incorrect;\n  A: %#v\n  E: %#v", i, actual[:], expected)
		}
	}
}

func TestDecodeHeaders(t *testing.T) {
	for i, tc := range headerTests {
		data, _ := hex.DecodeString(strings.Replace(tc.hex, " ", "", -1))
		var actual header
		actual.unmarshal(data)

		if !reflect.DeepEqual(actual, tc.hdr) {
			t.Errorf("Decode %d incorrect;\n  A: %#v\n  E: %#v", i, actual, tc.hdr)
		}
	}
}
