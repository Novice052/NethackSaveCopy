/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restores a NetHack save file from a backup.",
	Long:  `Restores a NetHack save file from a backup.  Accepts one argument that names the backup file to be restored.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restore called")
		saveDir := "C:\\Users\\immer\\AppData\\Local\\NetHack\\3.6"
		src := saveDir + "\\" + args[0] + ".NetHack-saved-game-bak"
		// checks that the file exists
		sourceFileStat, err := os.Stat(src)
		if err != nil {
			log.Fatal(err)
		}
		// checks that the file isn't funky
		if !sourceFileStat.Mode().IsRegular() {
			log.Fatalf("%s is not a regular file", src)
		}
		// opens the file
		source, err := os.Open(src)
		if err != nil {
			log.Fatal(err)
		}
		defer source.Close()
		dst := saveDir + "\\" + args[0] + ".NetHack-saved-game"
		// creates a new file at dst
		destination, err := os.Create(dst)
		if err != nil {
			log.Fatal(err)
		}
		defer destination.Close()
		nBytes, err := io.Copy(destination, source)
		log.Print(nBytes)
		log.Fatal(err)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
