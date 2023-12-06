/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day04"
	"github.com/spf13/cobra"
)

// day04Cmd represents the day04 command
var day04Cmd = &cobra.Command{
	Use:   "day04",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day04.Part1()
		day04.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day04Cmd)
}
