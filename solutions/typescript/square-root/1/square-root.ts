export function squareRoot(radicand: unknown): unknown {
  // Validate input
  if (typeof radicand !== "number" || !Number.isInteger(radicand) || radicand <= 0) {
    throw new Error("Input must be a positive whole number");
  }

  let n = radicand;

  // Binary search for integer square root
  let low = 1;
  let high = n;

  while (low <= high) {
    const mid = (low + high) >> 1; // integer midpoint
    const square = mid * mid;

    if (square === n) {
      return mid; // found exact root
    }

    if (square < n) {
      low = mid + 1;
    } else {
      high = mid - 1;
    }
  }

  // Per problem statement, radicand always has an integer square root
  throw new Error("Input is not a perfect square");
}
