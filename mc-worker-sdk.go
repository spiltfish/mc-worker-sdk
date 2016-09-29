package mc_worker_sdk

import (
	"net/http"
)

var ( WorkerAddress string )

func SetServer( workerAddress string){
	WorkerAddress = workerAddress
}

func GetMinecraftList()(resp http.Response){
	resp, _ = http.Get( WorkerAddress + "/minecraft")
	return resp
}

func GetMinecraftServer(name string)(resp http.Response){
	resp, _ = http.Get( WorkerAddress + "/minecraft/" + name)
	return resp
}

func CreateMinecraftServer(name string)(resp http.Response){
	resp, _ = http.NewRequest("PUT", WorkerAddress + "/minecraft/" + name, "")
	return resp
}

func PowerOnServer(name string)(resp http.Response){
	resp, _ = http.NewRequest("POST", WorkerAddress + "/minecraft/" + name,
		"{ \"Power\" : \"off\" } ")
	return resp
}

func PowerOffServer(name string)(resp http.Response){
	resp, _ = http.NewRequest("POST", WorkerAddress + "/minecraft/" + name,
		"{ \"Power\" : \"on\" } ")
	return resp
}

func DeleteMinecraftServer(name string)(resp http.Response){
	resp, _ = http.NewRequest("DELETE", WorkerAddress + "/minecraft/" + name, "")
	return resp
}