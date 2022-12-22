package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/ethereum/go-ethereum"
	"log"
	"net/http"
	"time"
)

/*
curl https://eth-mainnet.g.alchemy.com/v2/GrAWOBLihNvOotLnNfM2P-ygl8lRj_SK \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}'
*/

type NodeProvider struct {
	Name string   `json:"name"`
	Key  string   `json:"key"`
	Urls []string `json:"urls"`
}

type RPCCall struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int      `json:"id"`
}

func runTest(provider NodeProvider, call RPCCall) {
	byt, _ := json.Marshal(call)

	key := provider.Key
	log.Println("key:", key)
	log.Println("json:", string(byt))

	for _, url := range provider.Urls {
		request, error := http.NewRequest("POST", fmt.Sprintf("%s/%s", url, key), bytes.NewBuffer(byt))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		// start time
		startTime := time.Now()

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			//panic(error)
		}
		defer response.Body.Close()

		// end time
		duration := time.Since(startTime)

		//body, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode == 200 {
			log.Println(url)
			log.Println(duration)
			log.Println(duration.Nanoseconds())
		}
		/*log.Println(url)
		fmt.Println("R:", string(body))
		log.Println(response.StatusCode)
		log.Println(duration)
		log.Println(duration.Nanoseconds())*/
	}
}

func main() {
	log.Println("w3npet")

	infura := NodeProvider{
		Name: "Infura",
		Key:  "086d385e5266474fac356d54920e1e60",
		Urls: []string{
			"https://mainnet.infura.io/v3",
			"https://avalanche-mainnet.infura.io/v3",
			"https://starknet-mainnet.infura.io/v3",
			"https://palm-mainnet.infura.io/v3",
			"https://near-mainnet.infura.io/v3",
			"https://celo-mainnet.infura.io/v3",
		},
	}

	alchemy := NodeProvider{
		Name: "Alchemy",
		Key:  "GrAWOBLihNvOotLnNfM2P-ygl8lRj_SK",
		Urls: []string{
			"https://eth-mainnet.g.alchemy.com/v2",
		},
	}

	call := RPCCall{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []string{},
		Id:      1,
	}

	log.Println("infura", infura)
	log.Println("alchemy", infura)
	runTest(infura, call)
	runTest(alchemy, call)
}
