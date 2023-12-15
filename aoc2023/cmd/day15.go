/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day15"
	"github.com/spf13/cobra"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use:   "day15",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day15.Part1()
		day15.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day15Cmd)
}
