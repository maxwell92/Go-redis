package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"math/rand"
)

// Log Gen:
// logPathA: [Time] [ErrorType] [Event] [Status]
// e.g.: [2016.01.18 11:53] [Fatal|Error|Warning] [Segment Fault| Page Fault | Stack Overflow] [Stopped|Paused|Cleaned]
//
// logPathB: [Time] [LogNum] [Event]
// e.g.: [2016.01.18 11:53] [1222] [Segment Fault]

// logPathC: [Time] [IP][Node] [Status]
// e.g.: [2016.01.18 11:53] [172.21.1.21] [node1] [Down | Up | Unknow]




type A struct {
	time string
	errorType map[int]string
	event map[int]string
	status map[int]string
}

type B struct {
	time string
	logNum int
	event map[int]string

}

type C struct {
	time string
	ip string
	node string
	status map[int]string

}

type LogGen struct {

	fA *os.File
	fB *os.File
	fC *os.File

	writerA *bufio.Writer
	writerB *bufio.Writer
	writerC *bufio.Writer

	genIntervalA int //generate interval
	genIntervalB int
	genIntervalC int

}


//func initialize(filename[] string) (lg *LogGen) {
//
//	return 	&LogGen {
//		fA : os.OpenFile(filename[0], os.O_RDWR | os.O_APPEND, 0666),
//		fB : os.OpenFile(filename[1], os.O_RDWR | os.O_APPEND, 0666),
//		fC : os.OpenFile(filename[2], os.O_RDWR | os.O_APPEND, 0666),
//
//		//writerA : bufio.NewWriter(fA),
//		//writerB : bufio.NewWriter(fB),
//		//writerC : bufio.NewWriter(fC),
//		genIntervalA: 3,
//		genIntervalB: 5,
//		genIntervalC: 8,
//	}
//}

func initialize(filename[] string, lg *LogGen) (err error) {
	lg.fA,_ = os.OpenFile(filename[0], os.O_RDWR | os.O_APPEND, 0666)
	lg.fB,_ = os.OpenFile(filename[1], os.O_RDWR | os.O_APPEND, 0666)
	lg.fC,_ = os.OpenFile(filename[2], os.O_RDWR | os.O_APPEND, 0666)
	lg.writerA = bufio.NewWriter(lg.fA)
	lg.writerB = bufio.NewWriter(lg.fB)
	lg.writerC = bufio.NewWriter(lg.fC)
	lg.genIntervalA = 3
	lg.genIntervalB = 5
	lg.genIntervalC = 8
	return nil
}

func validateFilename(filename[] string) (err error) {
	for i := 0; i < 3; i++ {
		_,err := os.Open(filename[i])
		if err != nil {
			_, err = os.Create(filename[i])
			if err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func main()  {
	timeLayout := "2006-01-02 15:04:05"
	filename := make([]string, 3)
	for i := 0; i < 3; i++ {
		filename[i] = os.Args[i+1]
	}

	validateFilename(filename)
	var lg LogGen
	initialize(filename, &lg)
	//fmt.Println(lg.genIntervalA)
	//fmt.Println(lg.genIntervalB)
	//fmt.Println(lg.genIntervalC)
	//
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	//var logA A
	//var logB B
	//var logC C

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	trand := r.Intn(10)
	lg.genIntervalA = trand

	timerA := time.NewTimer(time.Second * time.Duration(lg.genIntervalA))
	//timerA := time.NewTimer(time.Second * 3)
	for i := 0; ; i++ {
		<-timerA.C
		fmt.Println(time.Now().Format(timeLayout))
		timerA.Reset(time.Second * time.Duration(lg.genIntervalA))
		//timerA.Reset(time.Second * 3)
	}


}


