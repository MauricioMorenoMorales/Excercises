En ciencias de la computacion un orden formal debe cumplir dos reglas
1. Ley de Tricotomia, para dos elemementos a y b deben cumplirse exactamente una de estas
a = b, a < b, a > b
2. Ley de transitividad, si a < b y b < c entonces a < c

Una inversión es un par de elementos que están en un orden incorrecto respecto al criterio elegido.

Ejemplo:
Con base en la longitud, esta lista:

```typescript
["are", "we", "sorting", "hello", "world", "learning"]
```
Tiene las siguientes inversiones:

("are", "we")

("sorting", "hello")

("sorting", "world")

Cuantas más inversiones haya, más desordenada está la lista.

Entonces, otra forma de ver un algoritmo de ordenamiento es: una serie de pasos para reducir el número de inversiones a 0.

♻️ Estabilidad de un algoritmo de ordenamiento
Un algoritmo es estable si mantiene el orden original entre elementos que son iguales según el criterio de comparación.

Ejemplo:
Secuencia original:

["hello", "world", "we", "are", "learning", "sorting"]

Ordenado por longitud:

["we", "are", "hello", "world", "sorting", "learning"] → estable

["we", "are", "world", "hello", "sorting", "learning"] → no estable

Ambos están bien ordenados, pero el primero conserva el orden de "hello" y "world" como estaban en la original.

Conceptos:
> Estabilidad de algoritmos
> Inversiones

## Swaps
En javascript no se pueden hacer swaps como en golang o python
Esta seria la tecnica comun para hacer los swaps
```typescript
const temp = lst[currentIndex];
lst[currentIndex] = lst[currentIndex - 1];
lst[currentIndex - 1] = temp;
```

## Selection Sort
Complexity: O(n^2)
Non stable algorithm, does not preserve the ordering of equal elements
```go
func selectionSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		minIndex := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		// Swap actual con mínimo
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}
}
```

## Bubble Sort
Complexity: O(n^2)
Stable algorithm, similar value elements are never swapped

```go
func bubbleSort(lst []int) {
	hasSwapped := true

	for hasSwapped {
		hasSwapped = false

		for i := 0; i < len(lst)-1; i++ {
			if lst[i] > lst[i+1] {
				// Swap elementos adyacentes
				lst[i], lst[i+1] = lst[i+1], lst[i]
				hasSwapped = true
			}
		}
	}
}
```
## Insertion Sort
Complexity: O(n^2) but better than previous algorithms
Stable algorithm
```go
func insertionSort(lst []int) {
	for i := 1; i < len(lst); i++ {
		currentIndex := i

		for currentIndex > 0 && lst[currentIndex-1] > lst[currentIndex] {
			// Intercambiar elementos fuera de orden
			lst[currentIndex], lst[currentIndex-1] = lst[currentIndex-1], lst[currentIndex]
			currentIndex--
		}
	}
}
```

# Non comparison based sorts

# Counting sort
Complexity: O(N)
Space: O(2N)

Insertion sort para elemenotos mayores o iguales a 0 (sin shift)
```go
func countingSort(numbers []int) {
	if len(numbers) == 0 {
		return
	}

	// Encontrar el valor maximo
	maximum := numbers[0]
	for _, number := range numbers {
		if number > maximum {
			maximum = number
		}
	}

	// Crear y llenar el arreglo de conteo
	counts := make([]int, maximum+1)
	for _, number := range numbers {
		counts[number]++
	}

	// Sobreescribir counts con los indices de inicio
	startingIndex := 0
	for i, count := range counts {
		counts[i] = startingIndex
		startingIndex += count
	}

	// Crear el arreglo ordenado
	sortedList := make([]int, len(numbers))
	for _, number := range numbers {
		sortedList[counts[number]] = number
		counts[number]++
	}

	// Copiar el arreglo ordenado al original
	for i := range numbers {
		numbers[i] = sortedList[i]
	}
}
```
Insertion sort para elementos negativos (con shift)
```go
func CountingSortWithShift(lst []int) {
	if len(lst) == 0 {
		return
	}

	// Encontrar el mínimo y máximo
	minVal := lst[0]
	maxVal := lst[0]
	for _, val := range lst {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}

	shift := minVal
	K := maxVal - shift

	// Crear y llenar el arreglo de conteo
	counts := make([]int, K+1)
	for _, elem := range lst {
		counts[elem-shift]++
	}

	// Calcular índices de inicio
	startingIndex := 0
	for i := range counts {
		count := counts[i]
		counts[i] = startingIndex
		startingIndex += count
	}

	// Ordenar usando counts
	sortedLst := make([]int, len(lst))
	for _, elem := range lst {
		idx := counts[elem-shift]
		sortedLst[idx] = elem
		counts[elem-shift]++
	}

	// Copiar el resultado a lst
	for i := range lst {
		lst[i] = sortedLst[i]
	}
}

```