package list

import (
	"fmt"
)

type List[T any] struct {
	head *Node[T]
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

// Add добавляет новый элемент в список
func (l *List[T]) Add(value T) {
	newNode := &Node[T]{value: value}

	// если в списке еще нет ни одного нода
	if l.head == nil {
		l.head = newNode
		return
	}

	node := l.head
	for node.next != nil {
		node = node.next
	}

	node.next = newNode
}

// Count возвращает кол-во элементов в списке
func (l *List[T]) Count() int {
	count := 0
	l.Each(func(_ int, _ T) {
		count++
	})

	return count
}

// Each обойти каждый элемент в списке
func (l *List[T]) Each(callback func(index int, item T)) {
	node := l.head
	index := 0
	for node != nil {
		callback(index, node.value)
		node = node.next
		index++
	}
}

// Filter фильтрует список по переданному callback выражению
func (l *List[T]) Filter(callback func(index int, item T) bool) List[T] {
	newList := List[T]{}
	l.Each(func(index int, item T) {
		if callback(index, item) {
			newList.Add(item)
		}
	})

	return newList
}

// Find ищет и возвращает первый элемент, соответствующий выражению в callback
func (l *List[T]) Find(callback func(item T) bool) (*T, int) {
	node := l.head
	index := 0
	for node != nil {
		if callback(node.value) {
			return &node.value, index
		}
		node = node.next
		index++
	}

	return nil, -1
}

// Map изменить каждый элемент в списке
func (l *List[T]) Map(callback func(index int, item T) T) {
	node := l.head
	index := 0
	for node != nil {
		node.value = callback(index, node.value)
		node = node.next
		index++
	}
}

// Print выводит все элементы в консоль
func (l *List[T]) Print() {
	l.Each(func(index int, item T) {
		fmt.Printf("%d. %v\n", index+1, item)
	})
}

// Remove ищет и удаляет первый элемент, соответствующий выражению в callback
func (l *List[T]) Remove(callback func(item T) bool) {
	node := l.head
	var prev *Node[T]

	for node != nil {
		if callback(node.value) {
			if prev == nil {
				l.head = node.next
				return
			}

			prev.next = node.next
			return
		}
		prev = node
		node = node.next
	}
}

// ToSlice возвращает все элементы из списка в виде слайса
func (l *List[T]) ToSlice() []T {
	result := make([]T, 0)
	l.Each(func(index int, item T) {
		result = append(result, item)
	})

	return result
}
