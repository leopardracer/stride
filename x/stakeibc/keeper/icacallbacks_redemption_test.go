package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	_ "github.com/stretchr/testify/suite"

	icacallbacktypes "github.com/Stride-Labs/stride/v27/x/icacallbacks/types"
	recordtypes "github.com/Stride-Labs/stride/v27/x/records/types"

	"github.com/Stride-Labs/stride/v27/x/stakeibc/types"
)

type RedemptionCallbackState struct {
	epochUnbondingNumbers   []uint64
	userRedemptionRecordIds []string
	epochNumber             uint64
	nativeTokenAmount       sdkmath.Int
}

type RedemptionCallbackArgs struct {
	packet      channeltypes.Packet
	ackResponse *icacallbacktypes.AcknowledgementResponse
	args        []byte
}

type RedemptionCallbackTestCase struct {
	initialState RedemptionCallbackState
	validArgs    RedemptionCallbackArgs
}

func (s *KeeperTestSuite) SetupRedemptionCallback() RedemptionCallbackTestCase {
	epochNumber := uint64(1)

	// individual userRedemptionRecords should be claimable, as long as the host zone unbonding allows for claims
	recordId1 := recordtypes.UserRedemptionRecordKeyFormatter(HostChainId, epochNumber, "receiver")
	userRedemptionRecord1 := recordtypes.UserRedemptionRecord{
		Id: recordId1,
	}
	recordId2 := recordtypes.UserRedemptionRecordKeyFormatter(HostChainId, epochNumber, "other_receiver")
	userRedemptionRecord2 := recordtypes.UserRedemptionRecord{
		Id: recordId2,
	}

	// the hostZoneUnbonding should have HostZoneUnbonding_EXIT_TRANSFER_QUEUE - meaning unbonding has completed, but the tokens
	// have not yet been transferred to the redemption account
	nativeTokenAmount := sdkmath.NewInt(1000)
	hostZoneUnbonding := recordtypes.HostZoneUnbonding{
		HostZoneId:            HostChainId,
		Status:                recordtypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE,
		NativeTokenAmount:     nativeTokenAmount,
		UserRedemptionRecords: []string{recordId1, recordId2},
	}

	epochUnbondingRecord := recordtypes.EpochUnbondingRecord{
		EpochNumber:        epochNumber,
		HostZoneUnbondings: []*recordtypes.HostZoneUnbonding{&hostZoneUnbonding},
	}
	hostZone := types.HostZone{
		ChainId:        HostChainId,
		HostDenom:      Atom,
		IbcDenom:       IbcAtom,
		RedemptionRate: sdk.NewDec(1.0),
	}
	s.App.StakeibcKeeper.SetHostZone(s.Ctx, hostZone)
	s.App.RecordsKeeper.SetEpochUnbondingRecord(s.Ctx, epochUnbondingRecord)
	s.App.RecordsKeeper.SetUserRedemptionRecord(s.Ctx, userRedemptionRecord1)
	s.App.RecordsKeeper.SetUserRedemptionRecord(s.Ctx, userRedemptionRecord2)

	packet := channeltypes.Packet{}
	ackResponse := icacallbacktypes.AcknowledgementResponse{Status: icacallbacktypes.AckResponseStatus_SUCCESS}
	callbackArgs := types.RedemptionCallback{
		HostZoneId:              HostChainId,
		EpochUnbondingRecordIds: []uint64{epochNumber},
	}
	callbackArgsBz, err := s.App.StakeibcKeeper.MarshalRedemptionCallbackArgs(s.Ctx, callbackArgs)
	s.Require().NoError(err)

	return RedemptionCallbackTestCase{
		initialState: RedemptionCallbackState{
			epochUnbondingNumbers:   []uint64{epochNumber},
			userRedemptionRecordIds: []string{userRedemptionRecord1.Id, userRedemptionRecord2.Id},
			epochNumber:             epochNumber,
			nativeTokenAmount:       nativeTokenAmount,
		},
		validArgs: RedemptionCallbackArgs{
			packet:      packet,
			ackResponse: &ackResponse,
			args:        callbackArgsBz,
		},
	}
}

func (s *KeeperTestSuite) TestRedemptionCallback_Successful() {
	tc := s.SetupRedemptionCallback()
	initialState := tc.initialState
	validArgs := tc.validArgs

	err := s.App.StakeibcKeeper.RedemptionCallback(s.Ctx, validArgs.packet, validArgs.ackResponse, validArgs.args)
	s.Require().NoError(err, "redemption callback succeeded")

	for _, epochNumber := range initialState.epochUnbondingNumbers {
		// fetch the epoch unbonding record
		epochUnbondingRecord, found := s.App.RecordsKeeper.GetEpochUnbondingRecord(s.Ctx, epochNumber)
		s.Require().True(found, "epoch unbonding record found")
		for _, hzu := range epochUnbondingRecord.HostZoneUnbondings {
			// check that the status is CLAIMABLE
			if hzu.HostZoneId == HostChainId {
				s.Require().Equal(recordtypes.HostZoneUnbonding_CLAIMABLE, hzu.Status, "host zone unbonding status is CLAIMABLE")
				s.Require().Equal(hzu.ClaimableNativeTokens, tc.initialState.nativeTokenAmount, "claimable native tokens")
			}
		}
	}
}

func (s *KeeperTestSuite) checkRedemptionStateIfCallbackFailed(tc RedemptionCallbackTestCase) {
	initialState := tc.initialState
	for _, epochNumber := range initialState.epochUnbondingNumbers {
		// fetch the epoch unbonding record
		epochUnbondingRecord, found := s.App.RecordsKeeper.GetEpochUnbondingRecord(s.Ctx, epochNumber)
		s.Require().True(found, "epoch unbonding record found")
		for _, hzu := range epochUnbondingRecord.HostZoneUnbondings {
			// check that the status is NOT CLAIMABLE
			s.Require().Equal(recordtypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE, hzu.Status, "host zone unbonding status is NOT CLAIMABLE (EXIT_TRANSFER_QUEUE)")
		}
	}
}

func (s *KeeperTestSuite) TestRedemptionCallback_RedemptionCallbackTimeout() {
	tc := s.SetupRedemptionCallback()

	// Update the ack response to indicate a timeout
	invalidArgs := tc.validArgs
	invalidArgs.ackResponse.Status = icacallbacktypes.AckResponseStatus_TIMEOUT

	err := s.App.StakeibcKeeper.RedemptionCallback(s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidArgs.args)
	s.Require().NoError(err)
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_RedemptionCallbackErrorOnHost() {
	tc := s.SetupRedemptionCallback()

	// an error ack means the tx failed on the host
	invalidArgs := tc.validArgs
	invalidArgs.ackResponse.Status = icacallbacktypes.AckResponseStatus_FAILURE

	err := s.App.StakeibcKeeper.RedemptionCallback(s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidArgs.args)
	s.Require().NoError(err)
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_WrongCallbackArgs() {
	tc := s.SetupRedemptionCallback()
	invalidArgs := tc.validArgs

	// random args should cause the callback to fail
	invalidCallbackArgs := []byte("random bytes")

	err := s.App.StakeibcKeeper.RedemptionCallback(s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidCallbackArgs)
	s.Require().ErrorContains(err, "unable to unmarshal redemption callback args: unexpected EOF")
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_EpochUnbondingRecordNotFound() {
	tc := s.SetupRedemptionCallback()

	// Update the callbacks args to have a different epoch number
	// (So that epoch unbonding record is not found)
	callbackArgs := types.RedemptionCallback{
		HostZoneId:              HostChainId,
		EpochUnbondingRecordIds: []uint64{tc.initialState.epochNumber + 1},
	}
	invalidCallbackArgs, err := s.App.StakeibcKeeper.MarshalRedemptionCallbackArgs(s.Ctx, callbackArgs)
	s.Require().NoError(err)

	err = s.App.StakeibcKeeper.RedemptionCallback(s.Ctx, tc.validArgs.packet, tc.validArgs.ackResponse, invalidCallbackArgs)
	s.Require().ErrorContains(err, "host zone unbonding record not found on epoch unbonding record")
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_HostZoneUnbondingNotFound() {
	tc := s.SetupRedemptionCallback()

	// remove the hzu from the epoch unbonding record
	epochUnbondingRecord, found := s.App.RecordsKeeper.GetEpochUnbondingRecord(s.Ctx, tc.initialState.epochNumber)
	s.Require().True(found)
	epochUnbondingRecord.HostZoneUnbondings = []*recordtypes.HostZoneUnbonding{}
	s.App.RecordsKeeper.SetEpochUnbondingRecord(s.Ctx, epochUnbondingRecord)

	err := s.App.StakeibcKeeper.RedemptionCallback(s.Ctx, tc.validArgs.packet, tc.validArgs.ackResponse, tc.validArgs.args)
	s.Require().ErrorContains(err, "host zone unbonding record not found on epoch unbonding record")
	s.checkRedemptionStateIfCallbackFailed(tc)
}
