package client

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	qcliacc "github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/module/stake/mapper"
	"github.com/QOSGroup/litewallet/litewallet/slim/module/stake/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/libs/common"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/client"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const (
	flagActive   = "active"
	activeDesc   = "active"
	inactiveDesc = "inactive"

	inactiveRevokeDesc        = "Revoked"
	inactiveMissVoteBlockDesc = "Kicked"
	inactiveMaxValidatorDesc  = "Replaced"
	inactiveDoubleDesc        = "DoubleSign"
)

type validatorDisplayInfo struct {
	Owner           btypes.Address    `json:"owner"`
	ValidatorAddr   string            `json:"validatorAddress"`
	ValidatorPubKey crypto.PubKey     `json:"validatorPubkey"`
	BondTokens      uint64            `json:"bondTokens"` //不能超过int64最大值
	Description     types.Description `json:"description"`
	Commission      types.Commission  `json:"commission"`

	Status         string    `json:"status"`
	InactiveDesc   string    `json:"InactiveDesc"`
	InactiveTime   time.Time `json:"inactiveTime"`
	InactiveHeight uint64    `json:"inactiveHeight"`

	BondHeight uint64 `json:"bondHeight"`
}

func toValidatorDisplayInfo(validator types.Validator) validatorDisplayInfo {
	info := validatorDisplayInfo{
		Owner:           validator.Owner,
		ValidatorPubKey: validator.ValidatorPubKey,
		BondTokens:      validator.BondTokens,
		Description:     validator.Description,
		InactiveTime:    validator.InactiveTime,
		InactiveHeight:  validator.InactiveHeight,
		BondHeight:      validator.BondHeight,
		Commission:      validator.Commission,
	}

	if validator.Status == types.Active {
		info.Status = activeDesc
	} else {
		info.Status = inactiveDesc
	}

	if validator.InactiveCode == types.Revoke {
		info.InactiveDesc = inactiveRevokeDesc
	} else if validator.InactiveCode == types.MissVoteBlock {
		info.InactiveDesc = inactiveMissVoteBlockDesc
	} else if validator.InactiveCode == types.MaxValidator {
		info.InactiveDesc = inactiveMaxValidatorDesc
	} else if validator.InactiveCode == types.DoubleSign {
		info.InactiveDesc = inactiveDoubleDesc
	}

	info.ValidatorAddr = strings.ToUpper(hex.EncodeToString(validator.ValidatorPubKey.Address()))

	return info
}

func QueryValidatorInfo(remote, address string) (types.Validator, error) {
	cliCtx := context.NewCLIContext(remote)
	ownerAddress, err := qcliacc.GetAddrFromValue(address)
	if err != nil {
		return types.Validator{}, err
	}
	validator, err := getValidator(cliCtx, ownerAddress)
	return validator, nil
}

func getValidator(ctx context.CLIContext, ownerAddress btypes.Address) (types.Validator, error) {
	result, err := ctx.Client.ABCIQueryWithOptions(string(types.BuildStakeStoreQueryPath()), types.BuildOwnerWithValidatorKey(ownerAddress), buildQueryOptions())
	if err != nil {
		return types.Validator{}, err
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		return types.Validator{}, errors.New("owner does't have validator")
	}

	var address btypes.Address
	txs.Cdc.UnmarshalBinaryBare(valueBz, &address)

	key := types.BuildValidatorKey(address)
	result, err = ctx.Client.ABCIQueryWithOptions(string(types.BuildStakeStoreQueryPath()), key, buildQueryOptions())
	if err != nil {
		return types.Validator{}, err
	}

	valueBz = result.Response.GetValue()
	if len(valueBz) == 0 {
		return types.Validator{}, errors.New("response empty value")
	}

	var validator types.Validator
	txs.Cdc.UnmarshalBinaryBare(valueBz, &validator)
	return validator, nil
}

func buildQueryOptions() client.ABCIQueryOptions {
	//height := viper.GetInt64(bctypes.FlagHeight)
	//if height <= 0 {
	//	height = 0
	//}

	//trust := viper.GetBool(bctypes.FlagTrustNode)

	return client.ABCIQueryOptions{
		Height: 0,
		Prove:  true,
	}
}

func QueryValidators(remote string) ([]byte, error) {
	cliCtx := context.NewCLIContext(remote)

	opts := client.ABCIQueryOptions{
		Height: 0,
		Prove:  true,
	}

	subspace := "/store/validator/subspace"
	result, err := cliCtx.Client.ABCIQueryWithOptions(subspace, []byte{0x01}, opts)

	if err != nil {
		return nil, err
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		return nil, errors.New("response empty value")
	}

	var validators []validatorDisplayInfo

	var vKVPair []common.KVPair
	err = txs.Cdc.UnmarshalBinaryLengthPrefixed(valueBz, &vKVPair)
	if err != nil {
		return nil, err
	}
	for _, kv := range vKVPair {
		var validator types.Validator
		fmt.Println(kv.Value)
		txs.Cdc.UnmarshalBinaryBare(kv.Value, &validator)
		validators = append(validators, toValidatorDisplayInfo(validator))
	}
	return txs.Cdc.MarshalJSON(validators)
}

func QueryValidatorMissedVoteInfo(remote, address string) (voteSummary, error) {
	cliCtx := context.NewCLIContext(remote).WithCodec(txs.Cdc)

	ownerAddress, err := qcliacc.GetAddrFromValue(address)
	if err != nil {
		return voteSummary{}, err
	}

	return queryVotesInfoByOwner(cliCtx, ownerAddress)
}

type voteSummary struct {
	StartHeight int64            `json:"startHeight"`
	EndHeight   int64            `json:"endHeight"`
	MissCount   int8             `json:"missCount"`
	Votes       []voteInfoDetail `json:"voteDetail"`
}

type voteInfoDetail struct {
	Height int64
	Vote   bool
}

func queryVotesInfoByOwner(ctx context.CLIContext, ownerAddress btypes.Address) (voteSummary, error) {
	voteSummaryDisplay := voteSummary{}

	windownLength, err := getStakeConfig(ctx)
	if err != nil {
		return voteSummaryDisplay, err
	}

	votesInfo := make([]voteInfoDetail, 0, windownLength)

	val, err := getValidator(ctx, ownerAddress)
	if err != nil {
		return voteSummaryDisplay, err
	}

	validatorAddress := btypes.Address(val.ValidatorPubKey.Address())

	voteInfo, err := getValidatorVoteInfo(ctx, validatorAddress)
	if err != nil {
		return voteSummaryDisplay, err
	}

	voteInfoMap, _, err := queryValidatorVotesInWindow(ctx, validatorAddress)
	if err != nil {
		return voteSummaryDisplay, err
	}

	votedBlockLength := voteInfo.IndexOffset - 1

	endWindowHeight := voteInfo.StartHeight + votedBlockLength
	startWindowHeight := uint64(1)
	if votedBlockLength <= windownLength {
		startWindowHeight = voteInfo.StartHeight
	} else {
		startWindowHeight = endWindowHeight - windownLength + 1
	}

	voteSummaryDisplay.StartHeight = int64(startWindowHeight)
	voteSummaryDisplay.EndHeight = int64(endWindowHeight)

	i := int8(0)
	for h := endWindowHeight; h >= startWindowHeight; h-- {
		index := h % windownLength
		voted := true

		if v, ok := voteInfoMap[index]; ok {
			voted = v
		}

		if !voted {
			i++
			votesInfo = append(votesInfo, voteInfoDetail{int64(h), voted})
		}
	}

	voteSummaryDisplay.Votes = votesInfo
	voteSummaryDisplay.MissCount = i
	return voteSummaryDisplay, nil
}

func getStakeConfig(ctx context.CLIContext) (uint64, error) {
	node, err := ctx.GetNode()
	if err != nil {
		return 0, err
	}

	path := "/store/params/key"
	key := BuildParamKey(types.ParamSpace, types.KeyValidatorVotingStatusLen)

	result, err := node.ABCIQueryWithOptions(path, key, buildQueryOptions())
	if err != nil {
		return 0, err
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		return 0, errors.New("response empty value. getStakeConfig is empty")
	}

	var length uint64
	ctx.Codec.UnmarshalBinaryBare(valueBz, &length)

	return length, nil

	return 0, nil
}

func BuildParamKey(paramSpace string, key []byte) []byte {
	return append([]byte(paramSpace), key...)
}

func getValidatorVoteInfo(ctx context.CLIContext, validatorAddr btypes.Address) (types.ValidatorVoteInfo, error) {
	node, err := ctx.GetNode()
	if err != nil {
		return types.ValidatorVoteInfo{}, err
	}

	path := string(types.BuildStakeStoreQueryPath())
	key := types.BuildValidatorVoteInfoKey(validatorAddr)

	result, err := node.ABCIQueryWithOptions(path, key, buildQueryOptions())
	if err != nil {
		return types.ValidatorVoteInfo{}, err
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		return types.ValidatorVoteInfo{}, errors.New("response empty value. validatorVoteInfo is empty")
	}

	var voteInfo types.ValidatorVoteInfo
	ctx.Codec.UnmarshalBinaryBare(valueBz, &voteInfo)

	return voteInfo, nil
}

func queryValidatorVotesInWindow(ctx context.CLIContext, validatorAddr btypes.Address) (map[uint64]bool, int64, error) {

	voteInWindowInfo := make(map[uint64]bool)

	node, err := ctx.GetNode()
	if err != nil {
		return voteInWindowInfo, 0, err
	}

	storePath := "/" + strings.Join([]string{"store", types.MapperName, "subspace"}, "/")
	key := types.BuildValidatorVoteInfoInWindowPrefixKey(validatorAddr)

	result, err := node.ABCIQueryWithOptions(storePath, key, buildQueryOptions())
	if err != nil {
		return nil, 0, err
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		return voteInWindowInfo, result.Response.Height, nil
	}

	var vKVPair []common.KVPair
	ctx.Codec.UnmarshalBinaryLengthPrefixed(valueBz, &vKVPair)

	for _, kv := range vKVPair {
		k := kv.Key
		var vote bool
		index := binary.LittleEndian.Uint64(k[(len(k) - 8):])
		ctx.Codec.UnmarshalBinaryBare(kv.Value, &vote)
		voteInWindowInfo[index] = vote
	}

	return voteInWindowInfo, result.Response.Height, nil
}

func QueryDelegationInfo(remote, ownerAddr, delegatorAddr string) (mapper.DelegationQueryResult, error) {
	cliCtx := context.NewCLIContext(remote).WithCodec(txs.Cdc)

	var owner btypes.Address
	var delegator btypes.Address

	if o, err := qcliacc.GetAddrFromValue(ownerAddr); err == nil {
		owner = o
	}

	if d, err := qcliacc.GetAddrFromValue(delegatorAddr); err == nil {
		delegator = d
	}

	var path = types.BuildGetDelegationCustomQueryPath(delegator, owner)

	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return mapper.DelegationQueryResult{}, err
	}

	var result mapper.DelegationQueryResult
	err = txs.Cdc.UnmarshalJSON(res, &result)
	return result, err
}

func QueryDelegations(remote, address string) ([]mapper.DelegationQueryResult, error) {
	cliCtx := context.NewCLIContext(remote).WithCodec(txs.Cdc)
	var delegator btypes.Address

	if d, err := qcliacc.GetAddrFromValue(address); err == nil {
		delegator = d
	}

	var path = types.BuildQueryDelegationsByDelegatorCustomQueryPath(delegator)

	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return nil, err
	}

	var result []mapper.DelegationQueryResult
	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	return result, err
}

func QueryUnbondings(remote, address string) ([]types.UnbondingDelegationInfo, error) {
	cliCtx := context.NewCLIContext(remote).WithCodec(txs.Cdc)

	var delegator btypes.Address

	if o, err := qcliacc.GetAddrFromValue(address); err == nil {
		delegator = o
	}

	var path = types.BuildQueryUnbondingsByDelegatorCustomQueryPath(delegator)

	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return nil, err
	}

	var result []types.UnbondingDelegationInfo
	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	return result, err
}

func QueryRedelegations(remote, address string) ([]types.RedelegationInfo, error) {
	cliCtx := context.NewCLIContext(remote).WithCodec(txs.Cdc)

	var delegator btypes.Address

	if o, err := qcliacc.GetAddrFromValue(address); err == nil {
		delegator = o
	}

	var path = types.BuildQueryRedelegationsByDelegatorCustomQueryPath(delegator)

	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return nil, err
	}

	var result []types.RedelegationInfo
	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	return result, err
}
