/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package solusvm

import "fmt"
// AccountsService handles communication with the user related
// methods of the Solusvm API.
//
// Solusvm API docs: http://docs.Solusvm.com/API#Client_Management
type NodesService struct {
	client *Client
}

type Nodes struct {
	Status     *string `json:"status"`
	Statusmsg  *string `json:"statusmsg"`
	Nodes 		 *string `json:"nodes"`
}

type NodeInfo struct {
Status      *string `json:"status"`
Statusmsg   *string `json:"statusmsg"`
NodeId 	 	  *string `json:"id"`
HypType 	  *string `json:"virt"`
NodeName 		*string `json:"name"`
NodeIp 			*string `json:"ip"`
DomainName 	*string `json:"hostname"`
TotVServers *int `json:"virtualservers"`
}

func (u NodesService) String() string {
	return Stringify(u)
}

func (s *NodesService) ListNodes(parms map[string]string) (*Nodes, *Response, error) {
	a := new(Nodes)
	resp, err := do(s.client, Params{parms: parms, u: "node-idlist"}, a)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return nil, resp, err
	}
	return a, resp, err
}

func (s *NodesService) NodeInfo(parms map[string]string) (*NodeInfo, *Response, error) {
	a := new(NodeInfo)
	resp, err := do(s.client, Params{parms: parms, u: "node-statistics"}, a)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return nil, resp, err
	}
	return a, resp, err
}
