package models

import "time"

type Server struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Pagination struct {
			CurrentPage int `json:"current_page"`
			TotalPage   int `json:"total_page"`
		} `json:"pagination"`
		ServerList []struct {
			BandwidthLimit int64  `json:"bandwidth_limit"`
			BandwidthUsage int    `json:"bandwidth_usage"`
			DeletedTime    string `json:"deleted_time,omitempty"`
			Ipv4           string `json:"ipv4"`
			Ipv6           string `json:"ipv6"`
			Name           string `json:"name"`
			Osapp          string `json:"osapp"`
			Service        string `json:"service"`
			Status         string `json:"status"`
			CreatedTime    string `json:"created_time,omitempty"`
		} `json:"server_list"`
	} `json:"data"`
}

type ServerConfig struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Application []struct {
			Code   int    `json:"code"`
			Name   string `json:"name"`
			System string `json:"system"`
		} `json:"application"`
		OperatingSystem []struct {
			Code    int    `json:"code"`
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"operating_system"`
		Package []struct {
			Code     int    `json:"code"`
			Features string `json:"features"`
			Name     string `json:"name"`
			Price    string `json:"price"`
		} `json:"package"`
	} `json:"data"`
}

type ServerCreate struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		ServerInfo struct {
			IPV4     string `json:"ipv4"`
			IPV6     string `json:"ipv6"`
			Name     string `json:"name"`
			Osapp    string `json:"osapp"`
			Password string `json:"password"`
			Service  string `json:"service"`
			Username string `json:"username"`
		} `json:"server_info"`
	} `json:"data"`
}

type ServerDetails struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Service        string `json:"service"`
		Status         string `json:"status"`
		IPV4           string `json:"ipv4"`
		IPV6           string `json:"ipv6"`
		BandwidthLimit int64  `json:"bandwidth_limit"`
		BandwidthUsage int64  `json:"bandwidth_usage"`
	} `json:"data"`
}

type ServerPower struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Status string `json:"status"`
	} `json:"data"`
}

type ServerIPList struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		IPV4 []struct {
			Address string `json:"address"`
			Reverse string `json:"reverse"`
			ACLList string `json:"acl_list"`
			Usage   string `json:"usage"`
		} `json:"ipv4"`
		IPV6 []struct {
			Address string `json:"address"`
			Reverse string `json:"reverse"`
			ACLList string `json:"acl_list"`
			Usage   string `json:"usage"`
		} `json:"ipv6"`
	} `json:"data"`
}

type ServerIPLogs struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    []struct {
		IPPrefix string    `json:"ip_prefix"`
		IsLog    bool      `json:"is_log"`
		IsPerson bool      `json:"is_person"`
		LogFile  string    `json:"log_file"`
		RuleIn   time.Time `json:"rule_in"`
		RuleOut  time.Time `json:"rule_out"`
		RuleType string    `json:"rule_type"`
	} `json:"data"`
}

type Snapshot struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Snapshots []struct {
			Name        string `json:"name"`
			Current     bool   `json:"current"`
			State       string `json:"state"`
			Location    string `json:"location"`
			Parent      string `json:"parent"`
			Children    int    `json:"children"`
			Descendants int    `json:"descendants"`
			Metadata    bool   `json:"metadata"`
			Params      struct {
				IsCron   bool      `json:"is_cron"`
				LastRun  time.Time `json:"last_run"`
				NextRun  time.Time `json:"next_run"`
				CronTime string    `json:"cron_time"`
			} `json:"params"`
			Date time.Time `json:"date"`
		} `json:"snapshots"`
	} `json:"data"`
}

type Backup struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Amount string `json:"amount"`
		Backup struct {
			Name       string    `json:"name"`
			FileSize   string    `json:"file_size"`
			FileHash   string    `json:"file_hash"`
			IsLocked   bool      `json:"is_locked"`
			BackupTime time.Time `json:"backup_time"`
		} `json:"backup"`
	} `json:"data"`
}

type BackupRestoreResp struct {
	Status  bool   `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Name       string    `json:"name"`
		FileSize   string    `json:"file_size"`
		FileHash   string    `json:"file_hash"`
		IsLocked   bool      `json:"is_locked"`
		BackupTime time.Time `json:"backup_time"`
	} `json:"data"`
}
