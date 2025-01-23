package main

import (
	"fmt"      // Paquete para imprimir mensajes en consola
	"sync"     // Paquete para sincronización de goroutines (Mutex, Cond, WaitGroup)
)

// Estructura FizzBuzz
type FizzBuzz struct {
	n       int        // Número límite hasta donde queremos ejecutar FizzBuzz
	counter int        // Contador actual que se va incrementando
	mu      sync.Mutex // Mutex para manejar la exclusión mutua entre goroutines
	cond    *sync.Cond // Condición para manejar señales entre goroutines
}

// Constructor para inicializar una nueva instancia de FizzBuzz
func NewFizzBuzz(n int) *FizzBuzz {
	// Creamos una instancia de FizzBuzz con el límite `n` y el contador inicializado en 1
	fb := &FizzBuzz{n: n, counter: 1}

	// Inicializamos la condición asociada al Mutex
	fb.cond = sync.NewCond(&fb.mu) // `&` obtiene la dirección de memoria del Mutex `mu`
	return fb
}

// Método para imprimir "fizz" si el número actual es divisible por 3 pero no por 5
func (fb *FizzBuzz) fizz(printFizz func()) {
	for { // Bucle infinito que dependerá de condiciones internas para detenerse
		fb.mu.Lock() // Bloqueamos el mutex para garantizar exclusión mutua
		// Mientras el número actual no cumple las condiciones de "fizz"
		for fb.counter <= fb.n && (fb.counter%3 != 0 || fb.counter%5 == 0) {
			fb.cond.Wait() // Esperamos una señal de otras goroutines
		}
		// Si ya hemos procesado todos los números, desbloqueamos y salimos
		if fb.counter > fb.n {
			fb.mu.Unlock()
			return
		}
		// Ejecutamos la función para imprimir "fizz"
		printFizz()
		fb.counter++        // Incrementamos el contador
		fb.cond.Broadcast() // Despertamos a todas las goroutines que están esperando
		fb.mu.Unlock()      // Desbloqueamos el mutex
	}
}

// Método para imprimir "buzz" si el número actual es divisible por 5 pero no por 3
func (fb *FizzBuzz) buzz(printBuzz func()) {
	for {
		fb.mu.Lock()
		// Mientras el número actual no cumple las condiciones de "buzz"
		for fb.counter <= fb.n && (fb.counter%5 != 0 || fb.counter%3 == 0) {
			fb.cond.Wait() // Esperamos una señal de otras goroutines
		}
		if fb.counter > fb.n {
			fb.mu.Unlock()
			return
		}
		printBuzz()
		fb.counter++
		fb.cond.Broadcast() // Notificamos a otras goroutines
		fb.mu.Unlock()
	}
}

// Método para imprimir "fizzbuzz" si el número actual es divisible tanto por 3 como por 5
func (fb *FizzBuzz) fizzbuzz(printFizzBuzz func()) {
	for {
		fb.mu.Lock()
		// Mientras el número actual no cumple las condiciones de "fizzbuzz"
		for fb.counter <= fb.n && (fb.counter%3 != 0 || fb.counter%5 != 0) {
			fb.cond.Wait() // Esperamos una señal
		}
		if fb.counter > fb.n {
			fb.mu.Unlock()
			return
		}
		printFizzBuzz()
		fb.counter++
		fb.cond.Broadcast() // Notificamos a otras goroutines
		fb.mu.Unlock()
	}
}

// Método para imprimir números si no son divisibles por 3 ni por 5
func (fb *FizzBuzz) number(printNumber func(int)) {
	for {
		fb.mu.Lock()
		// Mientras el número actual es divisible por 3 o por 5
		for fb.counter <= fb.n && (fb.counter%3 == 0 || fb.counter%5 == 0) {
			fb.cond.Wait() // Esperamos una señal
		}
		if fb.counter > fb.n {
			fb.mu.Unlock()
			return
		}
		printNumber(fb.counter) // Imprimimos el número actual
		fb.counter++
		fb.cond.Broadcast() // Notificamos a otras goroutines
		fb.mu.Unlock()
	}
}

func main() {
	n := 15                  // Límite hasta el que queremos ejecutar FizzBuzz
	fb := NewFizzBuzz(n)     // Creamos una nueva instancia de FizzBuzz

	var wg sync.WaitGroup    // WaitGroup para esperar a que todas las goroutines terminen
	wg.Add(4)                // Añadimos 4 al contador, una por cada goroutine

	// Lanzamos la goroutine para manejar "fizz"
	go func() {
		defer wg.Done() // Disminuimos el contador del WaitGroup al finalizar
		fb.fizz(func() { fmt.Println("fizz") })
	}()
	// Lanzamos la goroutine para manejar "buzz"
	go func() {
		defer wg.Done()
		fb.buzz(func() { fmt.Println("buzz") })
	}()
	// Lanzamos la goroutine para manejar "fizzbuzz"
	go func() {
		defer wg.Done()
		fb.fizzbuzz(func() { fmt.Println("fizzbuzz") })
	}()
	// Lanzamos la goroutine para manejar números
	go func() {
		defer wg.Done()
		fb.number(func(x int) { fmt.Println(x) })
	}()

	wg.Wait() // Esperamos a que todas las goroutines terminen
}
