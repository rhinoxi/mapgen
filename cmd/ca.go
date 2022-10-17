/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"math/rand"
	"time"

	ca "github.com/rhinoxi/mapgen/cellularautomata"
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
		fmt.Println("\nIsland area:")
		for i, count := range util.CountIsland(m) {
			fmt.Printf("    Island %d: %d\n", i, count)
		}
	},
}

func init() {
	rootCmd.AddCommand(caCmd)

	caCmd.Flags().IntVar(&caIters, "iter", 3, "iterations")
	caCmd.Flags().IntVar(&noiseDensity, "density", 60, "noise density")
}
