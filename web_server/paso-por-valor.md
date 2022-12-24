Go por defecto paso los datos vor valor, 

es decir si en una funcion le enviamos una variable esta funcion hara una copia de esa variable 
para modificar la misma variable que queremos modificar necesitamos hacer el paso por referencia
utilizando & y *

func main(){

  var a :=  5
  sum(a) // no modificara nada de a
  add(&a) // si modificara nada de a

}

func sum(b int){
  b + 1
}

// indica que recibira la referencia de memori de un numero entoro y la parsera a su valor 
func add(c  *int){
  c + 1
}