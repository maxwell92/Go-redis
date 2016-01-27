package main

import (
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"time"
	"math/rand"
)

type alog struct {
	Time string `json: "Time"`
	Event string `json: "Event"`
	Error string `json: "Error"`
	Description string `json: "Description"`
}

type blog struct {
	Time string `json: "Time"`
	Event string `json: "Event"`
	Error string `json: "Error"`
	Description string `json: "Description"`
}


type clog struct {
	Time string `json: "Time"`
	Event string `json: "Event"`
	Error string `json: "Error"`
	Description string `json: "Description"`
}


func validate(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		_, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}
	f.Close()
}


func main()  {
	timeLayout := "2006-01-02 15:04:05"

	afilename := os.Args[1]
	validate(afilename)
	fa, err := os.OpenFile(afilename, os.O_RDWR | os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	awriter := bufio.NewWriter(fa)


	bfilename := os.Args[2]
	validate(bfilename)
	fb, err := os.OpenFile(bfilename, os.O_RDWR | os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	bwriter := bufio.NewWriter(fb)

	cfilename := os.Args[3]
	validate(cfilename)
	fc, err := os.OpenFile(cfilename, os.O_RDWR | os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	cwriter := bufio.NewWriter(fc)

		go func() {
			loga := new(alog)
			for i := 0; i < 10; i++ {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				trand := r.Intn(10)
				time.Sleep(time.Second * time.Duration(trand))
				loga.Time = time.Now().Format(timeLayout)
				loga.Error = "Fatal Error"
				loga.Event = "Segment Fault"
				loga.Description = "This is a test of alog"
				//content := loga.Time + loga.Event + loga.Error + loga.Description
				ajson, _ := json.Marshal(loga)
				awriter.WriteString(string(ajson))
				awriter.Flush()
				fmt.Println("logA:"+string(ajson))
			}
		}()
		go func() {
			logb := new(blog)
			for i := 0; i < 10; i++ {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				trand := r.Intn(10)
				time.Sleep(time.Second * time.Duration(trand))
				logb.Time = time.Now().Format(timeLayout)
				logb.Error = "Warning"
				logb.Event = "Stack Overflow"
				logb.Description = "This is a test of blog"
				//content := logb.Time + logb.Event + logb.Error + logb.Description
				bjson, _ := json.Marshal(logb)
				bwriter.WriteString(string(bjson))
				bwriter.Flush()
				fmt.Println("logB:"+string(bjson))
			}
		}()
		go func() {
			logc := new(clog)
			for i := 0; i < 10; i++ {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				trand := r.Intn(10)
				time.Sleep(time.Second * time.Duration(trand))
				logc.Time = time.Now().Format(timeLayout)
				logc.Error = "Unknow Issue"
				logc.Event = "Out of Memory"
				logc.Description = "This is a test of clog"
				//content := logc.Time + logc.Event + logc.Error + logc.Description
				cjson, _ := json.Marshal(logc)
				cwriter.WriteString(string(cjson))
				cwriter.Flush()
				fmt.Println("logC:"+string(cjson))
			}
		}()
	fmt.Println("waiting...")
	time.Sleep(time.Second * time.Duration(120))

}



