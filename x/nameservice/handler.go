package nameservice

import (
	"fmt"

	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		case MsgCreateProduct:
			return handleMsgCreateProduct(ctx, keeper, msg)
		case MsgUpdateProduct:
			return handleMsgUpdateProduct(ctx, keeper, msg)
		case MsgDeleteProduct:
			return handleMsgDeleteProduct(ctx, keeper, msg)
		case MsgBuyProduct:
			return handleMsgBuyProduct(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to set name
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) (*sdk.Result, error) {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil                // return
}

// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) (*sdk.Result, error) {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough") // If not, throw an error
	}
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return nil, err
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return &sdk.Result{}, nil
}

// Handle a message to delete name
func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg MsgDeleteName) (*sdk.Result, error) {
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.Name)
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.DeleteWhois(ctx, msg.Name)
	return &sdk.Result{}, nil
}

// Handle a message to create product
func handleMsgCreateProduct(ctx sdk.Context, keeper Keeper, msg MsgCreateProduct) (*sdk.Result, error) {
	key := "Product-" + msg.ProductID

	if keeper.IsProductPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrProductAlreadyExists, msg.ProductID)
	}

	var product = Product{
		ProductID:   msg.ProductID,
		Description: msg.Description,
		Price:       msg.Price,
		Owner:       msg.Signer,
	}

	keeper.SetProduct(ctx, key, product) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil            // return
}

// Handle a message to update product
func handleMsgUpdateProduct(ctx sdk.Context, keeper Keeper, msg MsgUpdateProduct) (*sdk.Result, error) {
	key := "Product-" + msg.ProductID

	if !keeper.IsProductPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
	}

	product := keeper.GetProduct(ctx, key)

	if !msg.Signer.Equals(product.Owner) { // Checks if the the msg signer is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	product.Description = msg.Description
	product.Price = msg.Price

	keeper.SetProduct(ctx, key, product) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil            // return
}

// Handle a message to delete product
func handleMsgDeleteProduct(ctx sdk.Context, keeper Keeper, msg MsgDeleteProduct) (*sdk.Result, error) {
	key := "Product-" + msg.ProductID

	if !keeper.IsProductPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.ProductID)
	}
	if !msg.Signer.Equals(keeper.GetProduct(ctx, key).Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.DeleteProduct(ctx, key)
	return &sdk.Result{}, nil
}

// Handle a message to buy product
func handleMsgBuyProduct(ctx sdk.Context, keeper Keeper, msg MsgBuyProduct) (*sdk.Result, error) {
	key := "Product-" + msg.ProductID

	if !keeper.IsProductPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.ProductID)
	}

	product := keeper.GetProduct(ctx, key)

	if msg.Signer.Equals(product.Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You are product owner")
	}

	err := keeper.CoinKeeper.SendCoins(ctx, msg.Signer, product.Owner, product.Price)
	if err != nil {
		return nil, err
	}

	product.Owner = msg.Signer

	keeper.SetProduct(ctx, key, product)
	return &sdk.Result{}, nil
}
