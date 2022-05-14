package main
import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"os"
)
var oldid = "error"
var serverresponse = "error"
func main(){
	fmt.Println("Welcome to VideoLearner! \n------ \nEasily follow YouTube Tutorials. \nLicensed under GPL 3.0 - Free as in Freedom. \nProgram assumes internet access and access to vlserver.jontes.page.")
	handleRequests()
}
func getRequestInfo(w http.ResponseWriter, req *http.Request) {
	arguments := req.URL.Query()
	vidid := arguments.Get("vidid")
	time := arguments.Get("t")
	token := os.Getenv("VLTOKEN")
	fmt.Println("Someone requested " + vidid + " at " + time)
	if vidid != oldid {
		fmt.Println("VIDID CHANGED from " + oldid + " to " + vidid)
		oldid = vidid
		fmt.Println(oldid)
		serverresponse = getVideoInfo(vidid, token)
	 }
	fmt.Println(serverresponse)
}
func handleRequests() {
	http.HandleFunc("/video", getRequestInfo)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func getVideoInfo(vidid string, token string)string {
	resp, err := http.Get("https://vlserver.jontes.page/video?videoid="+vidid+"&token="+token)
	if err != nil {
	   log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
	sb := string(body)
	return sb
}