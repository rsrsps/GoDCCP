// Copyright 2010 GoDCCP Authors. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package dccp

import (
	"rand"
	"testing"
)

func TestWireEncodeDecode(t *testing.T) {

	buf := make([]byte, 8)
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	
	// 1 byte
	u8 = uint8(rand.Int31())
	encode1ByteUint(u8, buf[0:1])
	if decode1ByteUint(buf[0:1]) != u8 {
		t.Errorf("E/D 1 byte")
	}
	
	// 2 byte
	u16 = uint16(rand.Int31())
	encode2ByteUint(u16, buf[0:2])
	if decode2ByteUint(buf[0:2]) != u16 {
		t.Errorf("E/D 2 byte")
	}
	
	// 3 byte
	u32 = uint32(rand.Int31())
	u32 = (u32 << 8) >> 8
	encode3ByteUint(u32, buf[0:3])
	if decode3ByteUint(buf[0:3]) != u32 {
		t.Errorf("E/D 3 byte")
	}
	
	// 4 byte
	u32 = uint32(rand.Int31()) << 1
	encode4ByteUint(u32, buf[0:4])
	if decode4ByteUint(buf[0:4]) != u32 {
		t.Errorf("E/D 4 byte")
	}

	// 6 byte
	u64 = uint64(rand.Int63())
	u64 = (u64 << (2*8)) >> (2*8)
	encode6ByteUint(u64, buf[0:6])
	if decode6ByteUint(buf[0:6]) != u64 {
		t.Errorf("E/D 6 byte")
	}
}
