/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2024/day01"
	"github.com/spf13/cobra"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day01.Part1()
		day01.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day01Cmd)
}
