package main

import (
  "os"
  "image"
  _ "image/png"
  _ "image/jpeg"
)

func OpenImageFile(filename string) (image.Image, string, error) {
  file, err := os.Open(filename)
  HandleError(err)

  defer file.Close()

  return image.Decode(file)
}
