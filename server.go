package ilkbyte

import (
	"net/http"
	"strconv"

	"github.com/ilkbyte/api-go/models"
)

func (c *Client) GetAllServers(page int) (*models.Server, error) {
	params := map[string]string{
		"p": strconv.Itoa(page),
	}
	req, err := http.NewRequest("GET", c.BaseURL+"/server/list/all", nil)
	if err != nil {
		return nil, err
	}

	var res models.Server
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetActiveServers(page int) (*models.Server, error) {
	params := map[string]string{
		"p": strconv.Itoa(page),
	}

	req, err := http.NewRequest("GET", c.BaseURL+"/server/list", nil)
	if err != nil {
		return nil, err
	}

	var res models.Server
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerConfig() (*models.ServerConfig, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/create", nil)
	if err != nil {
		return nil, err
	}

	var res models.ServerConfig
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateServer(name, username, password, osid, appid, packageid, sshkey string) (*models.ServerCreate, error) {
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

	var res models.ServerCreate
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerDetails(serverName string) (*models.ServerDetails, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/show", nil)
	if err != nil {
		return nil, err
	}

	var res models.ServerDetails
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) PowerServer(serverName, set string) (*models.ServerPower, error) {
	params := make(map[string]string)
	params["set"] = set

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/power", nil)

	if err != nil {
		return nil, err
	}

	var res models.ServerPower
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerIPList(serverName string) (*models.ServerIPList, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/ip/list", nil)
	if err != nil {
		return nil, err
	}

	var res models.ServerIPList
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetServerIPLogs(serverName string) (*models.ServerIPLogs, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/ip/logs", nil)
	if err != nil {
		return nil, err
	}

	var res models.ServerIPLogs
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

	var res Response
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAllSnapshots(serverName string) (*models.Snapshot, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/snapshot", nil)
	if err != nil {
		return nil, err
	}

	var res models.Snapshot
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

func (c *Client) GetAllBackup(serverName string) (*models.Backup, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/backup", nil)
	if err != nil {
		return nil, err
	}

	var res models.Backup
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) BackupRestore(serverName, backupName string) (*models.BackupRestoreResp, error) {
	params := make(map[string]string)
	params["backup_name"] = backupName

	req, err := http.NewRequest("GET", c.BaseURL+"/server/manage/"+serverName+"/backup/restore", nil)
	if err != nil {
		return nil, err
	}

	var res models.BackupRestoreResp
	if err := c.sendRequest(req, &res, params); err != nil {
		return nil, err
	}

	return &res, nil
}
