package models

import "time"

type Log struct {
	TimeStamp time.Time
	Action    string
	User      string
	ItemId    int
	Quantily  int
	Reason    string
}
