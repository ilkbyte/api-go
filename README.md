### **Ilkbyte Package**

This package developed for go package.

#### **Installation**

`go get github.com/ilkbyte/api-go`
    
#### **Usage**

```go
package main

import (
	"fmt"

	ilkbyte "github.com/umtaktpe/api-go"
)

func main() {
	c := ilkbyte.NewClient("access_key", "secret_key")

	fmt.Println(c.GetAllServers())
}
```

#### **Avalaible Methods**

##### **Account**
```go
c := ilkbyte.NewClient("access_key", "secret_key")
// get your account's info.
c.GetAccountInfo()
// get your account users.
c.GetAccountUsers()
```
##### **Server**

```go
c := ilkbyte.NewClient("access_key", "secret_key")
// get all servers
c.GetAllServers()
// get only active servers
c.GetActiveServers()
// Get server configs you can choose
c.GetServerConfig()
// create a new server
c.CreateServer("name", "username", "password", "osid", "appid", "packageid", "sshkey")
// Get server details
c.GetServerDetails("servername")
// Server power settings
c.PowerServer("servername", "status")
// Get all ips from server
c.GetServerIPList("servername")
// Get ip logs
c.GetServerIPLogs("servername")
// Add a new rdns record
c.ServerIPRdns("servername", "ip", "rdns")
```

##### **Backup**
```go
c := ilkbyte.NewClient("access_key", "secret_key")
// Get backup list.
c.GetAllBackup("servername")
// Restore your backup.
c.BackupRestore("servername", "backupname")
```

##### **Snapshot**
```go
c := ilkbyte.NewClient("access_key", "secret_key")
// Get all snapshots.
c.GetAllSnapshots("servername")
// Create a new snapshot.
c.CreateSnapshot("servername", "snapshotname")
// Revert snapshot.
c.RevertSnapshot("servername", "snapshotname")
// Recreate your snapshot.
c.UpdateSnapshot("servername", "snapshotname")
// Delete snapshot.
c.DeleteSnapshot("servername", "snapshotname")
// Add cron to your snapshot.
c.AddCronSnapshot("servername", "snapshotname", "day", "hour", "minute")
// Delete cron.
c.DeleteCronSnapshot("snapshotname", "snapshotname")
```

##### **Domain**

```go
c := ilkbyte.NewClient("access_key", "secret_key")
// Get all domains
c.GetDomains()
// Create a new domain
c.CreateDomain("domain", "server_name", "ipv6")
// Get domain details
c.GetDNSDetails("domainname")
// Add a new record
c.DNSAdd("domainname", "record_name", "record_content", "record_priority", "record_type")
// Update an existing record
c.DNSUpdate("domainname", "record_content", "record_id", "record_priority")
// Delete domain
c.DNSDelete("domainname", "record_id")
// Push changes to server
c.DNSPush("domainname")
```
