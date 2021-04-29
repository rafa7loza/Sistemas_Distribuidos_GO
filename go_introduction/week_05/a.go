package main

import "fmt"

type Multimedia interface {
  mostrar();
}

type Video struct {
  titulo, formato string;
  frames int;
}

type Audio struct {
  titulo, formato string;
  duracion int;
}

type Imagen struct {
  titulo, formato string;
  canales int;
}

type ContenidoWeb struct {
  contenido []Multimedia;
}

func (v Video) mostrar(){
  fmt.Printf("Titulo: %s.%s\nFrames: %d\n\n",
    v.titulo, v.formato, v.frames);
}

func (a Audio) mostrar(){
  fmt.Printf("Titulo: %s.%s\nDuración: %ds\n\n",
    a.titulo, a.formato, a.duracion);
}

func (img Imagen) mostrar(){
  fmt.Printf("Titulo: %s.%s\nCanales: %d\n\n",
    img.titulo, img.formato, img.canales);
}

func wrap(vs ...interface{}) []interface{} {
    return vs
}


func leer(t string) (string, string, int) {
  var a,b string;
  var c int;

  fmt.Print("Nombre: ")
  fmt.Scanf("%s", &a);
  fmt.Print("Formato: ")
  fmt.Scanf("%s", &b);
  switch t {
  case "video":
    fmt.Print("Frames: ")
  case "imagen":
    fmt.Print("Canales: ")
  case "audio":
    fmt.Print("duracion: ")
  }
  fmt.Scanf("%d", &c);

  return a,b,c;
}

func menu() {
  fmt.Println("\tMenu\na) Agregar imagen\nb) Agregar Video\nc) Agregar audio\nd) Mostrar todos\ne) Salir");
}

func main() {

  web := ContenidoWeb{};
  var opt string;
  var obj Multimedia;

  for {
    menu();
    fmt.Scan(&opt);
    obj = nil
    if opt == "e" { break }
    switch opt {
    case "a":
      a,b,c := leer("imagen")
      obj = Imagen{ a,b,c }
    case "b":
      a,b,c := leer("video")
      obj = Video{ a,b,c }
    case "c":
      a,b,c := leer("audio")
      obj = Audio{ a,b,c }
    case "d":
      for _,v := range web.contenido {
        v.mostrar();
      }
    default:
      fmt.Println("Opcion inválida");
    }

    if obj != nil { web.contenido = append(web.contenido, obj)}
  }

}
