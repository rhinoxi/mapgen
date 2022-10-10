/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"math/rand"
	"time"

	ca "github.com/rhinoxi/mapgen/cellularAutomata"
	"github.com/rhinoxi/mapgen/util"
	"github.com/spf13/cobra"
)

var (
	caIters      int
	noiseDensity int
)

// caCmd represents the ca command
var caCmd = &cobra.Command{
	Use:   "ca",
	Short: "cellular automata map generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if seed == 0 {
			rand.Seed(time.Now().UnixNano())
			seed = rand.Int63n(1000000)
		}
		m := ca.Gen(mapWidth, mapHeight, noiseDensity, caIters, seed)
		if output == "" {
			util.Print(m)
		} else {
			util.Draw(m, output)
		}
	},
}

func init() {
	rootCmd.AddCommand(caCmd)

	caCmd.Flags().IntVar(&caIters, "iter", 3, "iterations")
	caCmd.Flags().IntVar(&noiseDensity, "density", 60, "noise density")
}
