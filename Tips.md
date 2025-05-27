> Puedes contar los digitos de un numero dividiendolo entre 10 por cada division cuentas un counter++ hasta llegar a 0
> Antes de hacer el algoritmo revisa la respuesta si se comporta como la respuesta que tienes en mente
> Puedes hacer tu descripcion del bug antes de empezarlo escribiendo un poco sobre él


# Golang tips
> Forma eficiente de recorrer o hacer maps en golang
```golang
func map(nums []int) []int {
  result := []int{}
  for _, value := range nums {
    result = append(result, function)
  }
  return result
}
```
# Arrays techniques
> Piensa si dos punteros pueden resolver el problema
> Piensa si dos punteros en distintas posiciones al inicio pueden funcionar
> Puedes hacer swap entre dos punteros
> Manejar distintas velocidades de punteros, o dejar un puntero en el punto que se hara un swap o se fijara info
> Prueba si funciona con sliding window tu algoritmo