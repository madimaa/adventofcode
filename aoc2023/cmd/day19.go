/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/madimaa/adventofcode/aoc2023/day19"
	"github.com/spf13/cobra"
)

// day19Cmd represents the day19 command
var day19Cmd = &cobra.Command{
	Use:   "day19",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		day19.Part1()
		day19.Part2()
	},
}

func init() {
	rootCmd.AddCommand(day19Cmd)
}
