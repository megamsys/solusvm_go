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
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"gopkg.in/check.v1"
)

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %s, want %s", header, got, want)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}

func (s *S) TestDoAction(c *check.C) {
  cl := NewClient(nil, s.url)
	req, _ := cl.NewWRequest(map[string]string{"id": ""+"iy9rRvifGKajunciPcu5V13ANyAmVnvklN2HV8cv"+"", "key": ""+ "8mQloZ1rjkl6bevOCW2o0mykZpSLnV8l8OwmCnEN"+"", "nodeid": "4","rdtype": "json",}, "node-virtualservers")
	//body := new(foo)
	body := new(VServers)
	s.client.Do(*req, &body)
	//c.Assert(body, check.DeepEquals,body)
	c.Assert(req,check.NotNil)
}
