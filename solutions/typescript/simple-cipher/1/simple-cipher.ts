export class SimpleCipher {
  key: string;

  constructor(key?: string) {
    // If no key is provided, generate a random 100-character lowercase key
    if (!key) {
      this.key = Array.from({ length: 100 }, () =>
        String.fromCharCode(97 + Math.floor(Math.random() * 26))
      ).join('');
    } else {
      // Validate key: must be lowercase letters only
      if (!/^[a-z]+$/.test(key)) {
        throw new Error('Key must contain only lowercase letters.');
      }
      this.key = key;
    }
  }

  // Helper: shift a lowercase character by a given amount (wraps around)
  private shiftChar(char: string, amount: number): string {
    const base = 'a'.charCodeAt(0);
    return String.fromCharCode(((char.charCodeAt(0) - base + amount + 26) % 26) + base);
  }

  encode(plainText: string): string {
    let result = '';
    for (let i = 0; i < plainText.length; i++) {
      const keyShift = this.key.charCodeAt(i % this.key.length) - 97;
      result += this.shiftChar(plainText[i], keyShift);
    }
    return result;
  }

  decode(cipherText: string): string {
    let result = '';
    for (let i = 0; i < cipherText.length; i++) {
      const keyShift = this.key.charCodeAt(i % this.key.length) - 97;
      result += this.shiftChar(cipherText[i], -keyShift);
    }
    return result;
  }
}
