package main

import "github.com/digitalocean/godo"

type (
	// Record represents a domain record from Digital Ocean.
	Record struct {
		godo.DomainRecord
		zoneName ZoneName
		Matched  bool
	}
)

// NewRecord will instantiate a new Record from a DomainRecord and a zoneName.
func NewRecord(record godo.DomainRecord, zoneName ZoneName) *Record {
	return &Record{
		DomainRecord: record,
		zoneName:     zoneName,
	}
}
