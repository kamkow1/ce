package main

type EditorMode = int
const (
  ModeView EditorMode = iota
  ModeEdit
)

type Editor struct {
  Mode EditorMode 
}

func NewEditor() *Editor {
  return &Editor{Mode: ModeView}
}
