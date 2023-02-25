package main

import "github.com/gdamore/tcell/v2"

type Buffer struct {
  Lines []string
}

func NewBuffer(textBuffer string) *Buffer {
  b := &Buffer{}
  line := ""
  for _, ch := range textBuffer {
    if ch == '\n' {
      b.Lines = append(b.Lines, line)
      line = ""
    } else {
      line += string(ch)
    }
  }

  return b
}

func (b *Buffer) Display(screen tcell.Screen, style tcell.Style) {
  row := 1
  for _, line := range b.Lines {
    col := 1
    for _, ch := range line {
      screen.SetCell(col, row, style, ch)
      col++
    }
    screen.SetCell(col, row, style, '\n')
    row++
  }
}
