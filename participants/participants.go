package participants

import (
	"fmt"

	"github.com/celestiaorg/cosmologger/database"
	"github.com/celestiaorg/valuter/tools"
	"github.com/celestiaorg/valuter/types"
)

type ParticipantRecord struct {
	FullLegalName  string `json:"full_legal_name"`
	GithubHandle   string `json:"github_handle"`
	EmailAddress   string `json:"email_address"`
	AccountAddress string `json:"account_address"`
	PubKey         string `json:"pub_key"` //TODO: Check if we need it anymore
	Country        string `json:"country"`
	KycSessionId   string `json:"kyc_session_id"`
	KycVerified    bool   `json:"kyc_verified"`
}

// This function receives a json string of the signed ID,
// verifies it with the given signature and if it passes,
// the data will be added to the database
func ImportBySignature(jsonStr string) error {

	// container, err := getAgSignerContainer(jsonStr)
	// if err != nil {
	// 	return err
	// }

	// // The input string was empty
	// if container == nil {
	// 	return nil
	// }

	// verified, err := container.VerifySubmission()

	// if err != nil {
	// 	return err
	// }
	// // The data is not verified
	// if !verified {
	// 	return fmt.Errorf("the data is not verified")
	// }

	// Let's add it to the database
	return AddNew(ParticipantRecord{
		// ID:           container.ID,
		KycSessionId: "",
		KycVerified:  false,
	})
}

func AddNew(participant ParticipantRecord) error {

	// Check if the record is already in the db
	queryRes, err := database.DB.Load(database.TABLE_PARTICIPANTS,
		database.RowType{
			database.FIELD_PARTICIPANTS_EMAIL_ADDRESS:   participant.EmailAddress,
			database.FIELD_PARTICIPANTS_ACCOUNT_ADDRESS: participant.AccountAddress,
		})
	if err != nil {
		return err
	}

	// Already exist, let's update it, a user might correct their signature in the next submissions
	if len(queryRes) > 0 && participant.AccountAddress != "" {
		_, err := participant.Update()
		return err
	}
	_, err = database.DB.Insert(database.TABLE_PARTICIPANTS, participant.getDBRow())
	return err
}

func GetParticipants() ([]ParticipantRecord, error) {

	rows, err := database.DB.Load(database.TABLE_PARTICIPANTS, nil)
	if err != nil {
		return nil, err
	}

	return DBRowToParticipantRecords(rows), err
}

func GetParticipantsWithPagination(limitOffset types.DBLimitOffset) ([]ParticipantRecord, types.Pagination, error) {

	// Prepare pagination
	totalRows := uint64(0)
	{
		SQL := fmt.Sprintf(`SELECT COUNT(*) AS "total" FROM "%s"`,
			database.TABLE_PARTICIPANTS,
		)
		rows, err := database.DB.Query(SQL, database.QueryParams{})
		if err != nil {
			return nil, types.Pagination{}, err
		}
		totalRows = uint64(rows[0]["total"].(int64))
	}
	pagination := tools.GetPagination(totalRows, limitOffset.Page)

	/*------*/

	SQL := fmt.Sprintf(`SELECT * FROM "%s" LIMIT $1 OFFSET $2`, database.TABLE_PARTICIPANTS)

	rows, err := database.DB.Query(SQL,
		database.QueryParams{
			limitOffset.Limit,
			limitOffset.Offset,
		})
	if err != nil {
		return nil, types.Pagination{}, err
	}

	return DBRowToParticipantRecords(rows), pagination, err
}

func GetParticipantByAddress(accAddress string) (ParticipantRecord, error) {

	rows, err := database.DB.Load(database.TABLE_PARTICIPANTS,
		database.RowType{
			database.FIELD_PARTICIPANTS_ACCOUNT_ADDRESS: accAddress,
		})

	if err != nil || rows == nil || len(rows) == 0 {
		return ParticipantRecord{}, err
	}

	return DBRowToParticipantRecord(rows[0]), err
}

func GetParticipantsByEmail(email string) ([]ParticipantRecord, error) {

	rows, err := database.DB.Load(database.TABLE_PARTICIPANTS,
		database.RowType{
			database.FIELD_PARTICIPANTS_EMAIL_ADDRESS: email,
		})

	if err != nil || rows == nil || len(rows) == 0 {
		return []ParticipantRecord{}, err
	}

	return DBRowToParticipantRecords(rows), err
}

// Returns RowsAffected, error
func (p *ParticipantRecord) UpdateKYC() (int, error) {

	if p.EmailAddress == "" {
		return 0, fmt.Errorf("email address cannot be empty")
	}

	uRes, err := database.DB.Update(
		database.TABLE_PARTICIPANTS,
		database.RowType{ // Fields to update
			database.FIELD_PARTICIPANTS_KYC_SESSION_ID: p.KycSessionId,
			database.FIELD_PARTICIPANTS_KYC_VERIFIED:   p.KycVerified,
		},
		database.RowType{ // Conditions
			database.FIELD_PARTICIPANTS_EMAIL_ADDRESS: p.EmailAddress,
		},
	)
	return int(uRes.RowsAffected), err
}

// Returns RowsAffected, error
func (p *ParticipantRecord) Update() (int, error) {

	if p.EmailAddress == "" {
		return 0, fmt.Errorf("email address cannot be empty")
	}

	uRes, err := database.DB.Update(
		database.TABLE_PARTICIPANTS,
		p.getDBRow(), // Fields to update
		database.RowType{ // Conditions
			database.FIELD_PARTICIPANTS_EMAIL_ADDRESS:   p.EmailAddress,
			database.FIELD_PARTICIPANTS_ACCOUNT_ADDRESS: p.AccountAddress,
		},
	)
	return int(uRes.RowsAffected), err
}

// Returns RowsAffected, error
func (p *ParticipantRecord) UpdateByEmail() (int, error) {

	if p.EmailAddress == "" {
		return 0, fmt.Errorf("email address cannot be empty")
	}

	uRes, err := database.DB.Update(
		database.TABLE_PARTICIPANTS,
		database.RowType{ // Conditions
			database.FIELD_PARTICIPANTS_FULL_LEGAL_NAME: p.FullLegalName,
			database.FIELD_PARTICIPANTS_COUNTRY:         p.Country,
		}, // Fields to update
		database.RowType{ // Conditions
			database.FIELD_PARTICIPANTS_EMAIL_ADDRESS: p.EmailAddress,
		},
	)
	return int(uRes.RowsAffected), err
}

func ImportByEmail(email string, fullName string, country string) error {

	participants, err := GetParticipantsByEmail(email)
	if err != nil {
		return err
	}

	if len(participants) == 0 { // Not found in the DB
		return AddNew(ParticipantRecord{
			EmailAddress:  email,
			FullLegalName: fullName,
			Country:       country,
		})
	}

	p := participants[0]

	// Let's update the found record
	if p.FullLegalName == "" {
		p.FullLegalName = fullName
	}
	if p.Country == "" {
		p.Country = country
	}

	_, err = p.UpdateByEmail() // Update the country name for all instances
	return err
}
