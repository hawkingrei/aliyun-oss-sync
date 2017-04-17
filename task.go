package main

import (
	"container/list"
	"errors"
	"github.com/mohae/deepcopy"
	"sync"
)

type Task struct {
	Keys []string
}

type Tasks struct {
	size     int
	tasklist *list.List
	lock     sync.Mutex
}

func NewTasks() (*Tasks, error) {
	return &Tasks{
		size:     0,
		tasklist: list.New(),
	}, nil
}

func (t *Tasks) Add(task Task) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.size += 1
	t.tasklist.PushFront(&task)
	return nil
}

func (t *Tasks) Remove() (*Task, error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.size == 0 {
		return &Task{}, errors.New("tasklist empty")
	}
	result := t.tasklist.Back()
	returnResult := deepcopy.Copy(result.Value).(*Task)
	t.tasklist.Remove(result)
	t.size -= 1
	return returnResult, nil
}

func (t *Tasks) Size() int {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.size
}
