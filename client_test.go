package main

import (
	"testing"
)

func TestGenerateTasks(t *testing.T) {
	config := NewConfig()
	client, err := NewClient(config)
	if err != nil {
		t.Error(err)
	}
	task, _ := client.GenerateTask("uploads/files/201401/20", config)
	t.Log(task.Keys)
}
