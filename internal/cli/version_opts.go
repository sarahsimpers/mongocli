// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/store"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type VersionOpts struct {
	store store.ServiceVersionDescriber
}

// InitVersion allow to init the VersionOpts in a functional way.
func (opts *VersionOpts) InitVersion() error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()))
	return err
}

// ServiceVersion retrieves and parses the service version.
func (opts *VersionOpts) ServiceVersion() (*semver.Version, error) {
	v, err := opts.store.ServiceVersion()
	if err != nil {
		return nil, err
	}
	return ParseServiceVersion(v)
}

// ParseServiceVersion parses service version into semver.Version.
func ParseServiceVersion(v *atlas.ServiceVersion) (*semver.Version, error) {
	versionParts := strings.Split(v.Version, ".")

	const maxVersionParts = 2

	if len(versionParts) > maxVersionParts {
		versionParts = versionParts[0:maxVersionParts] // ops manager versions are not semantic this is converting it to x.y
	}

	newVersion := strings.Join(versionParts, ".")

	return semver.NewVersion(newVersion)
}
