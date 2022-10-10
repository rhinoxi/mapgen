/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"math/rand"
	"time"

	"github.com/rhinoxi/mapgen/bsp"
	"github.com/rhinoxi/mapgen/util"
	"github.com/spf13/cobra"
)

var (
	bspDepth int
)

// bspCmd represents the bsp command
var bspCmd = &cobra.Command{
	Use:   "bsp",
	Short: "bsp tree",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if seed == 0 {
			rand.Seed(time.Now().UnixNano())
			seed = rand.Int63n(1000000)
		}
		m := bsp.Gen(mapWidth, mapHeight, bspDepth, seed)
		if output == "" {
			util.Print(m)
		} else {
			util.Draw(m, output)
		}
	},
}

func init() {
	rootCmd.AddCommand(bspCmd)

	bspCmd.Flags().IntVar(&bspDepth, "depth", 3, "depth")
}
