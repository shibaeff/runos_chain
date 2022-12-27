package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	"runos_chain/x/configstore/types"
)

const (
	constGetPortPath = "/runos_chain/config_store/hosts_database/%s/%s"
)

var (
	chainRest string
	addr      string
	cosmos    cosmosclient.Client
	account   cosmosaccount.Account
)

func errorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf("500 - error %s", err)))
}

func Init() {
	chainRest = "localhost"
	cosmos, err := cosmosclient.New(
		context.Background(),
	)
	accountName := "Alice"
	if err != nil {
		log.Fatal(err)
	}
	account, err := cosmos.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err = account.Address("")
	if err != nil {
		log.Fatal(err)
	}
}

func GetPortHandler(w http.ResponseWriter, r *http.Request) {
	dpid := r.URL.Query().Get("dpid")
	mac := r.URL.Query().Get("mac")
	path, _ := url.JoinPath(chainRest, fmt.Sprintf(constGetPortPath, dpid, mac))
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

type SetPortRequest struct {
	Dpid   string
	Mac    string
	Inport string
}

func SetPortHandler(w http.ResponseWriter, r *http.Request) {
	s, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, err)
		return
	}
	var tuple SetPortRequest
	err = json.Unmarshal(s, &tuple)
	if err != nil {
		errorHandler(w, err)
		return
	}
	msg := &types.MsgSetPort{
		Creator: addr,
		Dpid:    tuple.Dpid,
		Mac:     tuple.Mac,
		Inport:  tuple.Inport,
	}
	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := cosmos.BroadcastTx(context.Background(), account, msg)
	if err != nil {
		log.Fatal(err)
	}
	// Print response from broadcasting a transaction
	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)
}

func main() {
	// mux handlers
	r := mux.NewRouter()
	r.HandleFunc("/getPort", GetPortHandler).Methods("GET")
	r.HandleFunc("/setPort", SetPortHandler).Methods("POST")
	http.Handle("/", r)
}
