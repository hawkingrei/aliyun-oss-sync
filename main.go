package main

import (
	"fmt"
	"github.com/op/go-logging"
	"sync"
)

var log = logging.MustGetLogger("aliyun-oss-sync")

var quit_worker = make(chan bool)
var worker_channel = make(chan Task, 4)
var Productionor_channel = make(chan bool, 1)

func Productionor(config *Config, client Client) {
	var err error
	var wg sync.WaitGroup
	for {
		config.DateList = GenerateDateList(config.Day)
		for _, datepath := range config.DateList {
			wg.Add(1)
			go func(datepath string) {
				var task Task
				path := config.PrefixPath + datepath
				log.Info(path)
				task, err = client.GenerateTask(path, config)
				if err != nil {
					log.Info(err.Error())
					return
				}
				worker_channel <- task
				wg.Done()
			}(datepath)
		}
		config.DateList = GenerateDateList(config.Day)
		wg.Wait()
	}

}

func worker(config *Config, client Client) {
	for {
		select {
		case t := <-worker_channel:
			for _, key := range t.Keys {
				log.Info(key)
				err := client.ChangeContentType(string(key))
				if err != nil {
					fmt.Println("error " + string(key) + " " + err.Error())
				}
			}
		}
	}

}

func main() {
	config := NewConfig()
	//var SLOTID_CHANNEL = make(chan uint64, config.WORKER_NUM)
	client, err := NewClient(config)
	if err != nil {
		fmt.Println(err.Error())
		panic("init client error")
	}

	go Productionor(config, client)

	for i := 0; i < config.WORKER_NUM; i++ {
		go worker(config, client)
	}
	<-quit_worker

	//	}
	//	//

	//	sigCh := make(chan os.Signal, 1)
	//	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	//	<-sigCh

	/*
		for i := 0; i < WORKER_NUM; i++ {

			go func() {
				for {
					t, err := config.Taskslist.Remove()
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
					TASK_CHANNEL <- t
				}
			}()
			go func() {
				for {
					var t *Task
					t = <-TASK_CHANNEL
					for _, key := range t.Keys {
						fmt.Println(key)
						err = client.ChangeContentType(key)
						if err != nil {
							fmt.Println("error " + key)
						}
					}
				}
			}()
		}
	*/
	//client.Bucket.ListObjects(oss.Prefix("uploads/files"))
	//lsRes, err := client.Bucket.ListObjects(oss.Prefix("uploads/files/xxx"))
	//for _, object := range lsRes.Objects {
	//	fmt.Println("Objects:", object.Key)
	//}
	//props, _ := bucket.GetObjectDetailedMeta("uploads/files/201401/20/250px-Chitanda.Eru.full.1053233.jpg")
	//fmt.Println("Object Meta:", props)
	// := oss.ContentType("image/jpeg")

	//filename := "uploads/files/201401/20/250px-Chitanda.Eru.full.1053233.jpg"
	//option := checkmap[ReturnFilenameExtension(filename)]
	//_ = bucket.SetObjectMeta("uploads/files/201401/20/250px-Chitanda.Eru.full.1053233.jpg", option)

}
