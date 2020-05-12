// Copyright 2020 Google LLC
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

// Package publish defines the exposure keys publishing API.
package publish

import (
	"time"

	"github.com/google/exposure-notifications-server/internal/database"
	"github.com/google/exposure-notifications-server/internal/dbapiconfig"
)

// Environment repsresnts the supported environment variables for the exposure server.
type Environment struct {
	MinRequestDuration       time.Duration `envconfig:"TARGET_REQUEST_DURATION" default:"5s"`
	MaxKeysOnPublish         int           `envconfig:"MAX_KEYS_ON_PUBLISH" default:"14"`
	MaxIntervalAge           time.Duration `envconfig:"MAX_INTERVAL_AGE_ON_PUBLISH" default:"360h"`
	MaxIntervalFuture        time.Duration `envconfig:"MAX_INTERVAL_FUTURE_ON_PUBLISH" default:"2h"`
	APIConfigRefreshDuration time.Duration `envconfig:"CONFIG_REFRESH_DURATION" default:"5m"`

	APIConfigOpts dbapiconfig.ConfigOpts

	Database database.Environment
}