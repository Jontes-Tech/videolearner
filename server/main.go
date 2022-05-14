package main
import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)
func main(){
	fmt.Println("VideoLearner - Server - Port 3000.")
	handleRequests()
}
func getVideoInfo(w http.ResponseWriter, req *http.Request) {
	arguments := req.URL.Query()
	videoid := arguments.Get("videoid")
	token := arguments.Get("token")
	tokenvalid := checktokenvalidiy(token)
	fmt.Fprint(w, tokenvalid)
	if tokenvalid == "valid"{
		fmt.Fprint(w, "good")
	}
}
func checktokenvalidiy(token string)string {
	b, err := ioutil.ReadFile("acceptedtokens.txt")
    if err != nil {
      fmt.Print(err)
    }
    str := string(b)
	result := "nan"
    if !strings.Contains(str, token) {
		result := "auth token not valid"
		return result
	}else if strings.Contains("#", token) {
		result := "comments not counted or token empty"
		return result
	}else {
		result := "valid"
		return result
	}
	return(result)
}
func handleRequests() {
	http.HandleFunc("/video", getVideoInfo)
	log.Fatal(http.ListenAndServe(":3000", nil))
}