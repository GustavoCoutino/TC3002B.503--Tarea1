# Tarea 1

Implementacion generica de Stack, Queue, y HashMap en Go.

## Estructura del proyecto

```
.
├── stack/        # Stack (LIFO)
├── queue/        # Queue (FIFO)
├── hashmap/      # HashMap
├── utils/        # FNV hashing
└── main.go       # Ejemplo de uso
```

## Estructura de datos

### Stack

Un stack LIFO implementado con un Go slice.

```go
st := stack.New[int]()
st.Push(1)
st.Push(2)

top, err := st.Top()  // 2
st.Pop()
st.Empty()            // false
```

| Metodo             | Descripcion                                      |
| ------------------ | ------------------------------------------------ |
| `New[T]()`         | Crea un stack vacio                              |
| `Push(val T)`      | Empuja un valor al principio del stack           |
| `Pop() error`      | Remueve el ultimo elemento agregado del stack    |
| `Top() (T, error)` | Regresa el elemento en la cima del stack         |
| `Empty() bool`     | Regresa verdadero si el stack no tiene elementos |

---

### Queue

Un queue FIFO implementado con un Go slice.

```go
q := queue.New[int]()
q.Push(1)
q.Push(2)

front, err := q.Front()  // 1
back, err  := q.Back()   // 2
q.Pop()
q.Empty()                // false
```

| Metodo               | Descripcion                           |
| -------------------- | ------------------------------------- |
| `New[T]()`           | Crea un queue vacio                   |
| `Push(val T)`        | Pone en cola un elemento              |
| `Pop() error`        | Borrar el primer elemento.            |
| `Front() (T, error)` | Regresa el primero elemento           |
| `Back() (T, error)`  | Regresa el ultimo elemento            |
| `Empty() bool`       | Regresa verdadero si no hay elementos |

---

### HashMap

Un HashMap usando separate chaining para manejar colisones. Automaticamente aumenta su capacidad por el doble cuando esta llena. Acepte cualquier tipo `comparable` y una funcion de hashing personalizada.

```go
ht := hashmap.New[string, int](5, utils.FNVHash)
ht.Insert("Hello", 1)
ht.Insert("Goodbye", 2)

k, v, err := ht.Get("Hello")  // "Hello", 1, nil
ht.Remove("Hello")
ht.Size()                      // 1
```

| Method                                         | Description                                                        |
| ---------------------------------------------- | ------------------------------------------------------------------ |
| `New[K, V](size int, hashFunc func(K) uint32)` | Crea un nuevo hashmap con tamaño definido y una funcion de hashing |
| `Insert(key K, value V) error`                 | Inserta o actualiza un par key-value                               |
| `Get(key K) (K, V, error)`                     | Obtiene un valor usando una llave                                  |
| `Remove(key K) error`                          | Remueve un par key-value                                           |
| `Size() int`                                   | Regresa el numero de elementos almacenados en el mapa              |

## Utils

`utils.FNVHash` es una funcion de hash de string basada en el algoritmo 32-bit FNV-1a.

```go
hash := utils.FNVHash("Hello")  // 4116459851
```

## Ejecutar

```bash
# Ejecuta el archivo main
go run main.go

# Ejecuta todas las pruebas
go test ./...
```
