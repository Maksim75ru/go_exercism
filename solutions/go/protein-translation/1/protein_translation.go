package proteintranslation
import "errors"

var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop codon")

func FromRNA(rna string) ([]string, error) {
	var aminoAcids []string

	for i := 0; i < len(rna); i += 3 {
        if len(rna) < i+3 {
			return nil, ErrInvalidBase
		}
		codon := rna[i : i+3]
		acid, err := FromCodon(codon)

		if err != nil {
			if errors.Is(err, ErrStop) {
				return aminoAcids, nil
			}
			return nil, err
		}
		aminoAcids = append(aminoAcids, acid)
	}
	return aminoAcids, nil
}

func FromCodon(codon string) (string, error) {
	CodonAminoMap := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	acid, ok := CodonAminoMap[codon]

	if !ok {
		return "", ErrInvalidBase
	}
	if acid == "STOP" {
		return "", ErrStop

	}
	return acid, nil
}
