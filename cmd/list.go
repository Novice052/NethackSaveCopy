/*
Copyright © 2020 John Abbott <immersed101@protonmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists NetHack save files and backups in the save directory.",
	Long:  `Lists NetHack save files and backups in the default NetHack 3.6 save directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// CHANGE: list command now performs real directory scanning instead of placeholder output.
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		// CHANGE: use current user home path for portable Windows save directory detection.
		saveDir := usr.HomeDir + "\\AppData\\Local\\NetHack\\3.6"
		entries, err := os.ReadDir(saveDir)
		if err != nil {
			log.Fatal(err)
		}

		const saveSuffix = ".NetHack-saved-game"
		const backupSuffix = ".NetHack-saved-game-bak"

		fmt.Println("Save directory:", saveDir)
		fmt.Println("Detected NetHack saves/backups:")

		foundAny := false
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			name := entry.Name()
			if strings.HasSuffix(name, backupSuffix) {
				character := strings.TrimSuffix(name, backupSuffix)
				fmt.Printf("- %s (backup)\n", character)
				foundAny = true
				continue
			}

			if strings.HasSuffix(name, saveSuffix) {
				character := strings.TrimSuffix(name, saveSuffix)
				fmt.Printf("- %s (active save)\n", character)
				foundAny = true
			}
		}

		if !foundAny {
			fmt.Println("(none found)")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
