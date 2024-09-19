package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-19


A golang program to estimate the number of the sites with the unified base estimates. You can pass as many genomes aligned as you want.
The order of the genome aligned doent matter.

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var alignment string

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This estimates the site proportion in your whole genome or gene specific alignment",
	Run:  flagsFunc,
}

func init() {
	rootCmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
}

func flagsFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
	}

	counterA := 0
	counterT := 0
	counterG := 0
	counterC := 0

	counterAsite := []int{}
	counterTsite := []int{}
	counterGsite := []int{}
	counterCsite := []int{}

	for i := 0; i < len(sequenceCap)-1; i++ {
		for j := 0; j < len(sequenceCap[0]); j++ {
			if string(sequenceCap[i][j]) == "A" && string(sequenceCap[i+1][j]) == "A" {
				counterA++
				counterAsite = append(counterAsite, j)
			}
			if string(sequenceCap[i][j]) == "T" && string(sequenceCap[i+1][j]) == "T" {
				counterT++
				counterTsite = append(counterTsite, j)
			}
			if string(sequenceCap[i][j]) == "G" && string(sequenceCap[i+1][j]) == "G" {
				counterG++
				counterGsite = append(counterGsite, j)
			}
			if string(sequenceCap[i][j]) == "C" && string(sequenceCap[i+1][j]) == "C" {
				counterC++
				counterCsite = append(counterCsite, j)
			}

		}
	}
	fmt.Printf(
		"The alignment counts for the A unified bases to the rest of the same sites in the block are: %d",
		counterA,
	)
	fmt.Printf(
		"The alignment counts for the T unified bases to the rest of the same sites in the block are: %d",
		counterT,
	)
	fmt.Printf(
		"The alignment counts for the G unified bases to the rest of the same sites in the block are: %d",
		counterG,
	)
	fmt.Printf(
		"The alignment counts for the C unified bases to the rest of the same sites in the block are: %d",
		counterC,
	)

	for i := range counterAsite {
		fmt.Printf("The A sites are %d", counterAsite[i])
	}

	for i := range counterTsite {
		fmt.Printf("The T sites are %d", counterTsite[i])
	}

	for i := range counterGsite {
		fmt.Printf("The G sites are %d", counterGsite[i])
	}
	for i := range counterCsite {
		fmt.Printf("The C sites are %d", counterCsite[i])
	}
}
