/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day14"
	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use:   "day14",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day14.Part1()
		day14.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day14Cmd)
}
