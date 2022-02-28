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
	"io"
	"os"
	"os/user"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Args:  cobra.MinimumNArgs(1),
	Short: "Makes a backup copy of a NetHack save file.",
	Long:  `Makes a backup copy of a NetHack save file.  Accepts one argument that names the save file to be copied.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("copy called")
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(usr.HomeDir)

		// NetHack save files look like <Name.NetHack-saved-game>
		saveDir := usr.HomeDir + "\\AppData\\Local\\NetHack\\3.6"
		src := saveDir + "\\" + args[0] + ".NetHack-saved-game"
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
		dst := src + "-bak"
		// creates a new file at dst
		destination, err := os.Create(dst)
		if err != nil {
			log.Fatal(err)
		}
		defer destination.Close()
		nBytes, err := io.Copy(destination, source)
		log.Println(nBytes)
		log.Fatal(err)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
