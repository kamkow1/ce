package main


import (
  "fmt"
  "strconv"

  "github.com/gdamore/tcell/v2"
)

type UI struct {
  CurrentBufferName string
}

func NewUI() *UI {
  return &UI{}
}

func (ui *UI) PutText(screen tcell.Screen, style tcell.Style, startX, y int, text string) {
  for _, ch := range text {
    screen.SetCell(startX, y, style, ch)
    startX += 1
  }
}

func (ui *UI) ShowStatusBar(screen tcell.Screen, style tcell.Style, cursor Cursor, bufferLen int, editor Editor) {
  bufferName := "buffer: " + ui.CurrentBufferName

  spaces := ""
  for i := 0; i < bufferLen; i++ {
    spaces += " "
  }
  cursorInfo := fmt.Sprintf("(%d:%d)" + spaces, cursor.X - StartXPos, cursor.Y - StartYPos)

  modeInfo := "mode: "
  switch editor.Mode {
  case ModeEdit:
    modeInfo += "Edit"
  case ModeView:
    modeInfo += "View"
  }

  ui.PutText(screen, style, 1, 0, bufferName)
  ui.PutText(screen, style, len(bufferName) + 3, 0, cursorInfo)
  ui.PutText(screen, style, 1, 1, modeInfo)
}

func (ui *UI) DisplayLineNumbers(screen tcell.Screen, style tcell.Style, bufferLen int) {
  for i := 0; i < bufferLen; i++ {
    ui.PutText(screen, style, 1, i + StartYPos, strconv.Itoa(i))
  }
}

