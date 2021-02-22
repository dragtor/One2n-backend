package controller

import (
	"log"

	"github.com/dragtor/One2n-backend/backend/pkg"
)

type Controller struct {
	S3Iterator *pkg.AwsS3Iterator
}

func NewController(s3Iterator *pkg.AwsS3Iterator) *Controller {
	return &Controller{
		S3Iterator: s3Iterator,
	}
}

type ControllerResponse struct {
	LsDir []string
}

func convertResponseToControllerStructure(lsdir []string) *ControllerResponse {
	return &ControllerResponse{
		LsDir: lsdir,
	}
}

func (c *Controller) CommandS3ls(path string) (*ControllerResponse, error) {
	// Perform logic to fetch s3 ls command
	lsdir, err := c.S3Iterator.ListDir(path)
	if err != nil {
		log.Printf("Error : Failed to List dir")
		return nil, err
	}
	ctrlResponse := convertResponseToControllerStructure(lsdir)
	return ctrlResponse, nil
}
