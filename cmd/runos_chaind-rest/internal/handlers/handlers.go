package handlers

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"runos_chain/cmd/runos_chaind-rest/internal/logging"
	"runos_chain/x/configstore/types"
)

const (
	constGetPortPath = "/runos_chain/configstore/hosts_database/%s/%s"
)

type ConfigList struct {
	RunosChainApi string `yaml:"runosChainApi"`
	LogLevel      string `yaml:"logLevel"`
	AccountName   string `yaml:"accountName"`
	Port          string `yaml:"port"`
}

var (
	chainRest = "localhost:1317"
	addr      string
	cosmos    cosmosclient.Client
	account   cosmosaccount.Account
	Config    ConfigList
)

func errorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf("500 - error %s", err)))
}

func Init() {
	logging.LoggerInit()
	configPath := flag.String("config", "rest-config.yml", "path to the config file")
	flag.Parse()
	file, err := os.Open(*configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&Config); err != nil {
		panic(err)
	}
	chainRest = Config.RunosChainApi
	level, err := logrus.ParseLevel(Config.LogLevel)
	if err != nil {
		panic(err)
	}
	logging.Logger.SetLevel(level)
	logging.Logger.Debugln("Seeting blockchain variables")
	addressPrefix := "cosmos"
	// Create a Cosmos client instance
	cosmos, err = cosmosclient.New(
		context.Background(),
		cosmosclient.WithAddressPrefix(addressPrefix),
	)
	if err != nil {
		log.Fatal(err)
	}
	accountName := Config.AccountName
	logging.Logger.Debugf("Account name is %s", accountName)
	account, err = cosmos.Account(accountName)
	if err != nil {
		fmt.Println("ok")
		log.Fatal(err)
	}
	addr, err = account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPortHandler(w http.ResponseWriter, r *http.Request) {
	dpid := r.URL.Query().Get("dpid")
	mac := r.URL.Query().Get("mac")
	path := chainRest + fmt.Sprintf(constGetPortPath, dpid, mac)
	path = "http://" + path
	logging.Logger.Debugf("Blockchain request url: %s\n", path)
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
		logging.Logger.Errorln(err)
	}
	logging.Logger.Debugln(txResp)
}
