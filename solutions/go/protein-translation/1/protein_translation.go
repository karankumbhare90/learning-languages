package protein

import (
	"errors"
)

// ErrStop is returned when a STOP codon is found
var ErrStop = errors.New("STOP codon encountered")

// ErrInvalidBase is returned when the codon is invalid
var ErrInvalidBase = errors.New("invalid codon")

// FromCodon translates a single codon into an amino acid
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA splits RNA into codons and translates them until STOP
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			break
		}
		if err != nil {
			return nil, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
