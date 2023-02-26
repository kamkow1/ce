package main

import (
	"log"
	"os"
  "path/filepath"

	"github.com/gdamore/tcell/v2"
)

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

  editor := NewEditor()
  ui := NewUI()
  var textBuffer string
  if len(os.Args) > 0 {
    textBuffer = getInitialFile()
    absPath, err := filepath.Abs(os.Args[1])
    if err != nil {
      log.Fatalf("%+v", err)
    }

    ui.CurrentBufferName = absPath
  } else {
    textBuffer = "New Buffer"
  }

  buffer := NewBuffer(textBuffer)

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
          if cursor.X == len(buffer.Lines[cursor.Y-StartYPos]) {
            buffer.Lines[cursor.Y-StartYPos] += " "
          }
          oldLine := buffer.Lines[cursor.Y-StartYPos]
          buffer.Lines[cursor.Y-StartYPos] = oldLine[:cursor.X-StartXPos] + string(event.Rune()) + oldLine[cursor.X-StartXPos:]
          cursor.X += 1
        }
      case tcell.KeyUp:
        if cursor.Y > StartYPos {
          cursor.Y -= 1
        }
        
        if cursor.X > len(buffer.Lines[cursor.Y-StartYPos]) {
          cursor.X = len(buffer.Lines[cursor.Y-StartYPos])
        }
      case tcell.KeyDown:
        if cursor.X > len(buffer.Lines[cursor.Y-StartYPos]) {
          cursor.X = len(buffer.Lines[cursor.Y-StartYPos])
        }

        if cursor.Y < h && cursor.Y <= len(buffer.Lines){
          cursor.Y += 1
        }
      case tcell.KeyLeft:
        if cursor.X > StartXPos {
          cursor.X -= 1
        }
      case tcell.KeyRight:
        if cursor.X < w && cursor.X < len(buffer.Lines[cursor.Y-StartYPos]) + StartXPos {
          cursor.X += 1
        }
      case tcell.KeyEscape, tcell.KeyCtrlC:
        return
      case tcell.KeyCtrlI:
        editor.Mode = ModeView
      }
    }

    ui.ShowStatusBar(screen, defaultStyle, *cursor, len(textBuffer), *editor)
    buffer.Display(screen, defaultStyle)
    screen.SetCursorStyle(cursor.Style)
    screen.ShowCursor(cursor.X, cursor.Y)
    screen.Show()
  }
}
