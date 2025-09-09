package main

import (
	"fmt"
)

// ожидаемый клиентом
type Entity interface {
	Operation()
}

// структура нового интерфейса
type ClientEntity struct{}

func (clientEntity *ClientEntity) Operation() {
	fmt.Println("print from ClientEntity")
}

// старая (не ожидаемая клиентом) сущность
type OldEntity struct{}

func (oldEntity *OldEntity) OldOperation() {
	fmt.Println("print from OldEntity")
}

// адаптер
type OldEntityAdapter struct {
	*OldEntity
}

// делегирование вызова функции
func (adapter *OldEntityAdapter) Operation() {
	adapter.OldOperation()
}

// создание нового интерфейса по старой структуре
func NewEntity(oldEntity *OldEntity) Entity {
	return &OldEntityAdapter{oldEntity}
}

func main() {
	oldEntity := &OldEntity{}

	entity := NewEntity(oldEntity)

	entity.Operation()

	entity = &ClientEntity{}

	entity.Operation()
}
