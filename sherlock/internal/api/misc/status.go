package misc

import (
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const pburl = `https://pact-broker.dsp-eng-tools.broadinstitute.org`
const alphaID = `19e0aeee-4fc5-49a7-b502-065e3fbce967`
const productionID = `354e0bfa-9634-417c-b10a-4beea2ffc3bd`

type RecordDeploymentResponse struct {
	currentlyDeployed bool `json:"currentlyDeployed"`
}
type StatusResponse struct {
	OK bool `json:"ok"`
}

// statusGet godoc
//
//	@summary		Get Sherlock's current status
//	@description	Get Sherlock's current status. Right now, this endpoint always returned OK (if the server is online).
//	@description	This endpoint is acceptable to use for a readiness check.
//	@tags			Misc
//	@produce		json
//	@success		200	{object}	misc.StatusResponse
//	@router			/status [get]
func statusGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StatusResponse{OK: true})
}

type PactResponse struct {
	OK bool `json:"ok"`
}

//	type Participant struct {
//		ParticipantList []struct {
//			Name          string `json:"name"`
//			DisplayName   string `json:"displayName"`
//			RepositoryUrl string `json:"repositoryUrl"`
//		} `json:"environment"`
//	}
type ParticipantResponse struct {
	Deployed bool `json:"currentlyDeployed"`
}

// PactHandler godoc
//
//	@summary		Get Sherlock's current status
//	@description	Get Sherlock's current status. Right now, this endpoint always returned OK (if the server is online).
//	@description	This endpoint is acceptable to use for a readiness check.
//	@tags			Misc
//	@produce		json
//	@success		200	{object}	misc.StatusResponse
//	@router			/pact [get]
func PactHandler(ctx *gin.Context) {

	// config.Config.Bool("pactbroker.enabled")
	//var embeddedResponse ParticipantResponse
	//request, error := http.NewRequest(http.MethodGet, pburl+"/pacticipants", nil)
	//request.Header.Set("Content-Type", "application/json; charset=utf-8")
	//request.Header.Set("Accept", "application/hal+json")
	//request.SetBasicAuth(config.Config.MustString("pactbroker.auth.username"), config.Config.MustString("pactbroker.auth.password"))
	//// send the request
	//client := &http.Client{}
	//response, error := client.Do(request)
	//
	//if error != nil {
	//	fmt.Println(error)
	//}
	//
	//responseBody, error := io.ReadAll(response.Body)
	//defer response.Body.Close()
	//error = json.Unmarshal(responseBody, &embeddedResponse)
	//
	//if error != nil {
	//	fmt.Println(error)
	//}
	//for _, participant := range embeddedResponse.Participant.ParticipantList {
	//	if participant.RepositoryUrl == "https://github.com/DataBiosphere/terra-billing-profile-manager" {
	//		participantName := participant.Name
	//	}
	//}
	var eID string
	var chartName = "bpm-consumer"
	var rdp ParticipantResponse
	var chartverion = "beda2f1bf9282a11a674b6d1311104c788659c6a"
	var environmentName = "prod"
	switch en := environmentName; en {
	case "alpha":
		eID = alphaID
	case "prod":
		eID = productionID
	}
	request, err := http.NewRequest(http.MethodPost, pburl+"/pacticipants/"+chartName+
		"/versions/"+chartverion+"/deployed-versions/environment"+eID, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/hal+json")
	request.SetBasicAuth(config.Config.MustString("pactbroker.auth.username"), config.Config.MustString("pactbroker.auth.password"))
	// send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	responseBody, err := io.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, &rdp)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(response.StatusCode, rdp)
}

// function to format JSON data
//func formatJSON(data []byte) string {
//	var out bytes.Buffer
//	err := json.Indent(&out, data, "", " ")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	d := out.Bytes()
//	return string(d)
//}
//var eID string
//var rdp RecordDeploymentResponse
//var chartverion = "beda2f1bf9282a11a674b6d1311104c788659c6a"
//prod := "prod"
//switch en := prod; en {
//case "alpha":
//eID = alphaID
//case "prod":
//eID = productionID
//}
//request, err := http.NewRequest(http.MethodPost, pburl+"/pacticipants/bpm-consumer"+
//"/versions/"+chartverion+"/deployed-versions/environment"+eID, nil)
//if err != nil {
//return err
//}
//request.Header.Set("Content-Type", "application/json; charset=utf-8")
//request.Header.Set("Content-Type", "application/json; charset=utf-8")
//request.Header.Set("Accept", "application/hal+json")
//request.SetBasicAuth(config.Config.MustString("pactbroker.auth.username"), config.Config.MustString("pactbroker.auth.password"))
//// send the request
//client := &http.Client{}
//response, err := client.Do(request)
//if err != nil {
//return err
//}
//responseBody, err := io.ReadAll(response.Body)
//err = json.Unmarshal(responseBody, &rdp)
//if err != nil {
//return err
//}
//if err != nil {
//return err
//}
//if !rdp.currentlyDeployed || (response.StatusCode != 201) {
//return errors.New("deployment was not recorded successfully")
//}
//return nil
