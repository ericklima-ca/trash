package cmd

import (
	"log"
	"os"
	p "path"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var recursive bool
var force bool

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes (moves) files to the trash directory.",
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		for _, a := range args {
			isDir, err := afero.IsDir(fs, a)
			if err != nil {
				log.Fatalln(err)
			}
			if isDir && recursive {
				remove(fs, a)
			} else if !isDir {
				remove(fs, a)
			} else {
				log.Fatalln(a, "is a dir. try using the flag '-r'.")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "")
	removeCmd.Flags().BoolVarP(&force, "force", "f", true, "")
}

func remove(fs afero.Fs, path string) error {
	root := os.ExpandEnv(viper.GetString("DIR"))
	now := getTimestamp()
	tfs := afero.NewBasePathFs(afero.NewOsFs(), root)
	tfs.MkdirAll(now, 0700)
	newPath := p.Join(root, now, path)
	if e := fs.Rename(path, newPath); e != nil {
		return e
	}
	return nil
}

func getTimestamp() string {
	return time.Now().Format("20060102150405")
}
