package util

type List interface {
	Size()
	IsEmpty()
	Contains()
	ToArray()
	Add(interface{})
	Remove(interface{})
	ContainsAll(*List)
	AddAll(*List)
	RemoveAll(*List)
	RetainAll(*List)
	Clear()
	Equals()
	Get(int)
	Set(int, interface{})
	IndexOf(interface{})
	LastIndexOf(interface{})
	SubList(int, int)
}
