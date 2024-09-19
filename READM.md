go-alignment-proportion

- estimates the number of the unified bases in whole genome alignments
- faster calibration rates
- part of the alignmentGO package but can be used separately
- given a alignment it will profile, list the indices and the sites of the same bases.
```
[gauravsablok@fedora]~/Desktop/codecreatede/go-alignment-proportion% 
go run main.go -h
This estimates the site proportion in your whole genome or gene specific alignment

Usage:
  flags [flags]

Flags:
  -a, --alignmentfile string   a alignment file (default "align")
  -h, --help                   help for flags
[gauravsablok@fedora]~/Desktop/codecreatede/go-alignment-proportion% \
go run main.go -a ./samplefile/samplealignment.fasta
The alignment counts for the A unified bases to the rest of the same sites in the block are: 1
The alignment counts for the T unified bases to the rest of the same sites in the block are: 2
The alignment counts for the G unified bases to the rest of the same sites in the block are: 1
The alignment counts for the C unified bases to the rest of the same sites in the block are: 1
```

Gaurav Sablok
