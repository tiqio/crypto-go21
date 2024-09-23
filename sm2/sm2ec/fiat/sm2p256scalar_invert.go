// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by addchain. DO NOT EDIT.
package fiat

// Invert sets e = 1/x, and returns e.
//
// If x == 0, Invert returns e = 0.
func (e *SM2P256OrderElement) Invert(x *SM2P256OrderElement) *SM2P256OrderElement {
	// Inversion is implemented as exponentiation with exponent p − 2.
	// The sequence of 41 multiplications and 253 squarings is derived from the
	// following addition chain generated with github.com/mmcloughlin/addchain v0.4.0.
	//
	//	_10      = 2*1
	//	_11      = 1 + _10
	//	_100     = 1 + _11
	//	_101     = 1 + _100
	//	_111     = _10 + _101
	//	_1001    = _10 + _111
	//	_1101    = _100 + _1001
	//	_1111    = _10 + _1101
	//	_11110   = 2*_1111
	//	_11111   = 1 + _11110
	//	_111110  = 2*_11111
	//	_111111  = 1 + _111110
	//	_1111110 = 2*_111111
	//	i20      = _1111110 << 6 + _1111110
	//	x18      = i20 << 5 + _111111
	//	x31      = x18 << 13 + i20 + 1
	//	i42      = 2*x31
	//	i44      = i42 << 2
	//	i140     = ((i44 << 32 + i44) << 29 + i42) << 33
	//	i150     = ((i44 + i140 + _111) << 4 + _111) << 3
	//	i170     = ((1 + i150) << 11 + _1111) << 6 + _11111
	//	i183     = ((i170 << 5 + _1101) << 3 + _11) << 3
	//	i198     = ((1 + i183) << 7 + _111) << 5 + _11
	//	i219     = ((i198 << 9 + _101) << 5 + _101) << 5
	//	i231     = ((_1101 + i219) << 5 + _1001) << 4 + _1101
	//	i244     = ((i231 << 2 + _11) << 7 + _111111) << 2
	//	i262     = ((1 + i244) << 10 + _1001) << 5 + _111
	//	i277     = ((i262 << 5 + _111) << 4 + _101) << 4
	//	return     ((_101 + i277) << 9 + _1001) << 5 + 1
	//
	var z = new(SM2P256OrderElement).Set(e)
	var t0 = new(SM2P256OrderElement)
	var t1 = new(SM2P256OrderElement)
	var t2 = new(SM2P256OrderElement)
	var t3 = new(SM2P256OrderElement)
	var t4 = new(SM2P256OrderElement)
	var t5 = new(SM2P256OrderElement)
	var t6 = new(SM2P256OrderElement)
	var t7 = new(SM2P256OrderElement)
	var t8 = new(SM2P256OrderElement)
	var t9 = new(SM2P256OrderElement)

	t2.Square(x)
	t3.Mul(x, t2)
	t4.Mul(x, t3)
	t0.Mul(x, t4)
	t1.Mul(t2, t0)
	z.Mul(t2, t1)
	t4.Mul(t4, z)
	t6.Mul(t2, t4)
	t2.Square(t6)
	t5.Mul(x, t2)
	t2.Square(t5)
	t2.Mul(x, t2)
	t7.Square(t2)
	t8.Square(t7)
	for s := 1; s < 6; s++ {
		t8.Square(t8)
	}
	t7.Mul(t7, t8)
	t8.Square(t7)
	for s := 1; s < 5; s++ {
		t8.Square(t8)
	}
	t8.Mul(t2, t8)
	for s := 0; s < 13; s++ {
		t8.Square(t8)
	}
	t7.Mul(t7, t8)
	t7.Mul(x, t7)
	t8.Square(t7)
	t7.Square(t8)
	for s := 1; s < 2; s++ {
		t7.Square(t7)
	}
	t9.Square(t7)
	for s := 1; s < 32; s++ {
		t9.Square(t9)
	}
	t9.Mul(t7, t9)
	for s := 0; s < 29; s++ {
		t9.Square(t9)
	}
	t8.Mul(t8, t9)
	for s := 0; s < 33; s++ {
		t8.Square(t8)
	}
	t7.Mul(t7, t8)
	t7.Mul(t1, t7)
	for s := 0; s < 4; s++ {
		t7.Square(t7)
	}
	t7.Mul(t1, t7)
	for s := 0; s < 3; s++ {
		t7.Square(t7)
	}
	t7.Mul(x, t7)
	for s := 0; s < 11; s++ {
		t7.Square(t7)
	}
	t6.Mul(t6, t7)
	for s := 0; s < 6; s++ {
		t6.Square(t6)
	}
	t5.Mul(t5, t6)
	for s := 0; s < 5; s++ {
		t5.Square(t5)
	}
	t5.Mul(t4, t5)
	for s := 0; s < 3; s++ {
		t5.Square(t5)
	}
	t5.Mul(t3, t5)
	for s := 0; s < 3; s++ {
		t5.Square(t5)
	}
	t5.Mul(x, t5)
	for s := 0; s < 7; s++ {
		t5.Square(t5)
	}
	t5.Mul(t1, t5)
	for s := 0; s < 5; s++ {
		t5.Square(t5)
	}
	t5.Mul(t3, t5)
	for s := 0; s < 9; s++ {
		t5.Square(t5)
	}
	t5.Mul(t0, t5)
	for s := 0; s < 5; s++ {
		t5.Square(t5)
	}
	t5.Mul(t0, t5)
	for s := 0; s < 5; s++ {
		t5.Square(t5)
	}
	t5.Mul(t4, t5)
	for s := 0; s < 5; s++ {
		t5.Square(t5)
	}
	t5.Mul(z, t5)
	for s := 0; s < 4; s++ {
		t5.Square(t5)
	}
	t4.Mul(t4, t5)
	for s := 0; s < 2; s++ {
		t4.Square(t4)
	}
	t3.Mul(t3, t4)
	for s := 0; s < 7; s++ {
		t3.Square(t3)
	}
	t2.Mul(t2, t3)
	for s := 0; s < 2; s++ {
		t2.Square(t2)
	}
	t2.Mul(x, t2)
	for s := 0; s < 10; s++ {
		t2.Square(t2)
	}
	t2.Mul(z, t2)
	for s := 0; s < 5; s++ {
		t2.Square(t2)
	}
	t2.Mul(t1, t2)
	for s := 0; s < 5; s++ {
		t2.Square(t2)
	}
	t1.Mul(t1, t2)
	for s := 0; s < 4; s++ {
		t1.Square(t1)
	}
	t1.Mul(t0, t1)
	for s := 0; s < 4; s++ {
		t1.Square(t1)
	}
	t0.Mul(t0, t1)
	for s := 0; s < 9; s++ {
		t0.Square(t0)
	}
	z.Mul(z, t0)
	for s := 0; s < 5; s++ {
		z.Square(z)
	}
	z.Mul(x, z)
	return e.Set(z)
}
