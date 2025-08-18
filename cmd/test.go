/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
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
		problemType, _ := cmd.Flags().GetString("type")
		if problemType != "array" && problemType != "graph" && problemType != "tree" {
			cmd.Println("Only 'array', 'graph', and 'tree' options currently supported")
			os.Exit(1)
		}
		compile(args[0])
		compile(args[1])
		numCases, _ := cmd.Flags().GetInt("cases")
		multiQuery, _ := cmd.Flags().GetInt("multi-query")
		var pass bool = true
		for range numCases {
			var input string
			if problemType == "array" {
				num := rand.Intn(10) + 1
				input = gen.GenerateRandomArray(num, 1, 1e9)
			} else if problemType == "tree" {
				num := rand.Intn(10) + 1
				input = gen.GenerateRandomTree(num)
			} else if problemType == "graph" {
				num := rand.Intn(10) + 1
				num_edges := rand.Intn(num*(num-1)/2-num+2) + num - 1
				input = gen.GenerateRandomGraph(num, num_edges)
			}
			if multiQuery == 1 {
				input = "1\n" + input
			}
			var executable strings.Builder
			fmt.Fprintf(&executable, "./%s", args[0][:len(args[0])-4])
			testCmd := exec.Command(executable.String())
			testCmd.Stdin = strings.NewReader(input)
			output1, err1 := testCmd.Output()
			executable.Reset()
			fmt.Fprintf(&executable, "./%s", args[1][:len(args[1])-4])
			testCmd = exec.Command(executable.String())
			testCmd.Stdin = strings.NewReader(input)
			output2, err2 := testCmd.Output()
			if err1 == nil && err2 == nil {
				if string(output1) != string(output2) {
					pass = false
					cmd.Println("Outputs differed on following test case: ")
					cmd.Print(input)
					cmd.Printf("\nUser output: %s", string(output1))
					cmd.Printf("Solution output: %s", string(output2))
					os.Exit(1)
				}
			}
		}
		if pass {
			cmd.Println("All cases passed!")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().IntP("cases", "n", 1, "number of test cases")
	testCmd.Flags().StringP("type", "t", "array", "problem input category")
	testCmd.Flags().IntP("multi-query", "q", 1, "multiple queries in single test case (1/0)")
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
