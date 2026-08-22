package ssehub

import (
	"errors"

	"github.com/LYH2263/go-ssehub/internal/backpressure"
)

var (
	ErrClosed   = errors.New("ssehub: closed")
	ErrInvalid  = errors.New("ssehub: invalid")
	ErrNotFound = errors.New("ssehub: not found")
	ErrConflict = errors.New("ssehub: conflict")
	// ErrFull is the backpressure sentinel returned when a room subscriber's
	// send queue is full. It is the same value as internal/backpressure.ErrFull,
	// re-exported at the package boundary so that callers such as an edge
	// gateway can identify it via errors.Is across both Room.Publish and
	// Hub.Fanout without importing the internal package.
	ErrFull = backpressure.ErrFull
)
