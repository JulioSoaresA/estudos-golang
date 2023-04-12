package arraylist

import "errors"

type ArrayList struct {
	values []int
	size   int
}

func (l *ArrayList) Init(size int) error {
	if size > 0 {
		l.values = make([]int, size)
		return nil
	} else {
		return errors.New("Size must be greater than 0")
	}
}

func (l *ArrayList) DoubleSize() {
	doubleValues := make([]int, len(l.values)*2)
	copy(doubleValues, l.values)
	l.values = doubleValues
}

func (l *ArrayList) Add(value int) {
	if l.size == len(l.values) {
		l.DoubleSize()
	}
	l.values[l.size] = value
	l.size++
}

func (l *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index < l.size {
		if l.size == len(l.values) {
			l.DoubleSize()
		}
		for i := l.size; i > index; i-- {
			l.values[i] = l.values[i-1]
		}
		l.values[index] = value
		l.size++
		return nil
	} else {
		if index < 0 {
			return errors.New("index must be greater than 0")
		} else {
			return errors.New("index must be less than size")
		}
	}
}

func (l *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < l.size {
		for i := index; i < l.size-1; i++ {
			l.values[i] = l.values[i+1]
		}
		l.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("index must be greater than 0")
		} else {
			return errors.New("index must be less than size")
		}
	}
}

func (l *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < l.size {
		return l.values[index], nil
	} else {
		if index < 0 {
			return 0, errors.New("index must be greater than 0")
		} else {
			return 0, errors.New("index must be less than size")
		}
	}
}

func (l *ArrayList) Set(value int, index int) error {
	if index >= 0 && index < l.size {
		l.values[index] = value
		return nil
	} else {
		if index < 0 {
			return errors.New("index must be greater than 0")
		} else {
			return errors.New("index must be less than size")
		}
	}
}

func (l *ArrayList) Size() int {
	return l.size
}
