package main

type Config struct {
	Day            int
	DateList       []string
	Bucket         string
	ACCESS_ID      string
	ACCESS_SEC_KEY string
	Endpoint       string
	PrefixPath     string
	Taskslist      *TaskSlots
	WORKER_NUM     int
	Forever        bool
}

func NewConfig() *Config {
	taskslist, _ := NewTaskSlots(10)
	return &Config{
		Day:            3,
		DateList:       GenerateDateList(10),
		Bucket:         "",
		ACCESS_ID:      "",
		ACCESS_SEC_KEY: "",
		Endpoint:       "",
		PrefixPath:     "",
		Taskslist:      taskslist,
		WORKER_NUM:     4,
		Forever:        true,
	}
}
