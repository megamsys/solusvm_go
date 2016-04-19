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
type VServersService struct {
	client *Client
}

// VirtualServer represents an Solusvm user.
type VirtualServer struct {
	Vserverid   	*string `json:"vserverid"`
	Ctid_xid    	*string `json:"ctid-xid"`
	Clientid      *string `json:"clientid"`
	Ipaddress     *string `json:"ipaddress"`
	Hostname      *string `json:"hostname"`
	Template      *string `json:"template"`
	Hdd    				*string `json:"hdd"`
	Memory     		*string `json:"memory"`
	Swap_burst 		*string `json:"swap-burst"`
	Type    			*string `json:"type"`
	Mac      			*string `json:"mac"`
}

type VServers struct {
	Status     *string `json:"status"`
	Statusmsg  *string `json:"statusmsg"`
	VirtualServers *[]VirtualServer `json:"virtualservers"`
}

func (u VServersService) String() string {
	return Stringify(u)
}

// Solusvm API docs: http://docs.Solusvm.com/API:Get_VirtualServers_Details
func (s *VServersService) ListAll(parms map[string]string) (*VServers, *Response, error) {
	a := new(VServers)
	resp, err := do(s.client, Params{parms: parms, u: "node-virtualservers"}, a)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return nil, resp, err
	}
	return a, resp, err
}
