package simapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

// used to satisfy the interface for transfer, contract should escrow/unescrow.
type noopBankKeeper struct{}

// Implements the bank keeper interface
var _ ibctransfertypes.BankKeeper = noopBankKeeper{}

// NewBankKeeper returns a no-op bank keeper
func NewNoOpBankKeeper() ibctransfertypes.BankKeeper {
	return noopBankKeeper{}
}

// Implement the bank keeper interface as no-ops

func (k noopBankKeeper) SendCoins(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	return nil
}

func (k noopBankKeeper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	return nil
}

func (k noopBankKeeper) BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	return nil
}

func (k noopBankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return nil
}

// SendCoinsFromAccountToModule sends coins from an account to a module
func (k noopBankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return nil
}

// BlockedAddr returns if an address is blocked
func (k noopBankKeeper) BlockedAddr(addr sdk.AccAddress) bool {
	return false
}

// IsSendEnabledCoin returns if a coin can be sent
func (k noopBankKeeper) IsSendEnabledCoin(ctx sdk.Context, coin sdk.Coin) bool {
	return true
}

// GetBalance returns the balance of a single coin
func (k noopBankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return sdk.Coin{}
}

// GetAllBalances returns all the balances of an account
func (k noopBankKeeper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return sdk.Coins{}
}
