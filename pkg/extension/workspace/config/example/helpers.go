// Copyright 2020 Tetrate
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
// limitations under the License.

package example

import (
	"github.com/pkg/errors"
	"github.com/tetratelabs/multierror"

	"github.com/tetratelabs/getenvoy/pkg/extension/workspace/config"
	"github.com/tetratelabs/getenvoy/pkg/types"
)

// NewExampleDescriptor returns a new example descriptor instance.
func NewExampleDescriptor() *Descriptor {
	return &Descriptor{
		Meta: config.Meta{
			Kind: Kind,
		},
	}
}

// Default sets default values to optional fields.
func (d *Descriptor) Default() {
}

// Validate returns an error if Descriptor is not valid.
func (d *Descriptor) Validate() (errs error) {
	if err := d.Runtime.Validate(); err != nil {
		errs = multierror.Append(errs, errors.Wrap(err, "runtime configuration is not valid"))
	}
	return
}

// Validate returns an error if Runtime is not valid.
func (r *Runtime) Validate() (errs error) {
	if r == nil {
		return
	}
	if err := r.Envoy.Validate(); err != nil {
		errs = multierror.Append(errs, errors.Wrap(err, "Envoy configuration is not valid"))
	}
	return
}

// Validate returns an error if EnvoyRuntime is not valid.
func (r *EnvoyRuntime) Validate() (errs error) {
	if r == nil {
		return
	}
	if r.Version != "" {
		if _, err := types.ParseReference(r.Version); err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "Envoy version is not valid"))
		}
	}
	return
}

// GetRuntime returns Runtime the example should be run in.
func (d *Descriptor) GetRuntime() *Runtime {
	if d == nil {
		return nil
	}
	return d.Runtime
}

// GetEnvoy returns EnvoyRuntime the example should be run in.
func (r *Runtime) GetEnvoy() *EnvoyRuntime {
	if r == nil {
		return nil
	}
	return r.Envoy
}

// GetVersion returns Envoy version the example should be run in.
func (r *EnvoyRuntime) GetVersion() string {
	if r == nil {
		return ""
	}
	return r.Version
}
