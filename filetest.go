package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	//filename := "./test.txt"
	filename := os.Args[1]
	fmt.Println(filename)
	f, err := os.OpenFile(filename, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		f, err = os.OpenFile(filename, os.O_CREATE | os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("File open error")
			panic(err)
		}
	}

	w := bufio.NewWriter(f)
	content := "\nhello world"
	n, err := w.WriteString(content)
	if err != nil {
		fmt.Println("Write file error!")
		panic(err)
	}


	//var n int
	//n, err = f.WriteString(content)
	//if err != nil {
	//	panic(err)
	//}
	//err = f.Sync()
	//if err != nil {
	//	panic(err)
	//}

	fmt.Fprintln(f, content)

	fmt.Println(n)
	fmt.Println(w.Buffered())
	err = w.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println(w.Buffered())
	//fmt.Println(w.Available())
	fmt.Println(n)
	f.Close()
}
