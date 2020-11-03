### **Ilkbyte Package**

This package developed for go package.

#### **Installation**

`go get github.com/ilkbyte/api-go`


#### **Configuration**

You should add below parameters in .env file.

```
ILKBYTE_ACCESSKEY=''
ILKBYTE_SECRETKEY=''
```
    
#### **Usage**

```go
package main

import (
	"fmt"

	ilkbyte "github.com/umtaktpe/api-go/src"
)

func main() {
	server := &ilkbyte.Server{}

	fmt.Println(server.All())
}
```

#### **Avalaible Methods**

##### **Server**

```go
server := &ilkbyte.Server{}
// get all servers
server.All()
// get only active servers
server.Active()
// Get server configs you can choose
server.Config()
// create a new server
server.Create(map[string]string{
    "username": "username",
    "password": "password",
    "name": "name",
    "os_id": "osID",
    "app_id": "appID",
    "package_id": "packageID",
    "sshkey": "sshKey",
})
// Get server details
server.Show("server-name")
// Server power settings
server.Power("server-name", "status-value")
// Get all ips from server
server.IP("server-name")
// Get ip logs
server.IPLogs("server-name")
// Add a new rdns record
server.IPRdns("server-name", map[string]string{
    "ip": "ip-address",
    "rdns": "rdns-name",
})
```

##### **Domain**

```go
domain := &ilkbyte.Domain{}
// Get all domains
domain.All()
// Create a new domain
domain.Create(map[string]string{
    "domain": "domain-name",
})
// Get domain details
domain.Show("domain-name")
// Add a new record
domain.Add("domain-name", map[string]string{
    "record_name": "recordName",
    "record_type": "recordType",
    "record_content": "recordContent",
    "record_priority": "recordPriority"
})
// Update an existing record
domain.Update("domain-name", map[string]string{
    "record_id": "recordID",
    "record_content": "recordContent",
    "record_priority": "recordPriority",
})
// Delete domain
domain.Delete("domain-name")
// Push changes to server
domain.Push("domain-name")
```
