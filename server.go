package ilkbyte

import (
	"strconv"

	"github.com/ilkbyte/api-go/models"
)

// GetAllServers -- Code commentation required
func (c *Client) GetAllServers(page int) (*models.Server, error) {
	var (
		params = map[string]string{
			"p": strconv.Itoa(page),
		}
		res models.Server
	)

	if err := c.getRequest(&res, "/server/list/all", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetActiveServers -- Code commentation required
func (c *Client) GetActiveServers(page int) (*models.Server, error) {
	var (
		params = map[string]string{
			"p": strconv.Itoa(page),
		}
		res models.Server
	)

	if err := c.getRequest(&res, "/server/list", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetServerConfig -- Code commentation required
func (c *Client) GetServerConfig() (*models.ServerConfig, error) {
	var res models.ServerConfig

	if err := c.getRequest(&res, "/server/create", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateServer -- Code commentation required
func (c *Client) CreateServer(name, username, password, osid, appid, packageid, sshkey string) (*models.ServerCreate, error) {
	var (
		params = map[string]string{
			"name":       name,
			"username":   username,
			"password":   password,
			"os_id":      osid,
			"app_id":     appid,
			"package_id": packageid,
			"sshkey":     sshkey,
		}
		res models.ServerCreate
	)

	if err := c.getRequest(&res, "/server/create/config", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetServerDetails -- Code commentation required
func (c *Client) GetServerDetails(serverName string) (*models.ServerDetails, error) {
	var res models.ServerDetails

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/show", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// PowerServer -- Code commentation required
func (c *Client) PowerServer(serverName, set string) (*models.ServerPower, error) {
	var (
		params = map[string]string{
			"set": set,
		}
		res models.ServerPower
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/power", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetServerIPList -- Code commentation required
func (c *Client) GetServerIPList(serverName string) (*models.ServerIPList, error) {
	var res models.ServerIPList

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/ip/list", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetServerIPLogs -- Code commentation required
func (c *Client) GetServerIPLogs(serverName string) (*models.ServerIPLogs, error) {
	var res models.ServerIPLogs

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/ip/logs", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// ServerIPRdns -- Code commentation required
func (c *Client) ServerIPRdns(serverName, ip, rdns string) (*Response, error) {
	var (
		params = map[string]string{
			"ip":   ip,
			"rdns": rdns,
		}
		res Response
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/ip/rdns", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetAllSnapshots -- Code commentation required
func (c *Client) GetAllSnapshots(serverName string) (*models.Snapshot, error) {
	var res models.Snapshot

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/snapshot", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateSnapshot -- Code commentation required
func (c *Client) CreateSnapshot(serverName, snapshotName string) (*Response, error) {
	var (
		params = map[string]string{
			"name": snapshotName,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/snapshot/create", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// RevertSnapshot -- Code commentation required
func (c *Client) RevertSnapshot(serverName, snapshotName string) (*Response, error) {
	var (
		params = map[string]string{
			"name": snapshotName,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/snapshot/revert", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// UpdateSnapshot -- Code commentation required
func (c *Client) UpdateSnapshot(serverName, snapshotName string) (*Response, error) {
	var (
		params = map[string]string{
			"name": snapshotName,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/snapshot/update", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteSnapshot -- Code commentation required
func (c *Client) DeleteSnapshot(serverName, snapshotName string) (*Response, error) {
	var (
		params = map[string]string{
			"name": snapshotName,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/snapshot/delete", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// AddCronSnapshot -- Code commentation required
func (c *Client) AddCronSnapshot(serverName, snapshotName, day, hour, min string) (*Response, error) {
	var (
		params = map[string]string{
			"name": snapshotName,
			"day":  day,
			"hour": hour,
			"min":  min,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/snapshot/cron/add", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteCronSnapshot -- Code commentation required
func (c *Client) DeleteCronSnapshot(serverName, snapshotName string) (*Response, error) {
	var (
		params = map[string]string{
			"name": snapshotName,
		}
		res = Response{}
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/snapshot/cron/delete", params); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetAllBackup -- Code commentation required
func (c *Client) GetAllBackup(serverName string) (*models.Backup, error) {
	var res models.Backup

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/backup", nil); err != nil {
		return nil, err
	}

	return &res, nil
}

// BackupRestore -- Code commentation required
func (c *Client) BackupRestore(serverName, backupName string) (*models.BackupRestoreResp, error) {
	var (
		params = map[string]string{
			"backup_name": backupName,
		}
		res models.BackupRestoreResp
	)

	if err := c.getRequest(&res, "/server/manage/"+serverName+"/backup/restore", params); err != nil {
		return nil, err
	}

	return &res, nil
}
