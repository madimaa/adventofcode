/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day08"
	"github.com/spf13/cobra"
)

// day08Cmd represents the day08 command
var day08Cmd = &cobra.Command{
	Use:   "day08",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day08.Part1()
		day08.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day08Cmd)
}
