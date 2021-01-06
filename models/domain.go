package models

type Domain struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
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
	} `json:"data"`
}

type DNSDetails struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Records []struct {
			Content string `json:"content"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Prio    int    `json:"int"`
			TTL     int    `json:"ttl"`
			Type    string `json:"type"`
		} `json:"records"`
	} `json:"data"`
}

type Record struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Record struct {
			Name    string `json:"name"`
			Type    string `json:"type"`
			Content string `json:"content"`
			TTL     int    `json:"ttl"`
			Prio    int    `json:"prio"`
		} `json:"record"`
	} `json:"data"`
}
