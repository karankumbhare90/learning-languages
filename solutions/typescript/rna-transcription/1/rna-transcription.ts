export function toRna(dna: string): string {
  const mapping: Record<string, string> = {
    G: 'C',
    C: 'G',
    T: 'A',
    A: 'U'
  };

  return dna
    .split('')
    .map((nucleotide: string): string => {
      if (!(nucleotide in mapping)) {
        throw new Error('Invalid input DNA.');
      }
      return mapping[nucleotide];
    })
    .join('');
}
