export function score(x: unknown, y: unknown): unknown {
  if (typeof x !== 'number' || typeof y !== 'number') {
    throw new Error('Invalid input: x and y must be numbers');
  }

  // Calculate distance from center
  const distance = Math.sqrt(x * x + y * y);

  // Check which circle the dart landed in
  if (distance <= 1) return 10;      // Inner circle
  if (distance <= 5) return 5;       // Middle circle
  if (distance <= 10) return 1;      // Outer circle
  return 0;
}
