package main
import(
       "fmt"
       "net/http"
       )
func main(){
       http.HandleFunc("/", handler)
       http.ListenAndServe("0.0.0.0:8080", nil)
       
}

func handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Hello World!")
}
