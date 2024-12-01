/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day23"
	"github.com/spf13/cobra"
)

// day23Cmd represents the day23 command
var day23Cmd = &cobra.Command{
	Use:   "day23",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day23.Part1()
		day23.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day23Cmd)
}
