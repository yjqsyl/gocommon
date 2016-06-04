package util

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	elementData []interface{}
	size        int
	modCount    int
}

func New(initialCapacity int) (*ArrayList, error) {
	if initialCapacity < 0 {
		return nil, errors.New(fmt.Sprintf("Illegal Capacity:%d", initialCapacity))
	}
	var arrayList *ArrayList = new(ArrayList)
	arrayList.elementData = make([]interface{}, initialCapacity)
	arrayList.size = 0
	return arrayList, nil
}

func (arrayList *ArrayList) Size() int {
	return arrayList.size
}

func (arrayList *ArrayList) IsEmpty() bool {
	return arrayList.size == 0
}

func (arrayList *ArrayList) IndexOf(o interface{}) int {
	for i := 0; i < arrayList.size; i++ {
		if arrayList.elementData[i] == o {
			return i
		}
	}
	return -1
}

func (arrayList *ArrayList) LastIndexOf(e interface{}) int {
	for i := arrayList.size - 1; i >= 0; i++ {
		if arrayList.elementData[i] == 0 {
			return i
		}
	}
	return -1
}

func (arrayList *ArrayList) Contains(o interface{}) bool {
	return arrayList.IndexOf(o) > 0
}

func (arrayList *ArrayList) ToArray() []interface{} {
	array := make([]interface{}, arrayList.size)
	for i := 0; i < arrayList.size; i++ {
		array[i] = arrayList.elementData[i]
	}
	return array
}

func (arrayList *ArrayList) ensureCapacity(minCapacity int) {
	arrayList.modCount++
	oldCapacity := len(arrayList.elementData)
	if oldCapacity < minCapacity {
		newCapacity := oldCapacity*3/2 + 1
		if newCapacity < minCapacity {
			newCapacity = minCapacity
		}
		newArray := make([]interface{}, minCapacity-oldCapacity)
		arrayList.elementData = append(arrayList.elementData, newArray...)
	}
}

func (arrayList *ArrayList) Add(e interface{}) bool {
	arrayList.ensureCapacity(arrayList.size + 1)
	arrayList.elementData[arrayList.size] = e
	arrayList.size = arrayList.size + 1
	return true
}

func (arrayList *ArrayList) Remove(e interface{}) bool {
	index := arrayList.IndexOf(e)
	if index != -1 {
		arrayList.fastRemove(index)
		return true
	} else {
		return false
	}
}

func (arrayList *ArrayList) fastRemove(index int) {
	arrayList.modCount++
	arrayList.elementData = append(arrayList.elementData[0:index], arrayList.elementData[index+1:]...)
	arrayList.elementData[index] = nil
	arrayList.size = arrayList.size - 1
}

func (arrayList *ArrayList) Clear() {
	arrayList.modCount++
	for pos, _ := range arrayList.elementData {
		arrayList.elementData[pos] = nil
	}
	arrayList.size = 0
}

func (arrayList *ArrayList) Equals(arrayListSrc *ArrayList) bool {
	if arrayList == arrayListSrc {
		return true
	}
	if arrayList.size != arrayListSrc.size {
		return false
	}
	for i := 0; i < arrayList.size; i++ {
		if arrayList.elementData[i] != arrayListSrc.elementData[i] {
			return false
		}
	}
	return true
}

func (arrayList *ArrayList) Get(index int) (interface{}, error) {
	err := arrayList.rangeCheck(index)
	if err != nil {
		return nil, err
	}
	return arrayList.elementData[index], nil
}

func (arrayList *ArrayList) Set(index int, e interface{}) (interface{}, error) {
	err := arrayList.rangeCheck(index)
	if err != nil {
		return nil, err
	}
	oldE := arrayList.elementData[index]
	arrayList.elementData[index] = e
	return oldE, nil
}

func (arrayList *ArrayList) rangeCheck(index int) error {
	if index >= arrayList.size {
		return errors.New(fmt.Sprintln("Index: ", index, ", Size: ", arrayList.size))
	}
	return nil
}

func (arrayList *ArrayList) SubList(from int, to int) (*ArrayList, error) {
	size := arrayList.size
	if from < 0 {
		return nil, errors.New(fmt.Sprintln("fromIndex =", from))
	}
	if to >= size {
		return nil, errors.New(fmt.Sprintln("toIndex =", to))
	}
	if from > to {
		return nil, errors.New(fmt.Sprintln("fromIndex(", from, ") > toIndex(", to, ")"))
	}
	return &ArrayList{arrayList.elementData[from:to], arrayList.size, arrayList.modCount}, nil
}
