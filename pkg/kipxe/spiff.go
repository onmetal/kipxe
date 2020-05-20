/*
 * Copyright 2020 Mandelsoft. All rights reserved.
 *  This file is licensed under the Apache Software License, v. 2 except as noted
 *  otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package kipxe

import (
	"fmt"

	"github.com/gardener/controller-manager-library/pkg/types/infodata/simple"
	"github.com/mandelsoft/spiff/flow"
	"github.com/mandelsoft/spiff/yaml"
)

type SpiffTemplate struct {
	mapping yaml.Node
}

func (this *SpiffTemplate) AddStub(inp *[]yaml.Node, name string, v simple.Values) error {
	if v == nil {
		return nil
	}

	i, err := yaml.Sanitize(name, v)
	if err != nil {
		return fmt.Errorf("%s: invalid values: %s", name, err)
	}
	*inp = append(*inp, i)
	return nil
}

func (this *SpiffTemplate) MergeWith(inputs ...yaml.Node) (simple.Values, error) {
	stubs, err := flow.PrepareStubs(nil, false, inputs...)
	if err != nil {
		return nil, err
	}
	result, err := flow.Apply(nil, this.mapping, stubs)
	if err != nil {
		return nil, err
	}
	v, err := yaml.Normalize(result)
	if err != nil {
		return nil, err
	}

	m := v.(map[string]interface{})
	if out, ok := m["output"]; ok {
		if x, ok := out.(map[string]interface{}); ok {
			return simple.Values(x), nil
		}
		return nil, fmt.Errorf("unexpected type for mapping output")
	}
	return m, nil
}