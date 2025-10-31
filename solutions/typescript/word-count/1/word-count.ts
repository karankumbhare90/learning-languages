export function count(text: string): Map<string, number> {
  const lower = text.toLowerCase();

  // Extract valid words (letters/numbers, allow inner apostrophes)
  const words = lower.match(/[a-z0-9]+(?:'[a-z0-9]+)*/g) || [];

  const result = new Map<string, number>();

  for (const w of words) {
    result.set(w, (result.get(w) || 0) + 1);
  }

  return result;
}
