package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func archLinux(w http.ResponseWriter, r *http.Request) {
	rackspaceurl := getLatestArchLinux("rackspace")
	fmt.Fprint(w, rackspaceurl)
	fmt.Println("Endpoint Hit: Arch Linux Rackspace")
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
