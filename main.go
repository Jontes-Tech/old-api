package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func archLinux(w http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&w, req)
	rackspaceurl := getLatestArchLinux("rackspace")
	fmt.Fprint(w, rackspaceurl)
	timeOfRequest := time.Now()
	fmt.Println("Incomming Request: ",req.Method, req.URL.Path, timeOfRequest.Format("01-02-2006 15:04:05"))
}
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
 }
func handleRequests() {
	http.HandleFunc("/api/arch-rackspace-latest", archLinux)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func getLatestArchLinux(mirror string) string {
	fmt.Println(mirror)
	resp, err := http.Get("https://mirror.rackspace.com/archlinux/iso/latest/arch/version")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	outurl := "https://mirror.rackspace.com/archlinux/iso/latest/archlinux-" + sb + "-x86_64.iso"
	return outurl
}
func main() {
	fmt.Println("Jonte's Arch API Server - Written in Go.")
	fmt.Println("Rackspace Mirror accessible at: http://localhost:8080/api/arch-rackspace-latest")
	fmt.Println("Demo Accessible Here: https://jontes.page/api/arch")
	fmt.Println("Starting the application...")
	fmt.Println("Running Server on Port 8080")
	handleRequests()
}