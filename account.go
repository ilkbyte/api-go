package ilkbyte

import (
	"strconv"

	"github.com/ilkbyte/api-go/models"
)

// GetAccountInfo -- Code commentation required
func (c *Client) GetAccountInfo() (*models.Account, error) {
	var res models.Account

	if err := c.getRequest(&res, "/account", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetAccountUsers -- Code commentation required
func (c *Client) GetAccountUsers(page int) (*models.AccountUsers, error) {
	var (
		params = map[string]string{
			"p": strconv.Itoa(page),
		}
		res models.AccountUsers
	)

	if err := c.getRequest(&res, "/account/users", params); err != nil {
		return nil, err
	}

	return &res, nil
}
