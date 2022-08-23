package main
import stan "github.com/nats-io/stan.go"
import "fmt"
import "context"
import "encoding/json"
import "time"
import "log"
import "html/template"
import "net/http"
import "L0/mainstruct"
import "L0/cache_db"

func main() {
	clusterID := "test-cluster"
	clientID := "Person2"
	database, err := cache_db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	Cache := cache_db.CreateCache()
	err = database.UploadCache(&Cache)
	if err != nil {
		log.Fatal(err)
	}
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	data := *new(mainstruct.MainStruct)
	if err != nil {
		fmt.Print(err); 
		return
	}
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, &data)
		if err != nil {
			log.Print(err)
		} else {
			err := Cache.Insert(data)
			fmt.Print(Cache.Get(data.Order_uid))
			if err != nil {
				log.Print(err)
			}
		}
	}, stan.StartWithLastReceived())
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			if val, err := Cache.Get(req.PostFormValue("order_uid")); err == nil {
				b, err := json.MarshalIndent(val, "", "\t")
				if err != nil {
					log.Println(err)
				} else {
					log.Printf("SENDING INFO WITH ORDER_ID: %s\n", req.PostFormValue("order_uid"))
					fmt.Fprint(w, string(b))
				}
			} else {
				fmt.Fprint(w, "Pleace enter correct order_uid!", err)
			}
		case "GET":
			tmpl, err := template.ParseFiles("../filename.html")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	})
	server := &http.Server{Addr: ("127.0.0.1:8080")}
	server.ListenAndServe()
	<-time.After(60 * time.Second)
	sub.Unsubscribe()
	sc.Close()
	server.Shutdown(context.Background())
	database.InsertInDB(Cache.GetAllCache()...)
}
