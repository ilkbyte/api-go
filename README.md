### **Ilkbyte Package**

This package developed for go package.

#### **Installation**

`go get github.com/ilkbyte/api-go`
    
#### **Usage**

```go
package main

import (
	"fmt"

	ilkbyte "github.com/ilkbyte/api-go"
)

func main() {
	c := ilkbyte.NewClient("access_key", "secret_key")

	fmt.Println(c.GetAllServers())
}
```

#### **Avalaible Methods**

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
