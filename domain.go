package ilkbyte

import (
	"strconv"

	"github.com/ilkbyte/api-go/models"
)

// GetDomains -- Code commentation required
func (c *Client) GetDomains(page int) (*models.Domain, error) {
	var (
		params = map[string]string{
			"p": strconv.Itoa(page),
		}
		res models.Domain
	)

	if err := c.getRequest(&res, "/domain/list", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateDomain -- Code commentation required
func (c *Client) CreateDomain(domain, serverName, ipv6 string) (*Response, error) {
	var (
		params = map[string]string{
			"domain":      domain,
			"server_name": serverName,
			"ipv6":        ipv6,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/domain/create", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetDNSDetails -- Code commentation required
func (c *Client) GetDNSDetails(domainName string) (*models.DNSDetails, error) {
	var res models.DNSDetails

	if err := c.getRequest(&res, "/domain/manage/"+domainName+"/show", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// DNSAdd -- Code commentation required
func (c *Client) DNSAdd(domainName, recordName, recordContent, recordPriority, recordType string) (*models.Record, error) {
	var (
		params = map[string]string{
			"record_name":     recordName,
			"record_content":  recordContent,
			"record_priority": recordPriority,
			"record_type":     recordType,
		}
		res models.Record
	)

	if err := c.getRequest(&res, "/domain/manage/"+domainName+"/add", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// DNSUpdate -- Code commentation required
func (c *Client) DNSUpdate(domainName, recordContent, recordID, recordPriority string) (*models.Record, error) {
	var (
		params = map[string]string{
			"record_content":  recordContent,
			"record_id":       recordID,
			"record_priority": recordPriority,
		}
		res models.Record
	)

	if err := c.getRequest(&res, "/domain/manage/"+domainName+"/update", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// DNSDelete -- Code commentation required
func (c *Client) DNSDelete(domainName, recordID string) (*Response, error) {
	var (
		params = map[string]string{
			"record_id": recordID,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/domain/manage/"+domainName+"/delete", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// DNSPush -- Code commentation required
func (c *Client) DNSPush(domainName string) (*Response, error) {
	res := Response{}

	if err := c.getRequest(&res, "/domain/manage/"+domainName+"/push", nil); err != nil {
		return nil, err
	}

	return &res, nil
}
