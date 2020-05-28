package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "nameservice/DeleteName", nil)

	cdc.RegisterConcrete(MsgCreateProduct{}, "nameservice/CreateProduct", nil)
	cdc.RegisterConcrete(MsgUpdateProduct{}, "nameservice/UpdateProduct", nil)
	cdc.RegisterConcrete(MsgDeleteProduct{}, "nameservice/DeleteProduct", nil)
	cdc.RegisterConcrete(MsgBuyProduct{}, "nameservice/BuyProduct", nil)
}
