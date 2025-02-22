package tasks

import (
	"github.com/celestiaorg/valuter/blocksigners"
	"github.com/celestiaorg/valuter/configs"
	"github.com/celestiaorg/valuter/participants"
	"github.com/celestiaorg/valuter/validators"
	"github.com/celestiaorg/valuter/winners"
)

func GetGenesisValidatorsWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	if configs.Configs.Tasks.ValidatorGenesis.MaxWinners == 0 {
		return winnersList, nil
	}

	// Those who signged the first block are considered as genesis validators
	// Since some joins might not be able to make it to the first block we change it to a higher block like 20
	listOfValidators, err := blocksigners.GetSignersByBlockHeight(20)
	if err != nil {
		return winnersList, err
	}

	for i := range listOfValidators {

		valInfo, err := listOfValidators[i].GetValidatorInfo()
		if err != nil {
			return winnersList, err
		}
		if valInfo.UpTime < configs.Configs.Tasks.ValidatorGenesis.UptimePercent {
			// Let's just ignore this validator
			continue
		}

		pRecord, err := participants.GetParticipantByAddress(listOfValidators[i].AccAddr)
		if err != nil {
			return winnersList, err
		}

		// If the participant is not verified by KYC provider, just ignore it
		if !pRecord.KycVerified {
			continue
		}

		newWinner := winners.Winner{
			Address:         listOfValidators[i].AccAddr,
			Rewards:         configs.Configs.Tasks.ValidatorGenesis.Reward,
			ValidatorInfo:   valInfo,
			ParticipantData: pRecord,
		}

		winnersList.Append(newWinner)
		if winnersList.Length() >= configs.Configs.Tasks.ValidatorGenesis.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}

func GetJoinedAfterGenesisValidatorsWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	if configs.Configs.Tasks.ValidatorJoin.MaxWinners == 0 {
		return winnersList, nil
	}

	listOfValidators, err := validators.GetJoinedAfterGenesisValidators()
	if err != nil {
		return winnersList, err
	}

	for i := range listOfValidators {

		valInfo, err := listOfValidators[i].GetValidatorInfo()
		if err != nil {
			return winnersList, err
		}
		if valInfo.UpTime < configs.Configs.Tasks.ValidatorJoin.UptimePercent {
			// Let's just ignore this validator
			continue
		}

		pRecord, err := participants.GetParticipantByAddress(listOfValidators[i].AccAddr)
		if err != nil {
			return winnersList, err
		}

		// If the participant is not verified by KYC provider, just ignore it
		if !pRecord.KycVerified {
			continue
		}

		newWinner := winners.Winner{
			Address:         listOfValidators[i].AccAddr,
			Rewards:         configs.Configs.Tasks.ValidatorJoin.Reward,
			ValidatorInfo:   valInfo,
			ParticipantData: pRecord,
		}

		winnersList.Append(newWinner)
		if winnersList.Length() >= configs.Configs.Tasks.ValidatorJoin.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}
