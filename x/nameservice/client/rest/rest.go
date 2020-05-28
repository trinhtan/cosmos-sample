package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "name"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), namesHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), buyNameHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), setNameHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}", storeName, restName), resolveNameHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}/whois", storeName, restName), whoIsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), deleteNameHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc(fmt.Sprintf("/%s/product", storeName), createProductHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/product", storeName), updateProductHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/product/buyProduct", storeName), buyProductHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/product/{productID}", storeName), queryProductHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/product", storeName), allProductsHandler(cliCtx, storeName)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/name/{name}/address", storeName), accAddressHandler(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/tx/sign", storeName), signTxHandler(cliCtx)).Methods("POST")
}
