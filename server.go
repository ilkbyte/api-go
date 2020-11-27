package ilkbyte

import (
	"net/http"
	"time"
)

type Server struct {
	Pagination struct {
		CurrentPage int `json:"CurrentPage"`
		TotalPage   int `json:"TotalPage"`
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
}

type ServerConfig struct {
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
}

type ServerCreate struct {
	ServerInfo struct {
		IPV4     string `json:"ipv4"`
		IPV6     string `json:"ipv6"`
		Name     string `json:"name"`
		Osapp    string `json:"osapp"`
		Password string `json:"password"`
		Service  string `json:"service"`
		Username string `json:"username"`
	} `json:"server_info"`
}

type ServerDetails struct {
	Service        string `json:"service"`
	Status         string `json:"status"`
	IPV4           string `json:"ipv4"`
	IPV6           string `json:"ipv6"`
	BandwidthLimit int64  `json:"bandwidth_limit"`
	BandwidthUsage int64  `json:"bandwidth_usage"`
}

type ServerPower struct {
	Status string `json:"status"`
}

type ServerIPList struct {
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
}

type ServerIPLogs []struct {
	IPPrefix string `json:"ip_prefix"`
	IsLog    bool   `json:"is_log"`
	IsPerson bool   `json:"is_person"`
	LogFile  string `json:"log_file"`
	RuleIn   string `json:"rule_in"`
	RuleOut  string `json:"rule_out"`
	RuleType string `json:"rule_type"`
}

type Snapshot struct {
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
}

type Backup struct {
	Amount string `json:"amount"`
	Backup struct {
		Name       string    `json:"name"`
		FileSize   string    `json:"file_size"`
		FileHash   string    `json:"file_hash"`
		IsLocked   bool      `json:"is_locked"`
		BackupTime time.Time `json:"backup_time"`
	} `json:"backup"`
}

type BackupRestoreResp struct {
	Name       string    `json:"name"`
	FileSize   string    `json:"file_size"`
	FileHash   string    `json:"file_hash"`
	IsLocked   bool      `json:"is_locked"`
	BackupTime time.Time `json:"backup_time"`
}

func (c *Client) GetAllServers() (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/list/all", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Server{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetActiveServers() (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/list", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Server{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerConfig() (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/create", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &ServerConfig{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateServer(name, username, password, osid, appid, packageid, sshkey string) (*Response, error) {
	params := make(map[string]string)
	params["name"] = name
	params["username"] = username
	params["password"] = password
	params["os_id"] = osid
	params["app_id"] = appid
	params["package_id"] = packageid
	params["sshkey"] = sshkey

	req, err := http.NewRequest("GET", c.BaseURL+"/server/create/config", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &ServerCreate{},
	}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerDetails(serverName string) (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/show", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &ServerDetails{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) PowerServer(serverName, set string) (*Response, error) {
	params := make(map[string]string)
	params["set"] = set

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/power", nil)

	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &ServerPower{},
	}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerIPList(serverName string) (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/ip/list", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &ServerIPList{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerIPLogs(serverName string) (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/ip/logs", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ServerIPRdns(serverName, ip, rdns string) (*Response, error) {
	params := make(map[string]string)
	params["ip"] = ip
	params["rdns"] = rdns

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/ip/rdns", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAllSnapshots(serverName string) (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Snapshot{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateSnapshot(serverName, snapshotName string) (*Response, error) {
	params := make(map[string]string)
	params["name"] = snapshotName

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot/create", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) RevertSnapshot(serverName, snapshotName string) (*Response, error) {
	params := make(map[string]string)
	params["name"] = snapshotName

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot/revert", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateSnapshot(serverName, snapshotName string) (*Response, error) {
	params := make(map[string]string)
	params["name"] = snapshotName

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot/update", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteSnapshot(serverName, snapshotName string) (*Response, error) {
	params := make(map[string]string)
	params["name"] = snapshotName

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot/delete", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) AddCronSnapshot(serverName, snapshotName, day, hour, min string) (*Response, error) {
	params := make(map[string]string)
	params["name"] = snapshotName
	params["day"] = day
	params["hour"] = hour
	params["min"] = min

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot/cron/add", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCronSnapshot(serverName, snapshotName string) (*Response, error) {
	params := make(map[string]string)
	params["name"] = snapshotName

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot/cron/delete", nil)
	if err != nil {
		return nil, err
	}

	res := Response{}
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAllBackup(serverName string) (*Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/backup", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &Backup{},
	}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) BackupRestore(serverName, backupName string) (*Response, error) {
	params := make(map[string]string)
	params["backup_name"] = backupName

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/backup/restore", nil)
	if err != nil {
		return nil, err
	}

	res := Response{
		Data: &BackupRestoreResp{},
	}

	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}
