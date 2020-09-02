package main

import (
	"context"
	"errors"

	"github.com/digitalocean/godo"
)

type (
	// Domain represents a DNS domain from Digital Ocean.
	Domain struct {
		*godo.Domain
		ZoneName ZoneName
		Records  []*Record
	}
)

var (
	// ErrZoneNotFound will be returned if the zone cannot be found at Digital Ocean.
	ErrZoneNotFound = errors.New("Zone not found")
)

// NewDomain will instantiate a new Domain from the specified zoneName.
func NewDomain(zoneName ZoneName) *Domain {
	return &Domain{
		ZoneName: zoneName,
	}
}

// Find will try to find the matching domain at Digital Ocean.
func (d *Domain) Find(client *godo.Client) error {
	// Check if the domain is registered at Digital Ocean.
	domains, _, err := client.Domains.List(context.TODO(), nil)
	if err != nil {
		return err
	}

	for _, domain := range domains {
		if d.ZoneName == NewZoneName(domain.Name) {
			d.Domain = &domain
			return nil
		}
	}

	return ErrZoneNotFound
}
