package main

import (
	"github.com/spaolacci/murmur3"
)

func hashFunc(data []byte) uint64 {
	return murmur3.Sum64(data)
}

type TaskSlots struct {
	num   int
	slots []*Tasks
}

func NewTaskSlots(WORKER_NUM int) (*TaskSlots, error) {
	t := new(TaskSlots)
	t.num = WORKER_NUM
	for i := 0; i < WORKER_NUM; i++ {
		tmp, _ := NewTasks()
		t.slots = append(t.slots, tmp)
	}
	return t, nil
}

func (t TaskSlots) Add(path string, task Task) (error, uint64) {
	hashVal := hashFunc([]byte(path))
	slotId := hashVal & (uint64(t.num) - 1)
	return t.slots[slotId].Add(task), slotId
}

func (t TaskSlots) Remove(slotId uint64) (*Task, error) {
	return t.slots[slotId&(uint64(t.num)-1)].Remove()
}

func (t TaskSlots) Size(slotId uint64) int {
	return t.slots[slotId&(uint64(t.num)-1)].Size()
}
