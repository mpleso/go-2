// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id Uint64 -d VecType=Uint64Vec -d Type=uint64 vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type Uint64Vec []uint64

func (p *Uint64Vec) Resize(n uint) {
	c := Index(cap(*p))
	l := Index(len(*p)) + Index(n)
	if l > c {
		c = NextResizeCap(l)
		q := make([]uint64, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *Uint64Vec) validate(new_len uint, zero *uint64) *uint64 {
	c := Index(cap(*p))
	lʹ := Index(len(*p))
	l := Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l >= lʹ {
			*p = (*p)[:l]
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *Uint64Vec) validateSlowPath(zero *uint64,
	c, l, lʹ Index) *uint64 {
	if l > c {
		cNext := NextResizeCap(l)
		q := make([]uint64, cNext, cNext)
		copy(q, *p)
		if zero != nil {
			for i := c; i < cNext; i++ {
				q[i] = *zero
			}
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *Uint64Vec) Validate(i uint) *uint64 {
	return p.validate(i+1, (*uint64)(nil))
}

func (p *Uint64Vec) ValidateInit(i uint, zero uint64) *uint64 {
	return p.validate(i+1, &zero)
}

func (p *Uint64Vec) ValidateLen(l uint) (v *uint64) {
	if l > 0 {
		v = p.validate(l, (*uint64)(nil))
	}
	return
}

func (p *Uint64Vec) ValidateLenInit(l uint, zero uint64) (v *uint64) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p Uint64Vec) Len() uint { return uint(len(p)) }
