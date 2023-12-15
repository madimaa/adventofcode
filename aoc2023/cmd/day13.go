/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day13"
	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use:   "day13",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day13.Part1()
		day13.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day13Cmd)
}
