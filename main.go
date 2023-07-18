package main
import(
       "fmt"
       "net/http"
       )
func main(){
       http.HandleFunc("/", handler)
       http.ListenAndServe("127.0.0.1:8080", nil)
       
}

func handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Hello World!")
}
