package main

import (
  "strings"

  "github.com/gdamore/tcell/v2"
)

const (
  StartXPos = 6
  StartYPos = 3
)

type Buffer struct {
  Lines []string
}

func NewBuffer(textBuffer string) *Buffer {
  if strings.Count(textBuffer, "\n") == 0 {
    textBuffer += "\n"
  }

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
  row := StartYPos
  for _, line := range b.Lines {
    col := StartXPos
    for _, ch := range line {
      screen.SetCell(col, row, style, ch)
      col++
    }
    screen.SetCell(col, row, style, '\n')
    row++
  }
}
