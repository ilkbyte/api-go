package ilkbyte

import utils "github.com/umtaktpe/api-go"

// Server base struct.
type Server struct {
	QueryBuilder utils.QueryBuilder
}

// All servers.
func (c *Server) All() string {
	c.QueryBuilder.Endpoint = "/server/list/all"

	return utils.Request(&c.QueryBuilder)
}

// Active servers.
func (c *Server) Active() string {
	c.QueryBuilder.Endpoint = "/server/list"

	return utils.Request(&c.QueryBuilder)
}

// Config settings.
func (c *Server) Config() string {
	c.QueryBuilder.Endpoint = "/server/create"

	return utils.Request(&c.QueryBuilder)
}

// Create new server.
func (c *Server) Create(query map[string]string) string {
	c.QueryBuilder.Endpoint = "/server/create/config"

	data := make(map[string]string)
	for k, v := range query {
		data[k] = v
	}

	c.QueryBuilder.Query = data
	return utils.Request(&c.QueryBuilder)
}

// Show server details.
func (c *Server) Show(serverName string) string {
	c.QueryBuilder.Endpoint = "/server/manage/" + serverName + "/show"

	return utils.Request(&c.QueryBuilder)
}

// Power settings.
func (c *Server) Power(serverName string) string {
	c.QueryBuilder.Endpoint = "/server/manage/" + serverName + "/power"

	return utils.Request(&c.QueryBuilder)
}

// IP details.
func (c *Server) IP(serverName string) string {
	c.QueryBuilder.Endpoint = "/server/manage/" + serverName + "/ip/list"

	return utils.Request(&c.QueryBuilder)
}

// IPLogs data.
func (c *Server) IPLogs(serverName string) string {
	c.QueryBuilder.Endpoint = "/server/manage/" + serverName + "/ip/logs"

	return utils.Request(&c.QueryBuilder)
}

// IPRdns edit.
func (c *Server) IPRdns(serverName string, query map[string]string) string {
	c.QueryBuilder.Endpoint = "/server/manage/" + serverName + "/ip/rdns"

	data := make(map[string]string)
	for k, v := range query {
		data[k] = v
	}

	c.QueryBuilder.Query = data

	return utils.Request(&c.QueryBuilder)
}
