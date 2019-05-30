package service

import (
	"errors"
	"fmt"

	pipelinesv1 "github.com/VideoCoin/cloud-api/pipelines/v1"
	workorderv1 "github.com/VideoCoin/cloud-api/workorder/v1"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
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

func (ds *Datastore) GetPipelineByStreamId(streamId uint64) (*pipelinesv1.Pipeline, error) {
	workOrder := &workorderv1.WorkOrder{}
	err := ds.Db.Where("stream_id = ?", streamId).First(workOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrPipelineNotFound
		}

		return nil, fmt.Errorf("failed to get work order by stream id %d: %s", streamId, err.Error())
	}

	pipeline := &pipelinesv1.Pipeline{}
	err = ds.Db.Where("id = ?", workOrder.PipelineId).First(pipeline).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrPipelineNotFound
		}

		return nil, fmt.Errorf("failed to get pipeline by id %s: %s", workOrder.PipelineId, err.Error())
	}

	return pipeline, nil
}
