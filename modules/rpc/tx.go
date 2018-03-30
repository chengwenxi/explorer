package rpc

import (
	"github.com/gorilla/mux"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk"
	"net/http"
	"github.com/tendermint/tmlibs/common"
	"github.com/irisnet/irisplorer.io/modules/tools"
	"github.com/irisnet/irisplorer.io/modules/store/m"
	"github.com/spf13/cast"
)

func registerQueryTx(r *mux.Router) error {
	r.HandleFunc("/tx/{hash}", queryTx).Methods("GET")
	return nil
}

func registerQueryAllCoinTxByPage(r *mux.Router) error {
	r.HandleFunc("/tx/coin/{page}", queryAllCoinTxByPage).Methods("GET")
	return nil
}

func registerQueryCoinTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/coin/{address}", queryCoinTxByAccount).Methods("GET")
	return nil
}

func registerQueryPageCoinTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/coin/{address}/{page}", queryCoinPageTxByAccount).Methods("GET")
	return nil
}

func registerQueryAllStakeTxByPage(r *mux.Router) error {
	r.HandleFunc("/tx/stake/{page}", queryAllStakeTxByPage).Methods("GET")
	return nil
}

func registerQueryStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/stake/{address}", queryStakeTxByAccount).Methods("GET")
	return nil
}

func registerQueryPageStakeTxByAccount(r *mux.Router) error {
	r.HandleFunc("/tx/stake/{address}/{page}", queryPageStakeTxByAccount).Methods("GET")
	return nil
}

// queryTx is to query transaction by hash
func queryTx(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	hash := args["hash"]

	key, err := hex.DecodeString(common.StripHex(hash))
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	node := tools.GetNode()
	defer node.Release()

	res, err := node.Client.Tx(key, false)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	wrap, err := tools.BuildTxResp(res.Height, res.Proof.Data, false, hash)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}

	tools.FmtOutPutResult(w, wrap)
}

func queryAllCoinTxByPage(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	page := args["page"]
	p := cast.ToInt(page)
	result := m.QueryAllPageCoinTxs(p)
	tools.FmtOutPutResult(w, result)
}

func queryCoinTxByAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	account := args["address"]
	result := m.QueryCoinTxsByAccount(account)
	tools.FmtOutPutResult(w, result)
}

func queryCoinPageTxByAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	account := args["address"]
	page := args["page"]
	p := cast.ToInt(page)
	result := m.QueryPageCoinTxsByAccount(account,p)
	tools.FmtOutPutResult(w, result)
}

func queryAllStakeTxByPage(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	page := args["page"]
	p := cast.ToInt(page)
	result := m.QueryPageStakeTxs(p)
	tools.FmtOutPutResult(w, result)
}


func queryStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	account := args["address"]

	result := m.QueryStakeTxsByAccount(account)
	tools.FmtOutPutResult(w, result)
}

func queryPageStakeTxByAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	account := args["address"]
	page := args["page"]
	p := cast.ToInt(page)
	result := m.QueryPageStakeTxsByAccount(account,p)
	tools.FmtOutPutResult(w, result)
}

func RegisterTx(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryTx,
		registerQueryCoinTxByAccount,
		registerQueryPageCoinTxByAccount,
		registerQueryAllCoinTxByPage,
		registerQueryStakeTxByAccount,
		registerQueryPageStakeTxByAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}
