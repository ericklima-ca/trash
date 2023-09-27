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

var trashDir string
var configFilePath string
var userHomeDir string

var initCmd = &cobra.Command{
	Use:   "init",
	Args:  cobra.NoArgs,
	Short: "Initialize the trash directory. Default to '$HOME/.trash'",
	Run: func(cmd *cobra.Command, args []string) {
		var pathToTrash string

		if dir := os.ExpandEnv(viper.GetString("DIR")); dir != "" {
			pathToTrash = dir
		} else {
			pathToTrash = os.ExpandEnv(trashDir)
		}
		fs := afero.NewOsFs()
		b, e := afero.DirExists(fs, pathToTrash)
		if e != nil {
			panic(e)
		}
		if !b {
			initFolder(pathToTrash)
		} else {
			log.Fatalln("dir already exists")
		}
		if b, e := afero.Exists(fs, os.ExpandEnv(configFilePath)); e != nil {
			panic(e)
		} else if !b {
			if e := initDefault(); e != nil {
				panic(e)
			}
		}
	},
}

func init() {

	rootCmd.AddCommand(initCmd)
	userHomeDir, _ = os.UserHomeDir()
	configFilePath = path.Join(userHomeDir, ".trashconfig")
	trashDirPath := path.Join(userHomeDir, ".trash")

	initCmd.Flags().StringVarP(&trashDir, "dir", "d", trashDirPath, "Path to trash dir")

	viper.SetConfigName(".trashconfig")
	viper.SetConfigType("env")
	viper.AddConfigPath(userHomeDir)
	viper.AddConfigPath(".") // DEV
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func initFolder(path string) error {
	fs := afero.NewOsFs()
	if err := fs.MkdirAll(path, 0700); err != nil {
		return err
	}
	return nil
}

func initDefault() error {
	trashConfiFile := fmt.Sprintf("DIR=\"%s/.trash\"", userHomeDir)
	fs := afero.NewOsFs()
	if f, e := fs.Create(configFilePath); e != nil {
		return e
	} else {
		if _, err := f.Write([]byte(trashConfiFile)); err != nil {
			return err
		}
		f.Close()
		return nil
	}
}
