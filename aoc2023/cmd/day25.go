/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day25"
	"github.com/spf13/cobra"
)

// day25Cmd represents the day25 command
var day25Cmd = &cobra.Command{
	Use:   "day25",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day25.Part1()
		day25.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day25Cmd)
}
