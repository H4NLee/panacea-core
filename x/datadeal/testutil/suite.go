package testutil

import (
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/medibloc/panacea-core/v2/types/assets"
	"github.com/medibloc/panacea-core/v2/types/testsuite"
	"github.com/medibloc/panacea-core/v2/x/datadeal/types"
	oracletypes "github.com/medibloc/panacea-core/v2/x/oracle/types"
)

type DataDealBaseTestSuite struct {
	testsuite.TestSuite
}

func (suite *DataDealBaseTestSuite) MakeTestDeal(dealID uint64, buyerAddr sdk.AccAddress, maxNumData uint64) *types.Deal {
	return &types.Deal{
		Id:           dealID,
		Address:      types.NewDealAddress(dealID).String(),
		DataSchema:   []string{"http://jsonld.com"},
		Budget:       &sdk.Coin{Denom: assets.MicroMedDenom, Amount: sdk.NewInt(1000000000)},
		MaxNumData:   maxNumData,
		CurNumData:   0,
		BuyerAddress: buyerAddr.String(),
		Status:       types.DEAL_STATUS_ACTIVE,
	}
}

func (suite *DataDealBaseTestSuite) MakeNewDataSale(sellerAddr sdk.AccAddress, dataHash, verifiableCID string) *types.DataSale {
	return &types.DataSale{
		SellerAddress: sellerAddr.String(),
		DealId:        1,
		VerifiableCid: verifiableCID,
		DeliveredCid:  "",
		DataHash:      dataHash,
		Status:        types.DATA_SALE_STATUS_VERIFICATION_VOTING_PERIOD,
		VerificationVotingPeriod: &oracletypes.VotingPeriod{
			VotingStartTime: time.Now(),
			VotingEndTime:   time.Now().Add(5 * time.Second),
		},
		DeliveryVotingPeriod:    nil,
		VerificationTallyResult: nil,
		DeliveryTallyResult:     nil,
	}
}

func (suite *DataDealBaseTestSuite) SetValidator(pubKey cryptotypes.PubKey, amount sdk.Int, commission sdk.Dec) stakingtypes.Validator {
	varAddr := sdk.ValAddress(pubKey.Address().Bytes())
	validator, err := stakingtypes.NewValidator(varAddr, pubKey, stakingtypes.Description{})
	suite.Require().NoError(err)
	validator = validator.UpdateStatus(stakingtypes.Bonded)
	validator, _ = validator.AddTokensFromDel(amount)
	newCommission := stakingtypes.NewCommission(commission, sdk.OneDec(), sdk.NewDecWithPrec(5, 1))
	validator.Commission = newCommission
	validator.MinSelfDelegation = amount

	suite.StakingKeeper.SetValidator(suite.Ctx, validator)
	err = suite.StakingKeeper.SetValidatorByConsAddr(suite.Ctx, validator)
	suite.Require().NoError(err)

	return validator
}

func (suite *DataDealBaseTestSuite) MakeNewDataVerificationVote(voterAddr sdk.AccAddress, dataHash string) *types.DataVerificationVote {
	return &types.DataVerificationVote{
		VoterAddress: voterAddr.String(),
		DealId:       1,
		DataHash:     dataHash,
		VoteOption:   oracletypes.VOTE_OPTION_YES,
	}
}

func (suite *DataDealBaseTestSuite) MakeNewDataSaleDeliveryVoting(sellerAddr sdk.AccAddress, dataHash, verifiableCID string) *types.DataSale {
	return &types.DataSale{
		SellerAddress:            sellerAddr.String(),
		DealId:                   1,
		VerifiableCid:            verifiableCID,
		DeliveredCid:             "",
		DataHash:                 dataHash,
		Status:                   types.DATA_SALE_STATUS_DELIVERY_VOTING_PERIOD,
		VerificationVotingPeriod: nil,
		DeliveryVotingPeriod: &oracletypes.VotingPeriod{
			VotingStartTime: time.Now(),
			VotingEndTime:   time.Now().Add(5 * time.Second),
		},
		VerificationTallyResult: nil,
		DeliveryTallyResult:     nil,
	}
}

func (suite *DataDealBaseTestSuite) MakeNewDataSaleDeliveryFailed(sellerAddr sdk.AccAddress, dataHash, verifiableCID string) *types.DataSale {
	return &types.DataSale{
		SellerAddress:            sellerAddr.String(),
		DealId:                   1,
		VerifiableCid:            verifiableCID,
		DeliveredCid:             "",
		DataHash:                 dataHash,
		Status:                   types.DATA_SALE_STATUS_DELIVERY_FAILED,
		VerificationVotingPeriod: nil,
		DeliveryVotingPeriod: &oracletypes.VotingPeriod{
			VotingStartTime: time.Now(),
			VotingEndTime:   time.Now().Add(5 * time.Second),
		},
		VerificationTallyResult: nil,
		DeliveryTallyResult:     nil,
	}
}

func (suite *DataDealBaseTestSuite) MakeNewDataDeliveryVote(voterAddr sdk.AccAddress, dataHash, deliveredCID string, dealID uint64) *types.DataDeliveryVote {
	return &types.DataDeliveryVote{
		VoterAddress: voterAddr.String(),
		DealId:       dealID,
		DataHash:     dataHash,
		DeliveredCid: deliveredCID,
		VoteOption:   oracletypes.VOTE_OPTION_YES,
	}
}

func (suite *DataDealBaseTestSuite) CreateOracleValidator(pubKey cryptotypes.PubKey, amount sdk.Int) stakingtypes.Validator {
	suite.SetAccount(pubKey)

	val1Commission := sdk.NewDecWithPrec(1, 1)

	val := suite.SetValidator(pubKey, amount, val1Commission)

	oracleAccAddr := sdk.AccAddress(pubKey.Address().Bytes())
	oracle := &oracletypes.Oracle{
		Address:  oracleAccAddr.String(),
		Status:   oracletypes.ORACLE_STATUS_ACTIVE,
		Uptime:   0,
		JailedAt: nil,
	}

	suite.Require().NoError(suite.OracleKeeper.SetOracle(suite.Ctx, oracle))

	return val
}

func (suite *DataDealBaseTestSuite) SetAccount(pubKey cryptotypes.PubKey) {
	oracleAccAddr := sdk.AccAddress(pubKey.Address().Bytes())
	oracleAccount := suite.AccountKeeper.NewAccountWithAddress(suite.Ctx, oracleAccAddr)
	suite.Require().NoError(oracleAccount.SetPubKey(pubKey))
	suite.AccountKeeper.SetAccount(suite.Ctx, oracleAccount)
}

func (suite *DataDealBaseTestSuite) SetupValidatorRewards(valAddress sdk.ValAddress) {
	decCoins := sdk.DecCoins{sdk.NewDecCoinFromDec(assets.MicroMedDenom, sdk.ZeroDec())}
	historicalRewards := distrtypes.NewValidatorHistoricalRewards(decCoins, 1)
	suite.DistrKeeper.SetValidatorHistoricalRewards(suite.Ctx, valAddress, 1, historicalRewards)
	// setup current rewards
	currentRewards := distrtypes.NewValidatorCurrentRewards(decCoins, 2)
	suite.DistrKeeper.SetValidatorCurrentRewards(suite.Ctx, valAddress, currentRewards)
}
