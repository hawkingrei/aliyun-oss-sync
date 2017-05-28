package main

type NSQD struct {
	PreChan   chan string
	TaskChan  chan string
	ExitChan  chan int
	WaitGroup WaitGroupWrapper
}

func New() *NSQD {
	return &NSQD{
		PreChan:  make(chan string, 40960),
		TaskChan: make(chan string, 40960000),
		ExitChan: make(chan int),
	}
}

func main() {
	config := NewConfig()
	client, _ := NewClient(config)
	nsq := New()

	nsq.WaitGroup.Wrap(func() { Preproducor(config.Day, nsq, config) })
	for i := 0; i < config.Producer; i++ {
		nsq.WaitGroup.Wrap(func() { client.GenerateTask(nsq, config) })
	}
	for i := 0; i < config.Producer; i++ {
		nsq.WaitGroup.Wrap(func() { client.Worker(nsq) })
	}
	nsq.WaitGroup.Wait()
}
