/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day11"
	"github.com/spf13/cobra"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day11.Part1()
		day11.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)
}
