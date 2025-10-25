export function find(haystack: unknown, needle: unknown): number | never {
  // Validate inputs
  if (!Array.isArray(haystack)) {
    throw new Error('Haystack must be an array.');
  }
  if (typeof needle !== 'number') {
    throw new Error('Needle must be a number.');
  }

  let left = 0;
  let right = haystack.length - 1;

  while (left <= right) {
    const mid = Math.floor((left + right) / 2);
    const midValue = haystack[mid];

    if (midValue === needle) {
      return mid;
    } else if (midValue < needle) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  // If not found
  throw new Error('Value not in array');
}
