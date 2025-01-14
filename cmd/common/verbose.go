// Copyright 2023 The Inspektor Gadget authors
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

package common

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verbose bool

func checkVerboseFlag() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func AddVerboseFlag(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose", "v",
		false,
		"Print debug information",
	)
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		checkVerboseFlag()
	}
}
