package strand

func ToRNA(dna string) string {
	var rna []rune

    for _, nucleotide := range dna {
        switch nucleotide {
            case 'G':
            	rna = append(rna, 'C')
            case 'C':
            	rna = append(rna, 'G')
            case 'T':
            	rna = append(rna, 'A')
            case 'A':
            	rna = append(rna, 'U')
        }
    }

    return string(rna)
}
