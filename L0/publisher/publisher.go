package main
import stan "github.com/nats-io/stan.go"
import "fmt"
import "os"
import "time"
import "io/ioutil"
import "L0/random"

func main() {
	clusterID := "test-cluster"
	clientID := "Person1"
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		fmt.Print(err)
		return
	}
	file, err := os.Open("../model.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()
	info, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	information := []byte{}
	for {
		time.After(40 * time.Second)
		information = random.Random_variable(info)
		err = sc.Publish("foo", information)
		if err == nil {
			fmt.Printf("%s PUBLISHER IS SENDING THESE FIELDS:\n", string(information))
		} else {
			fmt.Print(err)
		}
	}
}
