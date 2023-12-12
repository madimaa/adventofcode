/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day12"
	"github.com/spf13/cobra"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day12.Part1()
		day12.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)
}
