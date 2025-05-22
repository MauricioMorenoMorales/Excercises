Forma en la que se hace un insertion en java
Primero se agrega a la derecha un nuevo elemento y pasamos todos los valores al lado
```java
for (int i = len(arr)-1; i >= $target; i--) {
  intArra[i + 1] = intArray[i]
}
```
Otro ejemplo de insertar, en este caso con un array fijo los valores que no caben salen de la respuesta
```go
func duplicateZeros(arr []int) {
	insertZero := func(index int) {
		for i := len(arr) - 1; i > index; i-- {
			arr[i] = arr[i-1]
		}
		arr[index] = 0
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			insertZero(i)
			i++ // Saltar el siguiente índice porque acabamos de insertar un 0
		}
	}
}

	insertZero := func(index int) {
		for i := len(arr) - 1; i > index; i-- {
			arr[i] = arr[i-1]
		}
		arr[index] = 0
	}
// En especial pon atencion aqui
```

# Eliminar elementos del array

Para eliminar un elemento del final del  array no necesitas eliminar el valor solo cambiar el indice propio de la estructura de datos con un index--

Para eliminar un elemento en determinado index tienes que usar esta funcion

```go
func deleteIndex(array []int, targetIndex int) {
	for i := targetIndex+1; i < len(array); i++ {
		array[i-1] = array[i]
	}
}
```

> Algoritmo para eliminar cierto valor 

```go
func removeElement(nums []int, val int) int {
	k := 0

	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k += 1
		}
	}
	return k
}

```

> Algoritmo eliminar duplicados
```go
func removeDuplicates(nums []int) int {
	currentValue := nums[0]
	k := 1

	for i, element := range nums {
		if currentValue != element {
			nums[k] = nums[i]
			currentValue = nums[i]
			k++
		}
	}
	return k
}
```

# Searching data
Para encontrar data se puede hacer una busqueda linear ir en cada elemento y en el peor de los escenarios seria una operacion O(N)
Binary search es util pero solo con data sorteada, no conviene sortear primero ya que sortear requiere mas tiempo que una busqueda linear, a menos que vayas a buscar informacion muy seguido en ese caso si conviene sortear la data primero y hacer busquedas despues