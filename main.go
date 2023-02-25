package main

import "log"

import "github.com/gdamore/tcell"

func main() {
  defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
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

  quit := func() {
    maybePanic := recover()
    screen.Fini()
    if maybePanic != nil {
      panic(maybePanic)
    }
  }
  defer quit()

  for {
    screen.Show()
    event := screen.PollEvent()
    switch event := event.(type) {
    case *tcell.EventResize:
      screen.Sync()
    case *tcell.EventKey:
      key := event.Key()
      if key == tcell.KeyEscape || key == tcell.KeyCtrlC {
        return
      }
    }
  }
}
