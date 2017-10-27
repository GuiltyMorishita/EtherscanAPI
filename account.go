package etherscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
)

// Get Ether Balance for a single Address
// https://api.etherscan.io/api?module=account&action=balance&address=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&tag=latest&apikey=YourApiKeyToken
type balRec struct {
	Status  string
	Message string
	Result  string
}

func (e *Etherscan) GetEtherBalance(addr string) (res *big.Int, err error) {
	var tr balRec
	var ok bool
	call := "http://api.etherscan.io/api?module=account&action=balance&address=" + addr + "&tag=latest&apikey=" + e.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return
	}
	if strings.Compare(tr.Status, "1") != 0 {
		err = errors.New(tr.Message)
		return
	}
	res, ok = strToWei(tr.Result)
	if !ok {
		err = errors.New("error in number " + tr.Result)
	}
	return
}

// Get Ether Balance for multiple Addresses in a single call
// https://api.etherscan.io/api?module=account&action=balancemulti&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a,0x63a9975ba31b0b9626b34300f7f627147df1f526,0x198ef1ec325a96cc354c7266a038be8b5c558f67&tag=latest&apikey=YourApiKeyToken

func (e *Etherscan) getMultiEtherBalances(addr []string) {
	var tr balRec
	addresses := strings.Join(addr, ",")
	call := "http://api.etherscan.io/api?module=account&action=balancemulti&address=" + addresses + "&tag=latest&apikey=" + e.apiKey
	fmt.Println(call)
	resp, err := http.Get(call)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tr)
}

// Get list of Blocks Mined by Address
// https://api.etherscan.io/api?module=account&action=getminedblocks&address=0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b&blocktype=blocks&apikey=YourApiKeyToken

// (To get paginated results use page=<page number> and offset=<max records to return>)
// ** type = blocks (full blocks only) or uncles (uncle blocks only)
// https://api.etherscan.io/api?module=account&action=getminedblocks&address=0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b&blocktype=blocks&page=1&offset=10&apikey=YourApiKeyToken