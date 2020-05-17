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

type InfoBase struct {
	Registry  *Registry
	Documents *Documents
	Profiles  *Profiles
	Matchers  *Matchers
}

func (this *InfoBase) SetDocument(e *Document) NameSet {
	users := this.Documents.Set(e)
	users = this.Profiles.Recheck(users)
	return this.Matchers.Recheck(users)
}

func (this *InfoBase) SetProfile(e *Profile) (NameSet, error) {
	users, err := this.Profiles.Set(e)
	return this.Matchers.Recheck(users), err
}

func (this *InfoBase) SetMatcher(e *Matcher) error {
	return this.Matchers.Set(e)
}
