/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day20"
	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use:   "day20",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day20.Part1()
		day20.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day20Cmd)
}
