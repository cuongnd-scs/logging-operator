// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/output"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
)

func TestSplunkHEC(t *testing.T) {
	CONFIG := []byte(`
hec_host: splunk.default.svc.cluster.local
hec_port: 8088
protocol: http

`)
	expected := `
  <match **>
	@type splunk_hec
	@id test_splunk_hec
	hec_host splunk.default.svc.cluster.local
	hec_port 8088
	protocol http
  </match>
`
	es := &output.SplunkHecOutput{}
	yaml.Unmarshal(CONFIG, es)
	test := render.NewOutputPluginTest(t, es)
	test.DiffResult(expected)
}