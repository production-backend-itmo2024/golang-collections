//go:build !solution

package list

import "fmt"

type IList[T any] interface {
	Size() int64
	IsEmpty() bool

	PushBack(elem T) error
	PopBack(elem T) error
	Insert(elem T, pos int64) error
	Erase(pos int64) (T, error)
}

var _ IList[string] = &List[string]{}

type List[T any] struct {
}

func NewList[T any]() *List[T] {
	return nil
}

func (l *List[T]) PushBack(elem T) error {
	return fmt.Errorf("implement me")
}

func (l *List[T]) PopBack(elem T) error {
	return fmt.Errorf("implement me")
}

func (l *List[T]) Insert(elem T, pos int64) error {
	return fmt.Errorf("implement me")
}

func (l *List[T]) Erase(pos int64) (T, error) {
	return *new(T), fmt.Errorf("implement me")
}

func (l *List[T]) IsEmpty() bool {
	return false
}

func (l *List[T]) Size() int64 {
	return 0
}
