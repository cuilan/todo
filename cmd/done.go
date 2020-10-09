package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/cuilan/todo/entity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark Item as Done.",
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := entity.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalf(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")
		sort.Sort(entity.ByPri(items))
		_ = entity.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items.")
	}
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
