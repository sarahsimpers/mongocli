// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mongodb/mongocli/internal/cli"

	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "metrics"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use, "measurements", "measurement"),
		Short:   "Get measurements on the state of the MongoDB process.",
	}
	cmd.AddCommand(
		ProcessBuilder(),
		DisksBuilder(),
		DatabasesBuilder(),
	)

	return cmd
}

// getHostnameAndPort return the hostname and the port starting from the string hostname:port.
func getHostnameAndPort(hostInfo string) (hostname string, port int, err error) {
	const hostnameParts = 2
	host := strings.Split(hostInfo, ":")
	if len(host) != hostnameParts {
		return "", 0, fmt.Errorf("expected hostname:port, got %s", host)
	}

	port, err = strconv.Atoi(host[1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid port number, got %s", host[1])
	}

	return host[0], port, nil
}
