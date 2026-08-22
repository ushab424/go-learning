package main

import "fmt"

type ServerInfo struct {
	Name   string
	IP     string
	Status string
}

func (S ServerInfo) Start() ServerInfo {
	S.Status = "Online"
	return S
}

func (S ServerInfo) Stop() ServerInfo {
	S.Status = "Offline"
	return S
}

func (S ServerInfo) Info() {
	fmt.Println(S.Name, S.IP, S.Status)
}

func main() {
	ServerList := []ServerInfo{
		{Name: "db-prod-01", IP: "192.168.10.25", Status: "Offline"},
		{Name: "web-stage-core", IP: "10.0.1.50", Status: "Online"},
		{Name: "backup-vault-us", IP: "172.16.5.112", Status: "Offline"},
	}
	ServerList[0] = ServerList[0].Start()
	ServerList[1] = ServerList[0].Stop()
	for _, Info := range ServerList {
		Info.Info()
	}
}
