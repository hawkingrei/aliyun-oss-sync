package main

import (
	"testing"
)

func TestTaskSlot(t *testing.T) {
	tasklist, _ := NewTaskSlots(4)
	tmmp := generateTestTask()
	tasklist.Add("1", tmmp)
	tasklist.Add("1", tmmp)
	tasklist.Add("1", tmmp)
	tasklist.Add("1", tmmp)
	//task, err := tasklist.Remove(1)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(task.Keys)
	//if tasklist.Size() == 3 {
	//	t.Log("Success")
	//	return
	//}
	//t.Error("fail")
}
