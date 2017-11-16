// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package common

// Common types used across all av packages.

// AVRational exposes the numerator and denominator part of the undrelying 'Rational' structure.
type AVRational struct {
	Num int
	Den int
}

type PixelFormat int

const (
	AV_PIX_FMT_RGB24 PixelFormat = 3
	AV_PIX_FMT_RGBA              = 28
)
