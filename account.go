package ilkbyte

import (
	"net/http"
	"strconv"

	"github.com/ilkbyte/api-go/models"
)

func (c *Client) GetAccountInfo() (*models.Account, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/account", nil)
	if err != nil {
		return nil, err
	}

	var res models.Account
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAccountUsers(page int) (*models.AccountUsers, error) {
	params := map[string]string{
		"p": strconv.Itoa(page),
	}

	req, err := http.NewRequest("GET", c.BaseURL+"/account/users", nil)
	if err != nil {
		return nil, err
	}

	var res models.AccountUsers
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}
