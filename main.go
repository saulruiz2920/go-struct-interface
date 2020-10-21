package main

import "fmt"

type Imagen struct {
    titulo string
    formato string
    canal string
}

type Audio struct {
    titulo string
    formato string
    duracion float32
}

type Video struct {
    titulo string
    formato string
    frames string
}

type Multimedia interface {
  mostrar()
}

func mostrar(m Multimedia) {
  m.mostrar()
}

func (i Imagen) mostrar() {
  fmt.Printf("Titulo: %s\n", i.titulo)
  fmt.Printf("Formato: %s\n", i.formato)
  fmt.Printf("Canal: %s\n", i.canal)
}

func (a Audio) mostrar() {
  fmt.Printf("Titulo: %s\n", a.titulo)
  fmt.Printf("Formato: %s\n", a.formato)
  fmt.Printf("Duracion: %f\n", a.duracion)
}

func (v Video) mostrar() {
  fmt.Printf("Titulo: %s\n", v.titulo)
  fmt.Printf("Formato: %s\n", v.formato)
  fmt.Printf("frames: %s\n", v.frames)
}

type ContenidoWeb struct {
  multimedias []Multimedia
}

func (cw ContenidoWeb) mostrar() {
  for _, cw := range cw.multimedias {
    fmt.Println("")
    cw.mostrar()
  }
}

func capImagen() Imagen {
  i := Imagen{}
  fmt.Print("Titulo: ")
  fmt.Scan(&i.titulo)
  fmt.Print("Formato: ")
  fmt.Scan(&i.formato)
  fmt.Print("Canal: ")
  fmt.Scan(&i.canal)
  return i
}

func capAudio() Audio {
  a := Audio{}
  fmt.Print("Titulo: ")
  fmt.Scan(&a.titulo)
  fmt.Print("Formato: ")
  fmt.Scan(&a.formato)
  fmt.Print("Duracion: ")
  fmt.Scan(&a.duracion)
  return a
}

func capVideo() Video {
  v := Video{}
  fmt.Print("Titulo: ")
  fmt.Scan(&v.titulo)
  fmt.Print("Formato: ")
  fmt.Scan(&v.formato)
  fmt.Print("Frames: ")
  fmt.Scan(&v.frames)
  return v
}


func main() {
  cw := ContenidoWeb{}
  opt := -1
  for opt != 0 {
    fmt.Println("\tMENU")
    fmt.Println("1.- Cap imagen")
    fmt.Println("2.- Cap audio")
    fmt.Println("3.- Cap video")
    fmt.Println("4.- Mostrar todo")
    fmt.Println("0.- Salir")
    fmt.Scan(&opt)
    switch opt {
      case 1:
        cw.multimedias = append(cw.multimedias, capImagen())
      case 2:
        cw.multimedias = append(cw.multimedias, capAudio())
      case 3:
        cw.multimedias = append(cw.multimedias, capVideo())
      case 4:
        cw.mostrar()
      
    }
    
  }

}