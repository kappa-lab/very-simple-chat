package main

type Room interface {
	Id() int
	Broadcast(message string)
	Unicast(reciever int, message string)
	AddListener(id int, listener chan string)
	RemoveListener(id int)
	CreateConnectionCtrl() ConnectionCtrl
	GetIds() []int
}

func NewRoom(id int) Room {
	return &room{
		id:        id,
		counter:   0,
		listeners: make(map[int]chan string, 0),
	}
}

type room struct {
	id        int
	counter   int
	listeners map[int]chan string
}

func (r *room) Id() int { return r.id }

func (r *room) Broadcast(message string) {
	//TODO mutex
	for _, v := range r.listeners {
		v <- message
	}
}

func (r *room) Unicast(reciever int, message string) {
	//TODO mutex
	ch := r.listeners[reciever]
	if ch != nil {
		ch <- message
	}
}

func (r *room) AddListener(id int, listener chan string) {
	//TODO mutex
	r.listeners[id] = listener
}

func (r *room) RemoveListener(id int) {
	//TODO mutex
	delete(r.listeners, id)
}

func (r *room) generateId() int {
	defer func() { r.counter += 1 }()
	return r.counter
}

func (r *room) CreateConnectionCtrl() ConnectionCtrl {
	return NewConnectionCtrl(r.generateId())
}

func (r *room) GetIds() []int {
	result := []int{}
	//TODO mutex
	for k := range r.listeners {
		result = append(result, k)
	}
	return result
}
