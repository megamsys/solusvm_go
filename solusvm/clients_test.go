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

import (
	"crypto/md5"
	"encoding/hex"
//	"fmt"
	"gopkg.in/check.v1"
)


func (s *S) TestClientsList(c *check.C) {
	//id and key is to be regitered as API access for your host ip
	clients, _, err := s.client.SolusClients.ListAllClients(map[string]string{ "id":""+ "iy9rRvifGKajunciPcu5V13ANyAmVnvklN2HV8cv" +"","key":""+ "8mQloZ1rjkl6bevOCW2o0mykZpSLnV8l8OwmCnEN" +"", "action": "client-list", "rdtype": "json"})
	c.Assert(err, check.IsNil)
  c.Assert(clients, check.NotNil)
}


func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
