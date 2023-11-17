package keeper

import (
	"context"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mycel-domain/mycel/x/registry/types"
)

func (k Keeper) Role(goCtx context.Context, req *types.QueryRoleRequest) (*types.QueryRoleResponse, error) {
	var role types.DomainRole
	if req == nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request: empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	dms := strings.Split(req.DomainName, ".")
	switch len(dms) {
	case 1: // TLD
		tld, found := k.GetTopLevelDomain(ctx, dms[0])
		if !found {
			return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "domain not found")
		}
		role = tld.AccessControl[req.Address]
	case 2: // SLD
		sld, found := k.GetSecondLevelDomain(ctx, dms[0], dms[1])
		if !found {
			return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "domain not found")
		}
		role = sld.AccessControl[req.Address]
	default:
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request: domain name")
	}

	return &types.QueryRoleResponse{Role: role.String()}, nil
}
