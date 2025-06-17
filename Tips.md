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
> A veces el swap se puede hacer entre dos punteros en el mismo lugar, empiezas ambos punteros en el mismo lugar y dependiendo las condiciones vas dejando uno atras mientras llenas los otrso
> Manejar distintas velocidades de punteros, o dejar un puntero en el punto que se hara un swap o se fijara info
> Prueba si funciona con sliding window tu algoritmo

# Iterations
> Si te enfrentas a ciclos infinitos preguntate si nay ciclos y si puedes usar la tecnica de liebre y tortuga

# Numbers
> Si trabajaras con digitos no necesitas hacer split o stringify, puedes usar esta tecnica

```go
func getNext(n int) int {
    totalSum := 0
    for n > 0 {
        digit := n % 10 // Extrae el valor del primer digito
        n /= 10 // Reduce el valor del primer digito del numero original
        totalSum += digit * digit
    }
    return totalSum
}
```


# Comparaciones

Cuando comparas cosas puedes usar la tecnica de transformar ambos datos a una abstraccion y compararlos como en este ejemplo

```go
func transformString(s string) string {
	indexMapping := make(map[rune]int)
	var newStr []string

	for i, c := range s {
		if _, exists := indexMapping[c]; !exists {
			indexMapping[c] = i
		}
		newStr = append(newStr, strconv.Itoa(indexMapping[c]))
	}

	return strings.Join(newStr, " ")
}

func isIsomorphic(s string, t string) bool {
	return transformString(s) == transformString(t)
}

```

# Multiples respuestas
Considera guardarlos en un array o pensar en arrays en los datos que usas








# Goals
🟢 Nivel 1: Fundamentos – Muy recomendados para empezar
The LeetCode Beginner’s Guide 🧭
√ Array and String
√ Arrays 101
√? Hash Table
Queue & Stack
√ Linked List

🔵 Nivel 2: Intermedio – Introduce técnicas clave
Binary Tree
Binary Search Tree
N-ary Tree
Binary Search
Sorting
Recursion
Recursion II
Bit Manipulation

🟣 Nivel 3: Avanzado – Conceptos más complejos
Heap
Graph
Trie
Dynamic Programming
Decision Tree

🟠 Extras (no algorítmicos pero útiles en tech interviews o backend)
SQL Language
Machine Learning 101
Systems Design
