package ilkbyte

import (
	utils "github.com/umtaktpe/api-go"
)

// Domain base struct.
type Domain struct {
	QueryBuilder utils.QueryBuilder
}

// All domains.
func (c *Domain) All() string {
	c.QueryBuilder.Endpoint = "/domain/list"

	return utils.Request(&c.QueryBuilder)
}

// Create new domain.
func (c *Domain) Create(query map[string]string) string {
	c.QueryBuilder.Endpoint = "/domain/create"

	data := make(map[string]string)
	for k, v := range query {
		data[k] = v
	}

	c.QueryBuilder.Query = data

	return utils.Request(&c.QueryBuilder)
}

// Show domain details.
func (c *Domain) Show(domainName string) string {
	c.QueryBuilder.Endpoint = "/domain/manage/" + domainName + "/show"

	return utils.Request(&c.QueryBuilder)
}

// Add new domain.
func (c *Domain) Add(domainName string, query map[string]string) string {
	c.QueryBuilder.Endpoint = "/domain/manage/" + domainName + "/add"

	data := make(map[string]string)
	for k, v := range query {
		data[k] = v
	}

	c.QueryBuilder.Query = data

	return utils.Request(&c.QueryBuilder)
}

// Update domain data.
func (c *Domain) Update(domainName string, query map[string]string) string {
	c.QueryBuilder.Endpoint = "/domain/manage/" + domainName + "/update"

	data := make(map[string]string)
	for k, v := range query {
		data[k] = v
	}

	c.QueryBuilder.Query = data

	return utils.Request(&c.QueryBuilder)
}

// Delete domain.
func (c *Domain) Delete() string {
	c.QueryBuilder.Endpoint = "/domain/delete"

	return utils.Request(&c.QueryBuilder)
}

// Push changed data.
func (c *Domain) Push(domainName string) string {
	c.QueryBuilder.Endpoint = "/domain/manage/" + domainName + "/push"

	return utils.Request(&c.QueryBuilder)
}
