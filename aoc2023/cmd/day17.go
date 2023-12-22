/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day17"
	"github.com/spf13/cobra"
)

// day17Cmd represents the day17 command
var day17Cmd = &cobra.Command{
	Use:   "day17",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day17.Part1()
		day17.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day17Cmd)
}
