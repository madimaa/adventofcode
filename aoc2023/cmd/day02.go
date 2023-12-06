/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day02"
	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day02.Part1()
		day02.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day02Cmd)
}
