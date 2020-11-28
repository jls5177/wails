package renderer

import (
	bridge "github.com/jls5177/wails/lib/renderer/bridge"
)

// NewBridge returns a new Bridge struct
func NewBridge() *bridge.Bridge {
	return &bridge.Bridge{}
}
