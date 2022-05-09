// Notes to Opensourcers:
// - The code is really messy, but it works.
// - The code is not optimized, but it works.
package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)
// Gets URL parameter, then returns the latest Arch Linux release URL usiong the getLatestArchLinux function.
func archLinux(w http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&w, req)
	mirrors, ok := req.URL.Query()["mirror"]
    if !ok || len(mirrors[0]) < 1 {
        fmt.Fprint(w, "URL parameter 'mirror' is missing")
        return
    }
    mirror := mirrors[0]
	rackspaceurl := getLatestArchLinux(string(mirror))
	fmt.Fprint(w, strings.ReplaceAll(rackspaceurl, "\n", ""))
}
func handleRequests() {
	http.HandleFunc("/api/arch", archLinux)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
}
// Main Function for getting the latest Arch Linux release URL. 
func getLatestArchLinux(mirror string) string {
	resp, err := http.Get("https://mirror.rackspace.com/archlinux/iso/latest/arch/version")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	outurl := "null"
	if mirror == "rackspace" {
		outurl := "https://mirror.rackspace.com/archlinux/iso/latest/archlinux-" + sb + "-x86_64.iso"
		return outurl
	} else if mirror == "acc-umu" {
		outurl := "https://ftp.acc.umu.se/mirror/archlinux/iso/" + sb + "/archlinux-" + sb + "-x86_64.iso"
		return outurl
	}
	fmt.Println("Mirror " + mirror + " requested.")
	return outurl
}
// Main Function - Used for printing friendly messages to the server console.
func main() {
	colorReset := "\033[0m"
    colorBlue := "\033[34m"
	fmt.Println(string(colorBlue),"Jonte's Arch Linux Mirror API")
	fmt.Println(string(colorReset),"Rackspace Mirror accessible at: http://localhost:8080/api/arch")
	fmt.Println("Demo Accessible Here: https://jontes.page/api/arch")
	fmt.Println("Starting the application...")
	fmt.Println("Running Server on Port 8080 (Or whatever you mapped your Docker container to)")
	handleRequests()
}