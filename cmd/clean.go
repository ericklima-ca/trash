package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var yes bool

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean the trash directory.",
	Run: func(cmd *cobra.Command, args []string) {
		if !yes {
			fmt.Printf("clean trash dir!\n")
			fmt.Printf("confirm?: (yes/no) ")
			var input string
			fmt.Scanln(&input)
			if input == "yes" {
				if e := clean(); e != nil {
					log.Fatalln(e)
				}
			}
		}
		fmt.Println("trash cleared.")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Skip question with 'yes' answer.")
}

func clean() error {
	fs := afero.NewOsFs()
	dir := os.ExpandEnv(viper.GetString("DIR"))
	if files, e := afero.ReadDir(fs, dir); e != nil {
		return e
	} else {
		for _, f := range files {
			if e := fs.RemoveAll(path.Join(dir, f.Name())); e != nil {
				return e
			}
		}
	}
	return nil
}
