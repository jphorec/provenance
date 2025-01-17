package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	cosmosauthtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/provenance-io/provenance/x/msgfees/types"
)

// StoreKey is the store key string for authz
const StoreKey = types.ModuleName

type baseAppSimulateFunc func(txBytes []byte) (sdk.GasInfo, *sdk.Result, sdk.Context, error)

// Keeper of the Additional fee store
type Keeper struct {
	storeKey         sdk.StoreKey
	cdc              codec.BinaryCodec
	paramSpace       paramtypes.Subspace
	feeCollectorName string // name of the FeeCollector ModuleAccount
	defaultFeeDenom  string
	simulateFunc     baseAppSimulateFunc
	txDecoder        sdk.TxDecoder
}

// NewKeeper returns a AdditionalFeeKeeper. It handles:
// CONTRACT: the parameter Subspace must have the param key table already initialized
func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	feeCollectorName string,
	defaultFeeDenom string,
	simulateFunc baseAppSimulateFunc,
	txDecoder sdk.TxDecoder,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:         key,
		cdc:              cdc,
		paramSpace:       paramSpace,
		feeCollectorName: feeCollectorName,
		defaultFeeDenom:  defaultFeeDenom,
		simulateFunc:     simulateFunc,
		txDecoder:        txDecoder,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) GetFeeCollectorName() string {
	return k.feeCollectorName
}

// GetFloorGasPrice  returns the current minimum gas price in sdk.Coin used in calculations for charging additional fees
func (k Keeper) GetFloorGasPrice(ctx sdk.Context) sdk.Coin {
	min := types.DefaultFloorGasPrice
	if k.paramSpace.Has(ctx, types.ParamStoreKeyFloorGasPrice) {
		k.paramSpace.Get(ctx, types.ParamStoreKeyFloorGasPrice, &min)
	}
	return min
}

// SetMsgFee sets the additional fee schedule for a Msg
func (k Keeper) SetMsgFee(ctx sdk.Context, msgFees types.MsgFee) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&msgFees)
	store.Set(types.GetMsgFeeKey(msgFees.MsgTypeUrl), bz)
	return nil
}

// GetMsgFee returns a MsgFee for the msg type if it exists nil if it does not
func (k Keeper) GetMsgFee(ctx sdk.Context, msgType string) (*types.MsgFee, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetMsgFeeKey(msgType)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, nil
	}

	var msgFee types.MsgFee
	if err := k.cdc.Unmarshal(bz, &msgFee); err != nil {
		return nil, err
	}

	return &msgFee, nil
}

// RemoveMsgFee removes MsgFee or returns an error if it does not exist
func (k Keeper) RemoveMsgFee(ctx sdk.Context, msgType string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetMsgFeeKey(msgType)
	bz := store.Get(key)
	if len(bz) == 0 {
		return types.ErrMsgFeeDoesNotExist
	}

	store.Delete(key)

	return nil
}

type Handler func(record types.MsgFee) (stop bool)

// IterateMsgFees  iterates all msg fees with the given handler function.
func (k Keeper) IterateMsgFees(ctx sdk.Context, handle func(msgFees types.MsgFee) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.MsgFeeKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.MsgFee{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// DeductFees deducts fees from the given account, the only reason it exists is that the
// cosmos method does not take in the custom fee collector which is a feature desired from msg fees.
func (k Keeper) DeductFees(bankKeeper cosmosauthtypes.BankKeeper, ctx sdk.Context, acc cosmosauthtypes.AccountI, fees sdk.Coins) error {
	if !fees.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %q", fees)
	}

	err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), k.feeCollectorName, fees)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}
	return nil
}
