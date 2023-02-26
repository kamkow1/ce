package main

import "github.com/gdamore/tcell/v2"

type Cursor struct {
  X, Y int // x, y coordinates for easy rendering
  Style tcell.CursorStyle 
}

func NewCursor() *Cursor {
  return &Cursor{X: StartXPos, Y: StartYPos}
}

func (c *Cursor) SetStyle(style tcell.CursorStyle) {
  c.Style = style
}

