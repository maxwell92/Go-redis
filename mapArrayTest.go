package main
import "fmt"

func main() {
	image := make(map[string][]string)
	tag := make([]string, 5, 20)
	tag[0] = "v1.0.0"
	tag[1] = "latest"
	image["zookeeper"] = tag

	fmt.Println(image["zookeeper"])


}
