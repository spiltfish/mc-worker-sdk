package mc_worker_sdk

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"log"
)

var ( WorkerAddress string )

func SetServer( workerAddress string){
	WorkerAddress = workerAddress
}

func GetMinecraftList()(resp *http.Response) {
	resp, err := http.Get( WorkerAddress + "/minecraft")
	if err != nil {
		log.Println("Error getting server list: ", err)
	}
	return resp
}

func GetMinecraftServerIp(name string)(ip []byte){
	response, err := http.Get( WorkerAddress + "/minecraft/" + name + "/ip")
	if err != nil {
		log.Println("Error getting server IP: ", err)
	}
	ip, _ = ioutil.ReadAll(response.Body)
	return ip
}

func GetMinecraftServerStatus(name string)(status []byte){
	response, err := http.Get( WorkerAddress + "/minecraft/" + name + "/status")
	if err != nil {
		log.Println("Error getting status: ", err)
	}
	status, _ = ioutil.ReadAll(response.Body)
	return status
}

func GetMinecraftServer(name string)(resp *http.Response){
	resp, err := http.Get( WorkerAddress + "/minecraft/" + name)
	if err != nil {
		log.Println("Error getting Server Info: ", err)
	}
	return resp
}

func CreateMinecraftServer(name string)(resp *http.Response){
	req, err := http.NewRequest("PUT", WorkerAddress + "/minecraft/" + name, nil)
	if err != nil {
		log.Println("Error Creating Server: ", err)
	}
	resp, _ = http.DefaultClient.Do(req)
	return resp
}

func PowerOnServer(name string)(resp *http.Response){
	req, err := http.NewRequest("POST", WorkerAddress + "/minecraft/" + name,
		bytes.NewBufferString("{ \"Power\" : \"off\" } "))
	if err != nil {
		log.Println("Error powering off server: ", err)
	}
        resp, _ = http.DefaultClient.Do(req)
	return resp
}

func PowerOffServer(name string)(resp *http.Response){
	req, err := http.NewRequest("POST", WorkerAddress + "/minecraft/" + name,
		bytes.NewBufferString("{ \"Power\" : \"on\" } "))
	if err != nil {
		log.Println("Error powering off server: ", err)
	}
	resp, _ = http.DefaultClient.Do(req)
	return resp
}

func DeleteMinecraftServer(name string)(resp *http.Response){
	req, err := http.NewRequest("DELETE", WorkerAddress + "/minecraft/" + name, nil)
	if err != nil {
		log.Println("Error deleting server: ", err)
	}
	resp, _ = http.DefaultClient.Do(req)
	return resp
}
