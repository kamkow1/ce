package main

import (
	"log"
	"os"
  "path/filepath"

	"github.com/gdamore/tcell/v2"
)

var editor *Editor
var cursor *Cursor
var buffer *Buffer
var ui *UI

func getInitialFile() string {
  filename := os.Args[1]
  b, err := os.ReadFile(filename)
  if err != nil {
    log.Fatalf("%+v", err)
  }
  return string(b)
}

func handleTextKey(key rune) {
  switch editor.Mode {
  case ModeView:
    switch key {
    case 'i':
      editor.Mode = ModeEdit
    }
  case ModeEdit:
    if cursor.X == len(buffer.Lines[cursor.Y-StartYPos]) {
      buffer.Lines[cursor.Y-StartYPos] += " "
    }
    oldLine := buffer.Lines[cursor.Y-StartYPos]
    buffer.Lines[cursor.Y-StartYPos] = oldLine[:cursor.X-StartXPos] + string(key) + oldLine[cursor.X-StartXPos:]
    cursor.X += 1
  }
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
  cursor = NewCursor()
  cursor.SetStyle(tcell.CursorStyleSteadyBlock)

  quit := func() {
    maybePanic := recover()
    screen.Fini()
    if maybePanic != nil {
      panic(maybePanic)
    }
  }
  defer quit()

  editor = NewEditor()
  ui = NewUI()
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

  buffer = NewBuffer(textBuffer)

  for {
    event := screen.PollEvent()
    w, h := screen.Size()

    switch event := event.(type) {
    case *tcell.EventResize:
      screen.Sync()
    case *tcell.EventKey:
      switch event.Key() {
      case tcell.KeyRune:
        handleTextKey(event.Rune())
      case tcell.KeyEnter:
        switch editor.Mode {
        case ModeView:
          cursor.MoveDown(buffer, h)
        case ModeEdit:
          buffer.Lines, err = ArrayInsert(buffer.Lines, cursor.Y-StartYPos, "")
          if err != nil {
            log.Fatalf("%+v", err)
          }
        }
      case tcell.KeyUp:
        cursor.MoveUp(buffer)
      case tcell.KeyDown:
        cursor.MoveDown(buffer, h)
      case tcell.KeyLeft:
        cursor.MoveLeft()
      case tcell.KeyRight:
        cursor.MoveRight(buffer, w)
      case tcell.KeyEscape, tcell.KeyCtrlC:
        return
      case tcell.KeyCtrlI:
        editor.Mode = ModeView
      }
    }

    screen.Clear()

    ui.ShowStatusBar(screen, defaultStyle, *cursor, len(textBuffer), *editor)
    ui.DisplayLineNumbers(screen, defaultStyle, len(buffer.Lines))

    buffer.Display(screen, defaultStyle)

    screen.SetCursorStyle(cursor.Style)
    screen.ShowCursor(cursor.X, cursor.Y)

    screen.Show()
  }
}
