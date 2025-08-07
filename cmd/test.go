/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"
	"strings"

	gen "github.com/amingma/stress/generators"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test [flags] <path to user solution> <path to model solution>",
	Short: "test user solution",
	Long: `Stress tests user solution against model implementation with optional flags to specify 
	input, time, and memory constraints.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Println("Too little arguments, please provide both paths")
			os.Exit(1)
		}
		if !verifyFile(args[0]) || !verifyFile(args[1]) {
			cmd.Println("Only c++ files supported soz")
			os.Exit(1)
		}
		compile(args[0])
		compile(args[1])
		numCases, _ := cmd.Flags().GetInt("test")
		for range numCases {
			arr := gen.GenerateRandomArray(3, 1, 6)
			testCmd := exec.Command("./main")
			testCmd.Stdin = strings.NewReader(arr)
			output, err := testCmd.Output()
			if err == nil {
				cmd.Println(string(output))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().IntP("test", "t", 1, "number of test cases")
	// testCmd.MarkFlagRequired("test")
}

func verifyFile(filePath string) bool {
	supportedExtensions := [1]string{".cpp"}
	var pos = false
	for _, s := range supportedExtensions {
		if len(filePath) >= len(s) && filePath[len(filePath)-len(s):] == s {
			pos = true
		}
	}
	return pos
}

func compile(filePath string) {
	cmd := exec.Command("g++", "--std=c++20", filePath, "-o", filePath[:len(filePath)-4])
	cmd.Run()
}
