/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2024/day06"
	"github.com/spf13/cobra"
)

// day06Cmd represents the day06 command
var day06Cmd = &cobra.Command{
	Use:   "day06",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day06.Part1()
		day06.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day06Cmd)
}
