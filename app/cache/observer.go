package cache

// Observer es una función que toma una URL como argumento
type Observer func(url string)

// Subject mantiene una lista de observadores
type Subject struct {
	observers []Observer
}

// NewSubject crea un nuevo sujeto
func NewSubject() *Subject {
	return &Subject{}
}

// AddObserver agrega un observador al sujeto
func (s *Subject) AddObserver(o Observer) {
	s.observers = append(s.observers, o)
}

// NotifyObservers notifica a todos los observadores
func (s *Subject) NotifyObservers(url string) {
	for _, o := range s.observers {
		o(url)
	}
}
