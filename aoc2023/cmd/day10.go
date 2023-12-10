/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day10"
	"github.com/spf13/cobra"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day10.Part1()
		day10.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}
