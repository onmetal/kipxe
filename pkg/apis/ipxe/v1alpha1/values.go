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

package v1alpha1

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/gardener/controller-manager-library/pkg/types/infodata/simple"
)

// Values is a workarround for kubebuilder to be able to generate
// an API spec. The Values MUST be marked with "-" to avoud errors.

// Values is used to specify any document structure
type Values struct {
	simple.Values `json:"-"`
}

func (in *Values) DeepCopy() *Values {
	if in == nil {
		return nil
	}
	return &Values{in.Values.DeepCopy()}
}

func (this Values) MarshalJSON() ([]byte, error) {
	if this.Values == nil {
		return []byte("null"), nil
	}
	return this.Values.Marshal()
}

func (this *Values) UnmarshalJSON(in []byte) error {
	if this == nil {
		return errors.New("Values: UnmarshalJSON on nil pointer")
	}
	if !bytes.Equal(in, []byte("null")) {
		return json.Unmarshal(in, &this.Values)
	}
	return nil
}
