package mc_worker_sdk

import (
	"net/http"
	"bytes"
)

var ( WorkerAddress string )

func SetServer( workerAddress string){
	WorkerAddress = workerAddress
}

func GetMinecraftList()(resp *http.Response){
	resp, _ = http.Get( WorkerAddress + "/minecraft")
	return resp
}

func GetMinecraftServerIp(name string)(ip []byte){
	resp, _ := http.Get( WorkerAddress + "/minecraft/" + name + "/ip")
	ip, _ = ioutil.Readall(response.Body)
	return ip
}

func GetMinecraftServerStatus(name string)(status []byte){
	resp, _ = http.Get( WorkerAddress + "/minecraft/" + name + "/status")
	status, _ = ioutil.ReadAll(response.Body)
	return status
}

func GetMinecraftServer(name string)(resp *http.Response){
	resp, _ = http.Get( WorkerAddress + "/minecraft/" + name)
	return resp
}

func CreateMinecraftServer(name string)(resp *http.Response){
	req, _ := http.NewRequest("PUT", WorkerAddress + "/minecraft/" + name, nil)
	resp, _ = http.DefaultClient.Do(req)
	return resp
}

func PowerOnServer(name string)(resp *http.Response){
	req, _ := http.NewRequest("POST", WorkerAddress + "/minecraft/" + name,
		bytes.NewBufferString("{ \"Power\" : \"off\" } "))
        resp, _ = http.DefaultClient.Do(req)
	return resp
}

func PowerOffServer(name string)(resp *http.Response){
	req, _ := http.NewRequest("POST", WorkerAddress + "/minecraft/" + name,
		bytes.NewBufferString("{ \"Power\" : \"on\" } "))
	resp, _ = http.DefaultClient.Do(req)
	return resp
}

func DeleteMinecraftServer(name string)(resp *http.Response){
	req, _ := http.NewRequest("DELETE", WorkerAddress + "/minecraft/" + name, nil)
	resp, _ = http.DefaultClient.Do(req)
	return resp
}
