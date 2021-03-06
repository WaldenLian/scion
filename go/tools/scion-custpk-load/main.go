// Copyright 2019 Anapaya Systems
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

package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/scionproto/scion/go/lib/common"
	"github.com/scionproto/scion/go/lib/config"
	"github.com/scionproto/scion/go/lib/env"
	"github.com/scionproto/scion/go/lib/infra/modules/trust/trustdb"
	"github.com/scionproto/scion/go/lib/truststorage"
)

var _ config.Config = (*Config)(nil)

type Config struct {
	TrustDB truststorage.TrustDBConf
}

func (cfg *Config) InitDefaults() {
	cfg.TrustDB.InitDefaults()
}

func (cfg *Config) Validate() error {
	return cfg.TrustDB.Validate()
}

func (cfg *Config) Sample(dst io.Writer, path config.Path, _ config.CtxMap) {
	config.WriteSample(dst, path, config.CtxMap{config.ID: "cs-1"}, &cfg.TrustDB)
}

func (cfg *Config) ConfigName() string {
	return "scion-custpk-load_config"
}

var (
	custDir = flag.String("customers", "", "The folder containing the customer keys")

	cfg     Config
	trustDB trustdb.TrustDB
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	env.AddFlags()
	flag.Parse()
	if ret, ok := checkFlags(); !ok {
		return ret
	}
	if err := loadConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		flag.Usage()
		return 1
	}
	files, loadedCusts, err := LoadCustomers(*custDir, trustDB)
	printSummary(files, loadedCusts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during loading: %s\n", err)
		return 1
	}
	return 0
}

// checkFlags checks the flags.
// The first return value is the return code of the program. The second value
// indicates whether the program can continue with its execution or should exit.
func checkFlags() (int, bool) {
	if ret, ok := env.CheckFlags(&cfg); !ok {
		return ret, ok
	}
	if *custDir == "" {
		fmt.Fprintln(os.Stderr, "Err: Missing customers argument")
		flag.Usage()
		return 1, false
	}
	return 0, true
}

func loadConfig() error {
	if _, err := toml.DecodeFile(env.ConfigFile(), &cfg); err != nil {
		return common.NewBasicError("Failed to load config", err)
	}
	cfg.InitDefaults()
	err := cfg.Validate()
	if err != nil {
		return common.NewBasicError("Unable to validate config", err)
	}
	if trustDB, err = cfg.TrustDB.New(); err != nil {
		return common.NewBasicError("Failed to init the database", err)
	}
	return nil
}

func printSummary(files []string, loadedCusts []*CustKeyMeta) {
	fmt.Println("Successfully processed files:")
	for _, f := range files {
		fmt.Println(f)
	}
	fmt.Println("Successfully stored customers:")
	for _, cust := range loadedCusts {
		fmt.Printf("IA: %s, Version: %d\n", cust.IA, cust.Version)
	}
}
