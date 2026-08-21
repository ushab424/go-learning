package main

import "fmt"

type Server struct {
	Name   string
	IP     string
	Status string
}

func main() {
	QuantityServerStatus := make(map[string]int)

	AllServers := []Server{
		{Name: "backup-nas-01.company.local", IP: "198.51.100.42", Status: "Online"},
		{Name: "corp-mail-01.company.local", IP: "203.0.113.185", Status: "Offline"},
		{Name: "prd-web-01.company.local ", IP: "192.168.1.45", Status: "Online"},
		{Name: "stg-db-02.company.local", IP: "10.0.4.112", Status: "Maintenance"},
		{Name: "dev-api-01.company.local", IP: "172.16.8.23", Status: "Offline"},
	}
	for _, ServerName := range AllServers {
		QuantityServerStatus[ServerName.Status]++
	}
	fmt.Println(QuantityServerStatus)
	for _, ServerStatus := range AllServers {
		if ServerStatus.Status == "Offline" {
			fmt.Println(ServerStatus.Name)
		}
	}
}
