package obser

import("parking/src/models/observer")

type Obser interface {
	RegisterObserver(o observer.Observer)
	RemoveObserver(o observer.Observer)
	NotifyObservers()
}
