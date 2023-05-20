package controllers

import (
	"dans-multi-pro/params"
	"dans-multi-pro/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	jobService services.JobService
}

func NewJobController(service *services.JobService) *JobController {
	return &JobController{
		jobService: *service,
	}
}

func (j *JobController) GetJobList(c *gin.Context) {
	var req params.GetJob

	req.Description = c.Query("description")
	req.Location = c.Query("location")
	req.Fulltime = c.Query("full_time")
	req.Page, _ = strconv.Atoi(c.Query("page"))

	result := j.jobService.GetJobList(req)

	c.JSON(result.Status, result.Payload)
}

func (j *JobController) GetJobDetail(c *gin.Context) {
	detailId := c.Param("detailID")
	result := j.jobService.GetJobDetail(detailId)

	c.JSON(result.Status, result.Payload)
}
