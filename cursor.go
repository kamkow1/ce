package main

import "github.com/gdamore/tcell/v2"

type Cursor struct {
  X, Y int // x, y coordinates for easy rendering
  AbsPos uint64 // absolute position within the buffer
  Style tcell.CursorStyle 
}

func NewCursor() *Cursor {
  return &Cursor{X: 1, Y: 1, AbsPos: 0}
}

func (c *Cursor) SetStyle(style tcell.CursorStyle) {
  c.Style = style
}

