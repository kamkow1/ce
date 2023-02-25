package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func replaceAt(in string, r rune, i int) string {
  out := []rune(in)
  out[i] = r
  return string(out)
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

  textBuffer := getInitialFile()
  buffer := NewBuffer(textBuffer)
  editor := NewEditor()

  for {
    event := screen.PollEvent()
    w, h := screen.Size()

    switch event := event.(type) {
    case *tcell.EventResize:
      screen.Sync()
    case *tcell.EventKey:
      switch event.Key() {
      case tcell.KeyRune:
        switch editor.Mode {
        case ModeView:
          switch event.Rune() {
          case 'i':
            editor.Mode = ModeEdit
          }
        case ModeEdit:
          if cursor.X == len(buffer.Lines[cursor.Y-1]) {
            buffer.Lines[cursor.Y-1] += " "
          }
          oldLine := buffer.Lines[cursor.Y-1]
          buffer.Lines[cursor.Y-1] = oldLine[:cursor.X-1] + string(event.Rune()) + oldLine[cursor.X-1:]
          cursor.X += 1
        }
      case tcell.KeyUp:
        if cursor.Y > 1 {
          cursor.Y -= 1
        }
        
        if cursor.X > len(buffer.Lines[cursor.Y-1]) {
          cursor.X = len(buffer.Lines[cursor.Y-1])
        }
      case tcell.KeyDown:
        if cursor.X > len(buffer.Lines[cursor.Y]) {
          cursor.X = len(buffer.Lines[cursor.Y])
        }

        if cursor.Y < h {
          cursor.Y += 1
        }
      case tcell.KeyLeft:
        if cursor.X > 1 {
          cursor.X -= 1
        }

      case tcell.KeyRight:
        if cursor.X < w && cursor.X < len(buffer.Lines[cursor.Y-1]) {
          cursor.X += 1
        }
      case tcell.KeyEscape, tcell.KeyCtrlC:
        return
      case tcell.KeyCtrlI:
        editor.Mode = ModeView
      }
    }

    buffer.Display(screen, defaultStyle)
    screen.SetCursorStyle(cursor.Style)
    screen.ShowCursor(cursor.X, cursor.Y)
    screen.Show()
  }
}
