package ilkbyte

import (
	"net/http"
	"strconv"

	"github.com/ilkbyte/api-go/models"
)

func (c *Client) GetDomains(page int) (*models.Domain, error) {
	params := map[string]string{
		"p": strconv.Itoa(page),
	}
	req, err := http.NewRequest("GET", c.BaseURL+"/domain/list", nil)
	if err != nil {
		return nil, err
	}

	var res models.Domain
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateDomain(domain, serverName, ipv6 string) (*Response, error) {
	params := make(map[string]string)
	params["domain"] = domain
	params["server_name"] = serverName
	params["ipv6"] = ipv6

	req, err := http.NewRequest("GET", c.BaseURL+"/domain/create", nil)

	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetDNSDetails(domainName string) (*models.DNSDetails, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/show", nil)
	if err != nil {
		return nil, err
	}

	var res models.DNSDetails
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DNSAdd(domainName, recordName, recordContent, recordPriority, recordType string) (*models.Record, error) {
	params := make(map[string]string)
	params["record_name"] = recordName
	params["record_content"] = recordContent
	params["record_priority"] = recordPriority
	params["record_type"] = recordType

	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/add", nil)

	if err != nil {
		return nil, err
	}

	var res models.Record
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DNSUpdate(domainName, recordContent, recordId, recordPriority string) (*models.Record, error) {
	params := make(map[string]string)
	params["record_content"] = recordContent
	params["record_id"] = recordId
	params["record_priority"] = recordPriority

	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/update", nil)

	if err != nil {
		return nil, err
	}

	var res models.Record
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DNSDelete(domainName, recordId string) (*Response, error) {
	params := make(map[string]string)
	params["record_id"] = recordId

	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/delete", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DNSPush(domainName string) (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/push", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}
