/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2024/day09"
	"github.com/spf13/cobra"
)

// day09Cmd represents the day09 command
var day09Cmd = &cobra.Command{
	Use:   "day09",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day09.Part1()
		day09.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day09Cmd)
}
