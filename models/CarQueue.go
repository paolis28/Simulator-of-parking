package models

import (
	"container/list"
	"sync"
)

// CarQueue representa una cola de autos con acceso seguro mediante un mutex
type CarQueue struct {
	mutex sync.Mutex // Mutex para asegurar operaciones concurrentes seguras
	queue *list.List // Lista enlazada doble para almacenar los autos en la cola
}

// NewCarQueue crea y devuelve una nueva instancia de CarQueue
func NewCarQueue() *CarQueue {
	return &CarQueue{
		queue: list.New(), // Inicializa una nueva lista enlazada vacía
	}
}

// Enqueue añade un auto al final de la cola
func (cq *CarQueue) Enqueue(car *Car) {
	cq.mutex.Lock()         // Adquiere el lock para acceso seguro
	defer cq.mutex.Unlock() // Libera el lock al final del método
	cq.queue.PushBack(car)  // Añade el auto al final de la cola
}

// Dequeue elimina y devuelve el primer auto de la cola
func (cq *CarQueue) Dequeue() *Car {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	if cq.queue.Len() == 0 {
		return nil // Si la cola está vacía, retorna nil
	}
	element := cq.queue.Front() // Obtiene el primer elemento de la cola
	cq.queue.Remove(element)    // Elimina el primer elemento de la cola
	return element.Value.(*Car) // Retorna el valor del elemento como *Car
}

// First devuelve el primer auto de la cola sin eliminarlo
func (cq *CarQueue) First() *Car {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	if cq.queue.Len() == 0 {
		return nil // Si la cola está vacía, retorna nil
	}
	element := cq.queue.Front() // Obtiene el primer elemento de la cola
	return element.Value.(*Car) // Retorna el valor del elemento como *Car
}

// Last devuelve el último auto de la cola sin eliminarlo
func (cq *CarQueue) Last() *Car {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	if cq.queue.Len() == 0 {
		return nil // Si la cola está vacía, retorna nil
	}
	element := cq.queue.Back()  // Obtiene el último elemento de la cola
	return element.Value.(*Car) // Retorna el valor del elemento como *Car
}

// Size devuelve el número de autos en la cola
func (cq *CarQueue) Size() int {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	return cq.queue.Len() // Retorna el tamaño de la cola
}
