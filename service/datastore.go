package service

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	jobsv1 "github.com/videocoin/cloud-api/jobs/v1"
	pipelinesv1 "github.com/videocoin/cloud-api/pipelines/v1"
)

var (
	ErrJobNotFound      = errors.New("job is not found")
	ErrPipelineNotFound = errors.New("pipeline is not found")
)

type Datastore struct {
	Db *gorm.DB
}

func NewDatastore(uri string) (*Datastore, error) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &Datastore{
		Db: db,
	}, err
}

func (ds *Datastore) GetJobByStreamId(streamId uint64) (*jobsv1.Job, error) {
	job := &jobsv1.Job{}
	err := ds.Db.Where("stream_id = ?", streamId).First(job).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrJobNotFound
		}

		return nil, fmt.Errorf("failed to get job by stream id %d: %s", streamId, err.Error())
	}

	return job, nil
}

func (ds *Datastore) GetPipelineById(pipelineId string) (*pipelinesv1.Pipeline, error) {
	pipeline := &pipelinesv1.Pipeline{}
	err := ds.Db.Where("id = ?", pipelineId).First(pipeline).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrPipelineNotFound
		}

		return nil, fmt.Errorf("failed to get pipeline by id %s: %s", pipelineId, err.Error())
	}

	return pipeline, nil
}
