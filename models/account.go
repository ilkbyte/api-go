package models

type Account struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Balance   float64 `json:"balance"`
		CreatedAt string  `json:"created_at"`
		Group     string  `json:"group"`
		Name      string  `json:"name"`
		Status    bool    `json:"status"`
	} `json:"data"`
}

type AccountUsers struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
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
	} `json:"data"`
}
