package roothints

import "math/rand"

// RootServer defines name and addresses of a root server
type RootServer struct {
	Name string
	A    string
	AAAA string
}

// RootServers are the official root servers
var RootServers = []RootServer{
	{
		Name: "a.root-servers.net",
		A:    "198.41.0.4",
		AAAA: "2001:503:ba3e::2:30",
	},
	{
		Name: "b.root-servers.net",
		A:    "192.228.79.201",
		AAAA: "2001:500:84::b",
	},
	{
		Name: "c.root-servers.net",
		A:    "192.33.4.12",
		AAAA: "2001:500:2::c",
	},
	{
		Name: "d.root-servers.net",
		A:    "199.7.91.13",
		AAAA: "2001:500:2d::d",
	},
	{
		Name: "e.root-servers.net",
		A:    "192.203.230.10",
	},
	{
		Name: "f.root-servers.net",
		A:    "192.5.5.241",
		AAAA: "2001:500:2f::f",
	},
	{
		Name: "g.root-servers.net",
		A:    "192.112.36.4",
	},
	{
		Name: "h.root-servers.net",
		A:    "198.97.190.53",
		AAAA: "2001:500:1::53",
	},
	{
		Name: "i.root-servers.net",
		A:    "192.36.148.17",
		AAAA: "2001:7fe::53",
	},
	{
		Name: "j.root-servers.net",
		A:    "192.58.128.30",
		AAAA: "2001:503:c27::2:30",
	},
	{
		Name: "k.root-servers.net",
		A:    "193.0.14.129",
		AAAA: "2001:7fd::1",
	},
	{
		Name: "l.root-servers.net",
		A:    "199.7.83.42",
		AAAA: "2001:500:9f::42",
	},
	{
		Name: "m.root-servers.net",
		A:    "202.12.27.33",
		AAAA: "2001:dc3::35",
	},
}

func Random() RootServer {
	return RootServers[rand.Intn(len(RootServers))]
}

func RandomA() RootServer {
	for {
		server := Random()
		if server.A != "" {
			return server
		}
	}
}

func RandomAAAA() RootServer {
	for {
		server := Random()
		if server.AAAA != "" {
			return server
		}
	}
}
