package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	"runos_chain/x/configstore/types"
)

const (
	constGetPortPath = "/runos_chain/configstore/hosts_database/%s/%s"
)

var (
	chainRest = "localhost:1317"
	addr      string
	cosmos    cosmosclient.Client
	account   cosmosaccount.Account
)

func errorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf("500 - error %s", err)))
}

//func Init() {
//	chainRest = "localhost"
//	prefix := "cosmos"
//	accountName := "alice"
//	var ar types2.AccountRetriever
//	cosmos, err := cosmosclient.New(
//		context.Background(),
//		cosmosclient.WithAddressPrefix(prefix),
//		cosmosclient.WithAccountRetriever(ar),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	account, err = cosmos.Account(accountName)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	addr, err = account.Address(prefix)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func GetPortHandler(w http.ResponseWriter, r *http.Request) {
	dpid := r.URL.Query().Get("dpid")
	mac := r.URL.Query().Get("mac")
	path := chainRest + fmt.Sprintf(constGetPortPath, dpid, mac)
	path = "http://" + path
	fmt.Println(path)
	resp, err := http.Get(path)
	if err != nil {
		errorHandler(w, err)
		return
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
	resp.Body.Close()
}
func SetPortHandler(w http.ResponseWriter, r *http.Request) {
	addressPrefix := "cosmos"

	// Create a Cosmos client instance
	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithAddressPrefix(addressPrefix),
	)
	if err != nil {
		log.Fatal(err)
	}
	accountName := "alice"
	account, err := cosmos.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	dpid := r.URL.Query().Get("dpid")
	mac := r.URL.Query().Get("mac")
	inport := r.URL.Query().Get("inport")
	msg := &types.MsgCreateHostsDatabase{
		Creator: addr,
		Dpid:    dpid,
		Mac:     mac,
		Inport:  inport,
	}
	txResp, err := cosmos.BroadcastTx(context.Background(), account, msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("MsgCreate:\n\n")
	fmt.Println(txResp)
}

func main() {
	// mux handlers
	fmt.Println("Started")
	r := mux.NewRouter()
	r.HandleFunc("/getPort", GetPortHandler).Methods("GET")
	r.HandleFunc("/setPort", SetPortHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
