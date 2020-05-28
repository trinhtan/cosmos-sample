package nameservice

import (
	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/keeper"
	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetName    = types.NewMsgSetName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec

	NewProduct          = types.NewProduct
	NewMsgCreateProduct = types.NewMsgCreateProduct
	NewMsgUpdateProduct = types.NewMsgUpdateProduct
	NewMsgDeleteProduct = types.NewMsgDeleteProduct
	NewMsgBuyProduct    = types.NewMsgBuyProduct
)

type (
	Keeper          = keeper.Keeper
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgDeleteName   = types.MsgDeleteName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois

	Product             = types.Product
	MsgCreateProduct    = types.MsgCreateProduct
	MsgUpdateProduct    = types.MsgUpdateProduct
	MsgDeleteProduct    = types.MsgDeleteProduct
	MsgBuyProduct       = types.MsgBuyProduct
	QueryResAllProducts = types.QueryResAllProducts
)
