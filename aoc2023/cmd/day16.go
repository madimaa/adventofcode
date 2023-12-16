/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day16"
	"github.com/spf13/cobra"
)

// day16Cmd represents the day16 command
var day16Cmd = &cobra.Command{
	Use:   "day16",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day16.Part1()
		day16.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day16Cmd)
}
