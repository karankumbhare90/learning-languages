export function isPangram(sentence: string): boolean {
  const lower = sentence.toLowerCase();

  const letters = new Set<string>();

  for (const ch of lower) {
    if (ch >= 'a' && ch <= 'z') {
      letters.add(ch);
    }
  }

  return letters.size === 26;
}
