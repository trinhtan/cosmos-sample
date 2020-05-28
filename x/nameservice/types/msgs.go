package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetName) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Route should return the name of the module
func (msg MsgBuyName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyName) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyName) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// MsgDeleteName defines a DeleteName message
type MsgDeleteName struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgDeleteName is a constructor function for MsgDeleteName
func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteName {
	return MsgDeleteName{
		Name:  name,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgDeleteName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteName) Type() string { return "delete_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgCreateProduct defines a CreateProduct message
type MsgCreateProduct struct {
	ProductID   string         `json:"productID"`
	Description string         `json:"description"`
	Price       sdk.Coins      `json:"price"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgCreateProduct is a constructor function for MsgCreateProduct
func NewMsgCreateProduct(productID string, description string, price sdk.Coins, signer sdk.AccAddress) MsgCreateProduct {
	return MsgCreateProduct{
		ProductID:   productID,
		Description: description,
		Price:       price,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgCreateProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateProduct) Type() string { return "create_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Description) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Description cannot be empty")
	}
	if !msg.Price.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgCreateProduct defines a UpdateProduct message
type MsgUpdateProduct struct {
	ProductID   string         `json:"productID"`
	Description string         `json:"description"`
	Price       sdk.Coins      `json:"price"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgUpdateProduct is a constructor function for MsgUpdateProduct
func NewMsgUpdateProduct(productID string, description string, price sdk.Coins, signer sdk.AccAddress) MsgUpdateProduct {
	return MsgUpdateProduct{
		ProductID:   productID,
		Description: description,
		Price:       price,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgUpdateProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUpdateProduct) Type() string { return "update_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Description) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Description cannot be empty")
	}
	if !msg.Price.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgDeleteProduct defines a DeleteName message
type MsgDeleteProduct struct {
	ProductID string         `json:"productID"`
	Signer    sdk.AccAddress `json:"signer"`
}

// NewMsgDeleteProduct is a constructor function for MsgDeleteProduct
func NewMsgDeleteProduct(productID string, signer sdk.AccAddress) MsgDeleteProduct {
	return MsgDeleteProduct{
		ProductID: productID,
		Signer:    signer,
	}
}

// Route should return the name of the module
func (msg MsgDeleteProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteProduct) Type() string { return "delete_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgBuyProduct defines a DeleteName message
type MsgBuyProduct struct {
	ProductID string         `json:"productID"`
	Signer    sdk.AccAddress `json:"signer"`
}

// NewMsgBuyProduct is a constructor function for MsgBuyProduct
func NewMsgBuyProduct(productID string, signer sdk.AccAddress) MsgBuyProduct {
	return MsgBuyProduct{
		ProductID: productID,
		Signer:    signer,
	}
}

// Route should return the name of the module
func (msg MsgBuyProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyProduct) Type() string { return "buy_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
