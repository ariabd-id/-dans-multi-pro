package services

import (
	"dans-multi-pro/constants"
	"dans-multi-pro/helpers"
	"dans-multi-pro/models"
	"dans-multi-pro/params"
	"dans-multi-pro/repositories"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type JobService struct {
	jobRepo repositories.JobRepo
}

func NewJobService(repo repositories.JobRepo) *JobService {
	return &JobService{
		jobRepo: repo,
	}
}

func (j *JobService) GetJobList(request params.GetJob) *params.Response {
	var responseData []models.Job
	apiUrl := constants.API_LIST

	resultAPI, err := helpers.FetchAPI(apiUrl)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	json.Unmarshal(resultAPI, &responseData)

	// filtering
	if request.Location != "" && request.Description == "" {
		var filteredResponse []models.Job

		for _, v := range responseData {
			if strings.EqualFold(strings.ToLower(request.Location), strings.ToLower(v.Location)) {
				filteredResponse = append(filteredResponse, v)
			}
		}

		responseData = filteredResponse
	}

	if request.Location == "" && request.Description != "" {
		var filteredResponse []models.Job

		for _, v := range responseData {
			if strings.Contains(strings.ToLower(v.Description), strings.ToLower(request.Description)) {
				filteredResponse = append(filteredResponse, v)
			}
		}

		responseData = filteredResponse
	}

	if request.Location != "" && request.Description != "" {
		var filteredResponse []models.Job

		for _, v := range responseData {
			if strings.EqualFold(strings.ToLower(request.Location), strings.ToLower(v.Location)) && strings.Contains(strings.ToLower(v.Description), strings.ToLower(request.Description)) {
				filteredResponse = append(filteredResponse, v)
			}
		}

		responseData = filteredResponse
	}

	// filtered full time
	if request.Fulltime == "true" {
		var filteredResponse []models.Job

		for _, v := range responseData {
			if strings.EqualFold(strings.ToLower(v.Type), strings.ToLower(constants.FULL_TIME)) {
				filteredResponse = append(filteredResponse, v)
			}
		}

		responseData = filteredResponse
	}

	// paginate
	if request.Page > 0 {
		maxIndex := (request.Page * constants.PER_PAGE) - 1
		minIndex := maxIndex - constants.PER_PAGE + 1

		log.Println(minIndex)
		log.Println(maxIndex)

		var pagedResponse []models.Job

		for i, v := range responseData {
			if i >= minIndex && i <= maxIndex {
				pagedResponse = append(pagedResponse, v)
			}
		}

		responseData = pagedResponse
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: responseData,
	}

}

func (j *JobService) GetJobDetail(id string) *params.Response {
	var responseData models.Job
	apiUrl := constants.API_DETAIL + "/" + id

	resultAPI, err := helpers.FetchAPI(apiUrl)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	json.Unmarshal(resultAPI, &responseData)

	return &params.Response{
		Status:  http.StatusOK,
		Payload: responseData,
	}
}
