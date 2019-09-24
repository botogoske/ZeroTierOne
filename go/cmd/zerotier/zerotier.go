/*
 * Copyright (c)2019 ZeroTier, Inc.
 *
 * Use of this software is governed by the Business Source License included
 * in the LICENSE.TXT file in the project's root directory.
 *
 * Change Date: 2023-01-01
 *
 * On the date above, in accordance with the Business Source License, use
 * of this software will be governed by version 2.0 of the Apache License.
 */
/****/

package main

import (
	"flag"
	"fmt"
	"os"
	"zerotier/cmd/zerotier/cli"
	"zerotier/pkg/zerotier"
)

var copyrightText = fmt.Sprintf(`ZeroTier Network Virtualization Service Version %d.%d.%d
(c)2019 ZeroTier, Inc.
Licensed under the ZeroTier BSL (see LICENSE.txt)`,
	zerotier.CoreVersionMajor, zerotier.CoreVersionMinor, zerotier.CoreVersionRevision)

func printHelp() {
	fmt.Println(copyrightText + `

Usage: zerotier [-options] <command> [-options] [command args]

Global Options
  -j                                   Output raw JSON where applicable
  -p <path>                            Connect to service running at this path
  -t <authtoken.secret path>           Use secret auth token from this file

Commands:
  help                                 Show this help
  version                              Print version
  service [path]                       Start in system service mode
  status                               Show ZeroTier service status and config
  peers                                Show VL1 peers
  roots                                Show VL1 root servers
  addroot <type> [options]             Add a VL1 root
    static <identity> <ip/port> [...]  Add a root with a set identity and IPs
    dynamic <name> [default locator]   Add a dynamic root fetched by name
  removeroot <type> [options]          Remove a VL1 root
    static <identity>                  Remove a root with a set identity
    dynamic <name>                     Remove a dynamic root fetched by name
  networks                             Show joined VL2 virtual networks
  join <network ID>                    Join a virtual network
  leave <network ID>                   Leave a virtual network
  show <network ID>                    Show verbose network info
  set <network ID> <option> <value>    Set a network local config option
    manageips <boolean>                Is IP management allowed?
    manageroutes <boolean>             Is route management allowed?
    globalips <boolean>                Can IPs in global IP space be managed?
    globalroutes <boolean>             Can global IP space routes be set?
    defaultroute <boolean>             Can default route be overridden?
  set <local config option> <value>    Set a local configuration option
    phy <IP/bits> blacklist <boolean>  Set or clear blacklist for CIDR
    phy <IP/bits> trust <path ID/0>    Set or clear trusted path ID for CIDR
    virt <address> try <IP/port> [...] Set explicit IPs for reaching a peer
    port <port>                        Set primary local port for VL1 P2P
    secondaryport <port/0>             Set or disable secondary VL1 P2P port
    tertiaryport <port/0>              Set or disable tertiary VL1 P2P port
    portsearch <boolean>               Set or disable port search on startup
    portmapping <boolean>              Set or disable use of uPnP and NAT-PMP
    explicitaddresses <IP/port> [...]  Set explicit external IPs to advertise

Most commands require a secret token to permit control of a running ZeroTier
service. The CLI will automatically try to read this token from the
authtoken.secret file in the service's working directory and then from a
file called .zerotierauth in the user's home directory. The -t option can be
used to explicitly specify a location.
`)
}

/*
func nodeStart() {
	osSignalChannel := make(chan os.Signal, 2)
	signal.Notify(osSignalChannel, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGBUS)
	signal.Ignore(syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		<-osSignalChannel
	}()
}
*/

func main() {
	globalOpts := flag.NewFlagSet("global", flag.ContinueOnError)
	hflag := globalOpts.Bool("h", false, "") // support -h to be canonical with other Unix utilities
	err := globalOpts.Parse(os.Args[1:])
	if err != nil {
		printHelp()
		os.Exit(1)
		return
	}
	args := globalOpts.Args()
	if len(args) < 1 || *hflag {
		printHelp()
		os.Exit(0)
		return
	}
	var cmdArgs []string
	if len(args) > 1 {
		cmdArgs = args[1:]
	}

	switch args[0] {
	case "help":
		printHelp()
		os.Exit(0)
		return
	default:
		printHelp()
		os.Exit(1)
		return

	case "version":
		fmt.Printf("%d.%d.%d\n", zerotier.CoreVersionMajor, zerotier.CoreVersionMinor, zerotier.CoreVersionRevision)
		os.Exit(0)

	case "service":
		cli.Service(cmdArgs)
	case "status":
		cli.Status(cmdArgs)
	case "peers":
	case "roots":
	case "addroot":
	case "removeroot":
	case "networks":
	case "join":
	case "leave":
	case "show":
	case "set":
	}
}