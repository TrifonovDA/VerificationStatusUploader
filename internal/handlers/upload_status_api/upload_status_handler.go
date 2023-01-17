package upload_status_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deliryo-io/VerificationStatusUpdater/pkg"
	"github.com/deliryo-io/VerificationStatusUpdater/pkg/database_tools"
	"github.com/jackc/pgx/v5/pgconn"
	"io/ioutil"
	"log"
	"net/http"
)

type Sumsub_request struct {
	ApplicantId    string `json:"applicantId,omitempty"`
	InspectionId   string `json:"inspectionId,omitempty"`
	ApplicantType  string `json:"applicantType,omitempty"`
	CorrelationId  string `json:"correlationId,omitempty"`
	LevelName      string `json:"levelName,omitempty"`
	SandboxMode    bool   `json:"sandboxMode,omitempty"`
	ExternalUserId string `json:"externalUserId,omitempty"`
	Type           string `json:"type,omitempty"`
	ReviewResult   struct {
		ReviewAnswer string `json:"reviewAnswer"`
	} `json:"reviewResult,omitempty"`
	ReviewStatus string `json:"reviewStatus,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	CreatedAtMs  string `json:"createdAtMs,omitempty"`
	ClientId     string `json:"clientId,omitempty"`
}

var DbConnection = database_tools.NewConnection(context.Background())

const query_GREEN = "UPDATE postgredb.production.user_info set verification_status = 3, update_dttm = now() where user_id = $1;"
const query_RED = "UPDATE postgredb.production.user_info set verification_status = 2, update_dttm = now() where user_id = $1;"

func Upload_status_handler(w http.ResponseWriter, r *http.Request) {
	pkg.EnableCors(&w)
	var request Sumsub_request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body_reading_err: %v\n", err)
	} else {
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Printf("Unmarshaling_err: %v\n", err)
		}
	}
	defer r.Body.Close()

	log.Printf("Id: %v, result %v", request.ExternalUserId, request.ReviewResult.ReviewAnswer)

	if err != nil {
		log.Printf("Error")
	}
	if request.ReviewResult.ReviewAnswer == "GREEN" {
		_, err1 := DbConnection.Query(context.Background(), query_GREEN, request.ExternalUserId)
		if err1 != nil {
			log.Printf("There's error: %v", err1)
			if pgErr, ok := err1.(*pgconn.PgError); ok { //обработка ошибок бд
				newErr := fmt.Sprintf("SQL Error: %s, Detail: %s, Code: %s, SQLState: %%", pgErr.Message, pgErr.Detail, pgErr.Code, pgErr.SQLState())
				fmt.Println(newErr)
				log.Fatal(newErr)
			}
		}
	} else {
		_, err1 := DbConnection.Query(context.Background(), query_RED, request.ExternalUserId)
		if err1 != nil {
			log.Printf("There's error: %v", err1)
			if pgErr, ok := err1.(*pgconn.PgError); ok { //обработка ошибок бд
				newErr := fmt.Sprintf("SQL Error: %s, Detail: %s, Code: %s, SQLState: %%", pgErr.Message, pgErr.Detail, pgErr.Code, pgErr.SQLState())
				fmt.Println(newErr)
				log.Fatal(newErr)
			}
		}
	}
}
