/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day05"
	"github.com/spf13/cobra"
)

// day05Cmd represents the day05 command
var day05Cmd = &cobra.Command{
	Use:   "day05",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day05.Part1()
		day05.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day05Cmd)
}
