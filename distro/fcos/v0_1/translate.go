// Copyright 2019 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.)

package fcos_0_1

import (
	"github.com/coreos/fcct/translate"

	types3_0 "github.com/coreos/ignition/v2/config/v3_0/types"
	types3_1 "github.com/coreos/ignition/v2/config/v3_1_experimental/types"
)

// ToIgn3_0 takes a config and merges in the distro specific bits.
func (f Fcos) ToIgn3_0(in types3_0.Config) (types3_0.Config, translate.TranslationSet, error) {
	return in, translate.TranslationSet{}, nil
}

func (f Fcos) ToIgn3_1(in types3_1.Config) (types3_1.Config, translate.TranslationSet, error) {
	return in, translate.TranslationSet{}, nil
}
