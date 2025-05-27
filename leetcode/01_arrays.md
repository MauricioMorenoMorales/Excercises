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

> Algoritmo para eliminar cierto valor, este algoritmo es clave aprenderlo
!!!
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

# En golang puedes hacer swaps facilmente
```go
			nums[left], nums[right] = nums[right], nums[left]
```

# Arrays bidimencionales
En golang puedes seguir esta estrategia para moverte dentro de un array e ir cambiando de direcciones
```go
func spiralOrder(matrix [][]int) []int {
    var response []int
    // Matrix dimensions
    rightX, bottomY := len(matrix[0])-1, len(matrix)-1
    leftX, topY := 0, 0
    totalElements := (rightX + 1) * (bottomY + 1)

    // pointers
    x, y := 0, 0

    // Direction variable: 0 = right, 1 = down, 2 = left, 3 = up
    direction := 0

    // Logic
    for len(response) < totalElements {
        // Append the current element to the response
        response = append(response, matrix[y][x])

        // Move in the current direction
        switch direction {
        case 0: // Moving right
            x++
            if x > rightX {
                x--
                y++
                topY++
                direction = 1 // Change direction to down
            }
        case 1: // Moving down
            y++
            if y > bottomY {
                y--
                x--
                rightX--
                direction = 2 // Change direction to left
            }
        case 2: // Moving left
            x--
            if x < leftX {
                x++
                y--
                bottomY--
                direction = 3 // Change direction to up
            }
        case 3: // Moving up
            y--
            if y < topY {
                y++
                x++
                leftX++
                direction = 0 // Change direction to right
            }
        }
    }

    return response
}
```

# Strings
Estudiar el siguiente algoritmo para encontrar prefijos
[flower, flowers flowerist] = flower

```go
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, value := range strs[1:] {
		for len(prefix) > 0 && len(value) < len(prefix) || value[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}

	return prefix
}
```
#

```go

[1,2,3,4,5,6,7]
 0 1 2 3 4 5 6
         ^
len = 7
len-k = 4
right := len(nums)-k


```