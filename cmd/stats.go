package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show stats of the trash directory, like 'du -sh'",
	Run: func(cmd *cobra.Command, args []string) {
		stats()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func stats() {
	root := os.ExpandEnv(viper.GetString("DIR"))
	var total float64
	e := afero.Walk(afero.NewOsFs(), root, func(path string, info fs.FileInfo, err error) error {
		total += float64(info.Size())
		return nil
	})
	if e != nil {
		panic(e)
	}

	units := []string{"B", "KB", "MB", "GB", "TB"}

	var unitIndex int

	for total >= 1024 && unitIndex < len(units)-1 {
		total /= 1024
		unitIndex++
	}

	fmt.Printf("Size: %.2f %s\n", total, units[unitIndex])
}
