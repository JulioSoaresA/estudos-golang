package linkedlist

import (
	"banco/contas"
	"errors"
)

type LinkedList struct {
	head *Node
	size int
}

type Node struct {
	value contas.ContaCorrente
	next  *Node
}

func (l *LinkedList) Append(conta contas.ContaCorrente) {
	node := &Node{value: conta}
	aux := l.head
	prev := aux

	for aux != nil {
		prev = aux
		aux = aux.next
	}

	if prev == nil {
		l.head = node
	} else {
		prev.next = node
	}

	l.size++
}

func (l *LinkedList) AddOnIndex(conta contas.ContaCorrente, index int) error {
	if index >= 0 && index <= l.size {
		node := &Node{value: conta}

		if index == 0 {
			if l.head != nil {
				node.next = l.head
			}

			l.head = node
		} else {
			prev := l.head

			for i := 0; i < index-1; i++ {
				prev = prev.next
			}

			node.next = prev.next
			prev.next = node
		}
		l.size++
		return nil
	} else {
		if index < 0 {
			return errors.New("index out of bounds")
		} else {
			return errors.New("index out of bounds")
		}
	}
}

func (l *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < l.size {
		if index == 0 {
			l.head = l.head.next
		} else {
			del := l.head
			prev := l.head

			for i := 0; i < index-1; i++ {
				prev = prev.next
				del = del.next
			}

			prev.next = del.next
		}

		l.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("index out of bounds")
		} else {
			return errors.New("index out of bounds")
		}
	}
}

func (l *LinkedList) Get(index int) (contas.ContaCorrente, error) {
	if index >= 0 && index < l.size {
		aux := l.head

		for i := 0; i < index; i++ {
			aux = aux.next
		}

		return aux.value, nil
	} else {
		if index < 0 {
			return contas.ContaCorrente{}, errors.New("index out of bounds")
		} else {
			return contas.ContaCorrente{}, errors.New("index out of bounds")
		}
	}
}

func (l *LinkedList) Set(conta contas.ContaCorrente, index int) error {
	if index >= 0 && index < l.size {
		aux := l.head

		for i := 0; i < index; i++ {
			aux = aux.next
		}

		aux.value = conta
		return nil
	} else {
		if index < 0 {
			return errors.New("index out of bounds")
		} else {
			return errors.New("index out of bounds")
		}
	}
}

func (l *LinkedList) Size() int {
	return l.size
}
