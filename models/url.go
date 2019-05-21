package models

import (
	"github.com/jinzhu/gorm"
)

type UrlState int

const (
	UrlStateIdle    UrlState = iota + 1 //1
	UrlStateRunning                     //2
)

type UrlStatus int

const (
	UrlStatusReady          UrlStatus = iota + 1 //1
	UrlStatusSuccess                             //2
	UrlStatusStopped                             //3
	UrlStatusError                               //4
	UrlStatusNotFoundParser                      //5
)

type Url struct {
	gorm.Model
	Url              string
	State            UrlState  //~enum UrlState
	Status           UrlStatus //~enum UrlStatus
	DownloadHttpCode int
}
