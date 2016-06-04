package util

import "testing"

func TestNew(t *testing.T) {
	_, err := New(-1)
	if err == nil {
		t.Error("Not panic when nagetive initial capacity")
	}
	if err.Error() != "Illegal Capacity:-1" {
		t.Error("Excepte", "Illegal Capacity:-1", "but was", err.Error())
	}
}

func TestAdd(t *testing.T) {
	arrayList, err := New(1)
	if err != nil {
		t.Error("Error create array list")
	}
	firstE := "first"
	arrayList.Add(firstE)
	if arrayList.Size() != 1 {
		t.Error("Fail to add element to array list")
	}

	secondE := "second"
	arrayList.Add(secondE)
	if arrayList.Size() != 2 {
		t.Error("fail to add second element to array list")
	}

	thirdE := "third"
	arrayList.Add(thirdE)
	if arrayList.Size() != 3 {
		t.Error("fail to add third element to array list")
	}
}
