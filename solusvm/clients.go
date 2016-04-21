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
type ClientsService struct {
	client *Client
}

// VirtualServer represents an Solusvm user.
type SolusClient struct {
	Id   	          *string `json:"id"`
	Username    	  *string `json:"username"`
	Email           *string `json:"email"`
	FirstName       *string `json:"firstname"`
	LastName        *string `json:"lastname"`
	Company    			*string `json:"company"`
	Level     		  *string `json:"level"`
	Status 		      *string `json:"status"`
	Created    			*string `json:"created"`
	Lastlogin  			*string `json:"lastlogin"`
}

type SClients struct {
	Status     *string `json:"status"`
	Statusmsg  *string `json:"statusmsg"`
	SolusClients *[]SolusClient `json:"clients"`
}

func (u ClientsService) String() string {
	return Stringify(u)
}

// Solusvm API docs: http://docs.Solusvm.com/API:Get_VirtualServers_Details
func (s *ClientsService) ListAllClients(parms map[string]string) (*SClients, *Response, error) {
	a := new(SClients)
	resp, err := do(s.client, Params{parms: parms, u: "client-list"}, a)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return nil, resp, err
	}
	return a, resp, err
}
