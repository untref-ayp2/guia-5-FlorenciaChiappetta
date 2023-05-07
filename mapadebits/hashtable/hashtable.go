package hashtable

//Diseñar e implementar un TAD con nombre HashTable que utilice la estrategia de hashing abierto para resolver las colisiones.
//El TAD debe ser genérico y poder parametrizarse para distintos tipos tanto de claves como de valores.

type Nodo[T int] struct {
	value   T
	puntero *Nodo[T]
}

func newNode[T int](value T) *Nodo[T] {
	return &Nodo[T]{value: value, puntero: nil}
}

type linkedlist[T int] struct {
	head *Nodo[T]
	tail *Nodo[T]
	size int
}

type HashTable[T int] struct {
	posiciones []linkedlist[T]
}

func (h *HashTable[int]) newHashTable(arreglo ...int) *HashTable[int] {
	N := 17
	for _, valor := range arreglo {
		indice := valor % N
		if h[indice] != 0 {
			h[indice].next.value = valor
		}
		h = append(h[indice], valor)
	}

	return &HashTable[int]{}
}
