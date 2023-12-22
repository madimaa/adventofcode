/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day18"
	"github.com/spf13/cobra"
)

// day18Cmd represents the day18 command
var day18Cmd = &cobra.Command{
	Use:   "day18",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day18.Part1()
		day18.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day18Cmd)
}
