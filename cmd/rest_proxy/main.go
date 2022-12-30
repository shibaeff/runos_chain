package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	//types2 "github.com/tendermint/spn/x/launch/types"
	"io"
	"log"
	"net/http"
	//"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/flags"
	//"github.com/cosmos/cosmos-sdk/client/tx"
	//"github.com/spf13/cobra"
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

/*
Sample json

	{
		dpid: 0000000000000001,
		mac: 00:00:5e:00:53:af,
		inport: 345
	}
*/
//type SetPortRequest struct {
//	Dpid   string `json:"dpid"`
//	Mac    string `json:"mac"`
//	Inport string `json:"inport"`
//}

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

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := cosmos.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	// Define a message to create a post
	//s, err := io.ReadAll(r.Body)
	//if err != nil {
	//	errorHandler(w, err)
	//	return
	//}
	dpid := r.URL.Query().Get("dpid")
	mac := r.URL.Query().Get("mac")
	inport := r.URL.Query().Get("inport")
	msg := &types.MsgCreateHostsDatabase{
		Creator: addr,
		Dpid:    dpid,
		Mac:     mac,
		Inport:  inport,
	}
	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := cosmos.BroadcastTx(context.Background(), account, msg)
	if err != nil {
		fmt.Println(err)
	}

	// Print response from broadcasting a transaction
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
