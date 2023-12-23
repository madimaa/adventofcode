/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day22"
	"github.com/spf13/cobra"
)

// day22Cmd represents the day22 command
var day22Cmd = &cobra.Command{
	Use:   "day22",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day22.Part1()
		day22.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day22Cmd)
}
