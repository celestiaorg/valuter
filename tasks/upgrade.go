package tasks

import (
	"github.com/celestiaorg/valuter/blocksigners"
	"github.com/celestiaorg/valuter/configs"
	"github.com/celestiaorg/valuter/participants"
	"github.com/celestiaorg/valuter/winners"
)

func GetNodeUpgradeWinners() (winners.WinnersList, error) {

	var winnersList winners.WinnersList

	if configs.Configs.Tasks.NodeUpgrade.MaxWinners == 0 {
		return winnersList, nil
	}

	listOfValidators, err := blocksigners.GetSignersByBlockHeight(configs.Configs.Tasks.NodeUpgrade.Condition.UpgradeHight)
	if err != nil {
		return winnersList, err
	}

	for i := range listOfValidators {

		valInfo, err := listOfValidators[i].GetValidatorInfo()
		if err != nil {
			return winnersList, err
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
			Rewards:         configs.Configs.Tasks.NodeUpgrade.Reward,
			ValidatorInfo:   valInfo,
			ParticipantData: pRecord,
		}

		winnersList.Append(newWinner)
		if winnersList.Length() >= configs.Configs.Tasks.NodeUpgrade.MaxWinners {
			break // Max winners reached
		}
	}

	return winnersList, nil
}
