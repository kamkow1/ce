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

func (c *Cursor) MoveDown(buffer *Buffer, h int) {
  if cursor.X > len(buffer.Lines[cursor.Y-StartYPos]) {
    cursor.X = len(buffer.Lines[cursor.Y-StartYPos])
  }

  if cursor.Y < h && cursor.Y <= len(buffer.Lines)+1 {
    cursor.Y += 1
  }
}

func (c *Cursor) MoveUp(buffer *Buffer) {
  if cursor.Y > StartYPos {
    cursor.Y -= 1
  }
        
  if cursor.X > len(buffer.Lines[cursor.Y-StartYPos]) {
    cursor.X = len(buffer.Lines[cursor.Y-StartYPos])
  }
}

func (c *Cursor) MoveLeft() {
  if cursor.X > StartXPos {
    cursor.X -= 1
  }
}
func (c *Cursor) MoveRight(buffer *Buffer, w int) {
  if cursor.X < w && cursor.X < len(buffer.Lines[cursor.Y-StartYPos]) + StartXPos {
    cursor.X += 1
  }
}
