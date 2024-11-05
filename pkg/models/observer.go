package models

// Observer define la interfaz para los observadores en el patrón Observer.
type Observer interface {
	Update(data interface{})
}
