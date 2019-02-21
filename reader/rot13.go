package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func Rot13(a byte) (b byte) {
  if a >= 'a' && a <= 'm' {
    b = a + 13
  } else if a >= 'n' && a <= 'z' {
    b = a - 13
  } else if a >= 'A' && a <= 'M' {
    b = a + 13
  } else if a >= 'N' && a <= 'Z' {
    b = a - 13
  } else {
    b = a
  }

  return b
}

func (a rot13Reader) Read (p []byte) (n int, err error){
  n, e := a.r.Read(p)
  if e != io.EOF {

    for i := 0; i < len(p); i++{
      p[i] = Rot13(p[i])
    }
  }

  return n, e
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}

