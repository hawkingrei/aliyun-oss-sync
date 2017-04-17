package main

import (
	"testing"
)

func generateTestTask() Task {
	var tasktmp []string
	tasktmp = append(tasktmp, "abb")
	return Task{Keys: tasktmp}
}

func TestTasks(t *testing.T) {
	tasklist, _ := NewTasks()
	tmmp := generateTestTask()
	tasklist.Add(tmmp)
	tasklist.Add(tmmp)
	tasklist.Add(tmmp)
	tasklist.Add(tmmp)
	task, err := tasklist.Remove()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(task.Keys)
	if tasklist.Size() == 3 {
		t.Log("Success")
		return
	}
	t.Error("fail")
}
