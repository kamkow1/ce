package main

import "errors"

func ArrayInsert[T interface{}](og []T, index int, value T) ([]T, error) {
  if index < 0 {
    return nil, errors.New("Index cannot be less than 0")
  }
  if index >= len(og) {
    return append(og, value), nil
  }

  og = append(og[:index+1], og[index:]...)
  og[index] = value
  return og, nil
}
