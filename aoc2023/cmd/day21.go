/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day21"
	"github.com/spf13/cobra"
)

// day21Cmd represents the day21 command
var day21Cmd = &cobra.Command{
	Use:   "day21",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day21.Part1()
		day21.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day21Cmd)
}
