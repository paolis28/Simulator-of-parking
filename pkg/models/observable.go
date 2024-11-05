package models

// Observable define la interfaz para objetos observables en el patrón Observer.
type Observable interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}
