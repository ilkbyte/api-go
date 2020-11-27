package ilkbyte

import "net/http"

type Account struct {
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
	Group     string  `json:"group"`
	Name      string  `json:"name"`
	Status    bool    `json:"status"`
}

type AccountUsers struct {
	Pagination struct {
		CurrentPage int `json:"current_page"`
		TotalPage   int `json:"total_page"`
	} `json:"pagination"`
	UserList []struct {
		Email     string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		LastLogin string `json:"last_login"`
		Level     string `json:"level"`
		Status    string `json:"status"`
		Verify    bool   `json:"verify"`
	} `json:"user_list"`
}

func (c *Client) GetAccountInfo() (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/account", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Account{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAccountUsers() (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/account/users", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &AccountUsers{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}
