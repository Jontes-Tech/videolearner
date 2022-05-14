package main
import (
	"fmt"
	"net/http"
	"log"
)
func main(){
	fmt.Println("VideoLearner!")
	handleRequests()
}
func getRequestInfo(w http.ResponseWriter, req *http.Request) {
	arguments := req.URL.Query()
	vidid := arguments.Get("vidid")
	time := arguments.Get("t")
	fmt.Println("Someone requested " + vidid + " at " + time)
	getVideoInfo(vidid, time)
}
func handleRequests() {
	http.HandleFunc("/video", getRequestInfo)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func getVideoInfo(vidid string, time string)string {
	resp, err := http.Get("https://vlserver.jontes.page/video?videoid="+vidid+"&token=rickastley")
	if err != nil {
	   log.Fatalln(err)
	}
}