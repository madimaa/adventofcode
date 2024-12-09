/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2024/day07"
	"github.com/spf13/cobra"
)

// day07Cmd represents the day07 command
var day07Cmd = &cobra.Command{
	Use:   "day07",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day07.Part1()
		day07.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day07Cmd)
}
