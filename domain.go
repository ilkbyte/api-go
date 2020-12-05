package ilkbyte

import (
	"net/http"
)

type Domain struct {
	Pagination struct {
		CurrentPage int `json:"current_page"`
		TotalPage   int `json:"total_page"`
	} `json:"pagination"`
	DomainList []struct {
		Domain     string `json:"domain"`
		Nameserver struct {
			Master string `json:"master"`
			Slave  string `json:"slave"`
		} `json:"ipv4"`
	} `json:"domain_list"`
}

type DNSDetails struct {
	Records []struct {
		Content string `json:"content"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Prio    int    `json:"int"`
		TTL     int    `json:"ttl"`
		Type    string `json:"type"`
	} `json:"records"`
}

type Record struct {
	Record struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Content string `json:"content"`
		TTL     int    `json:"ttl"`
		Prio    int    `json:"prio"`
	} `json:"record"`
}

func (c *Client) GetDomains() (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/domain/list", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Domain{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
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

func (c *Client) GetDNSDetails(domainName string) (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/show", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &DNSDetails{},
	}

	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DNSAdd(domainName, recordName, recordContent, recordPriority, recordType string) (*Response, error) {
	params := make(map[string]string)
	params["record_name"] = recordName
	params["record_content"] = recordContent
	params["record_priority"] = recordPriority
	params["record_type"] = recordType

	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/add", nil)

	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Record{},
	}

	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DNSUpdate(domainName, recordContent, recordId, recordPriority string) (*Response, error) {
	params := make(map[string]string)
	params["record_content"] = recordContent
	params["record_id"] = recordId
	params["record_priority"] = recordPriority

	req, err := http.NewRequest("GET", c.BaseURL+"/domain/manage/"+domainName+"/update", nil)

	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Record{},
	}

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
