export function encode(plainText: string): string {
  if (typeof plainText !== "string") return "";

  const alphabet = "abcdefghijklmnopqrstuvwxyz";
  const reversed = alphabet.split("").reverse().join("");
  const map: Record<string, string> = {};

  // build substitution map
  for (let i = 0; i < 26; i++) {
    map[alphabet[i]] = reversed[i];
  }

  // normalize text (only letters + digits)
  const cleaned = plainText.toLowerCase().replace(/[^a-z0-9]/g, "");

  // apply substitution
  let transformed = "";
  for (const ch of cleaned) {
    transformed += map[ch] ?? ch;
  }

  // group into 5-character chunks
  return transformed.replace(/(.{5})(?=.)/g, "$1 ");
}

export function decode(cipherText: string): string {
  if (typeof cipherText !== "string") return "";

  const alphabet = "abcdefghijklmnopqrstuvwxyz";
  const reversed = alphabet.split("").reverse().join("");
  const map: Record<string, string> = {};

  for (let i = 0; i < 26; i++) {
    map[alphabet[i]] = reversed[i];
  }

  // remove spaces, reverse substitution
  const cleaned = cipherText.replace(/\s+/g, "").toLowerCase();

  let decoded = "";
  for (const ch of cleaned) {
    decoded += map[ch] ?? ch;
  }

  return decoded;
}
