package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EndPoint struct {
	Path     string
	Function func(http.ResponseWriter, *http.Request)
}

var endpoints = []EndPoint{
	{
		Path:     "/buckets",
		Function: Buckets,
	},
	{
		Path:     "/contact",
		Function: Contact,
	},
}

func PrintHead(w *http.ResponseWriter) {
	fmt.Fprintf(*w, "DSpotify Kademlia Server\nID: %08b\n", kademliaServer.Contact.Id[:])
	fmt.Fprintf(*w, "IP: %s PORT: %d\n\n", kademliaServer.Contact.Ip, kademliaServer.Contact.Port)
}

func Buckets(w http.ResponseWriter, req *http.Request) {
	PrintHead(&w)
	for i, bucket := range kademliaServer.Buckets.Buckets {
		fmt.Fprintf(w, "BUCKET %3d:\n", i)
		for j, b := range bucket {
			fmt.Fprintf(w, "|-- %3d %08b  %s:%4d\n", j, b.Id[:], b.Ip.String(), b.Port)
		}
	}
}

func Contact(w http.ResponseWriter, req *http.Request) {
	b, _ := json.Marshal(kademliaServer.Contact)
	fmt.Fprintln(w, string(b))
}
