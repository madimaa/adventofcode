/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2024/day03"
	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day03.Part1()
		day03.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day03Cmd)
}
