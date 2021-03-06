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

package ipxe

import (
	"github.com/gardener/controller-manager-library/pkg/controllermanager/controller"
	"github.com/gardener/controller-manager-library/pkg/ctxutil"

	"github.com/mandelsoft/kipxe/pkg/controllers"
	"github.com/mandelsoft/kipxe/pkg/kipxe"
)

var infobaseKey = ctxutil.SimpleKey("infobase")

func GetSharedInfoBase(controller controller.Interface) *InfoBase {
	return controller.GetOrCreateSharedValue(infobaseKey, func() interface{} {
		return NewInfoBase(controller)
	}).(*InfoBase)
}

type InfoBase struct {
	controller controller.Interface
	registry   *kipxe.Registry
	cache      *kipxe.DirCache
	mappers    *MetaDataMappers
	matchers   *BootMatchers
	profiles   *BootProfiles
	resources  *BootResources
}

func NewInfoBase(controller controller.Interface) *InfoBase {
	b := &InfoBase{
		controller: controller,
		registry:   controllers.GetSharedRegistry(controller),
	}

	b.resources = newResources(b)
	b.profiles = newProfiles(b)
	b.matchers = newMatchers(b)
	b.mappers = newMappers(b)
	return b
}

func (this *InfoBase) Setup() {
	this.resources.Setup(this.controller)
	this.profiles.Setup(this.controller)
	this.matchers.Setup(this.controller)
	this.mappers.Setup(this.controller)
}
