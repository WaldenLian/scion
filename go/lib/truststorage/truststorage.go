// Copyright 2018 Anapaya Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package truststorage provides a "factory" for trust database.
// A config containing the backend type and the connection string
// are used to create a specific trust db.
package truststorage

import (
	"fmt"
	"io"

	"github.com/scionproto/scion/go/lib/common"
	"github.com/scionproto/scion/go/lib/config"
	"github.com/scionproto/scion/go/lib/infra/modules/db"
	"github.com/scionproto/scion/go/lib/infra/modules/trust/trustdb"
	"github.com/scionproto/scion/go/lib/infra/modules/trust/trustdb/trustdbsqlite"
	"github.com/scionproto/scion/go/lib/log"
	"github.com/scionproto/scion/go/lib/util"
)

type Backend string

const (
	BackendNone   Backend = ""
	BackendSqlite Backend = "sqlite"
)

const (
	BackendKey    = "backend"
	ConnectionKey = "connection"
)

var _ (config.Config) = (*TrustDBConf)(nil)

// TrustDBConf is the configuration for the connection to the trust database.
type TrustDBConf map[string]string

// InitDefaults choses the sqlite backend if no backend is set.
func (cfg *TrustDBConf) InitDefaults() {
	if *cfg == nil {
		*cfg = make(TrustDBConf)
	}
	m := *cfg
	util.LowerKeys(m)
	if cfg.Backend() == BackendNone {
		m[BackendKey] = string(BackendSqlite)
	}
}

func (cfg *TrustDBConf) Backend() Backend {
	return Backend((*cfg)[BackendKey])
}

func (cfg *TrustDBConf) Connection() string {
	return (*cfg)[ConnectionKey]
}

func (cfg *TrustDBConf) MaxOpenConns() (int, bool) {
	return db.ConfiguredMaxOpenConns(*cfg)
}

func (cfg *TrustDBConf) MaxIdleConns() (int, bool) {
	return db.ConfiguredMaxIdleConns(*cfg)
}

func (cfg *TrustDBConf) Sample(dst io.Writer, path config.Path, ctx config.CtxMap) {
	config.WriteString(dst, fmt.Sprintf(trustDbSample, ctx[config.ID]))
}

func (cfg *TrustDBConf) ConfigName() string {
	return "trustDB"
}

func (cfg *TrustDBConf) Validate() error {
	if err := db.ValidateConfigLimits(*cfg); err != nil {
		return err
	}
	if err := cfg.validateBackend(); err != nil {
		return err
	}
	return nil
}

func (cfg *TrustDBConf) validateBackend() error {
	switch cfg.Backend() {
	case BackendSqlite:
		return nil
	case BackendNone:
		return common.NewBasicError("No backend set", nil)
	}
	return common.NewBasicError("Unsupported backend", nil, "backend", cfg.Backend())
}

// New creates a TrustDB from the config.
func (cfg *TrustDBConf) New() (trustdb.TrustDB, error) {
	log.Info("Connecting TrustDB", "backend", cfg.Backend(), "connection", cfg.Connection())
	var err error
	var tdb trustdb.TrustDB

	switch cfg.Backend() {
	case BackendSqlite:
		tdb, err = trustdbsqlite.New(cfg.Connection())
	default:
		return nil, common.NewBasicError("Unsupported backend", nil, "backend", cfg.Backend())
	}

	if err != nil {
		return nil, err
	}
	db.SetConnLimits(cfg, tdb)
	return tdb, nil
}
