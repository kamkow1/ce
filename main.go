package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func makeBox(screen tcell.Screen, width, height int, style tcell.Style) {
  w, h := screen.Size()
  if w == 0 || h == 0 {
    return
  }
  for row := 0; row < h; row++ {
    for col := 0; col < w; col++ {
      screen.SetCell(row, col, style, '@')
    }
  }
  screen.Show()
} 

func displayBuffer(screen tcell.Screen, style tcell.Style, buffer string) {
  row := 1
  col := 1
  for _, ch := range buffer {
    if ch == '\n' {
      row++
      col = 1
    } else {
      col++
    }
    screen.SetCell(col, row, style, ch)
  }
}

func getInitialFile() string {
  filename := os.Args[1]
  b, err := os.ReadFile(filename)
  if err != nil {
    log.Fatalf("%+v", err)
  }
  return string(b)
}

func main() {
  defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
  screen, err := tcell.NewScreen()
  if err != nil {
    log.Fatalf("%+v", err)
  }
  if err := screen.Init(); err != nil {
    log.Fatalf("%+v", err)
  }
  screen.SetStyle(defaultStyle)
  screen.EnableMouse()
  screen.Clear()

  // cursor
  cursor := NewCursor()
  cursor.SetStyle(tcell.CursorStyleSteadyBlock)

  quit := func() {
    maybePanic := recover()
    screen.Fini()
    if maybePanic != nil {
      panic(maybePanic)
    }
  }
  defer quit()

  buffer := getInitialFile()

  for {
    event := screen.PollEvent()
    switch event := event.(type) {
    case *tcell.EventResize:
      screen.Sync()
    case *tcell.EventKey:
      switch event.Key() {
      case tcell.KeyRune:
      case tcell.KeyUp:
        cursor.Y -= 1
      case tcell.KeyDown:
        cursor.Y += 1
      case tcell.KeyLeft:
        cursor.X -= 1
      case tcell.KeyRight:
        cursor.X += 1
      case tcell.KeyEscape, tcell.KeyCtrlC:
        return
      }
    }

    displayBuffer(screen, defaultStyle, buffer)
    screen.SetCursorStyle(cursor.Style)
    screen.ShowCursor(cursor.X, cursor.Y)
    screen.Show()
  }
}
