export class Squares {
  private n: number;

  constructor(count: number) {
    this.n = count;
  }

  // Formula for sum of squares: n(n + 1)(2n + 1) / 6
  get sumOfSquares(): number {
    return (this.n * (this.n + 1) * (2 * this.n + 1)) / 6;
  }

  // Formula for square of sum: (n(n + 1)/2)²
  get squareOfSum(): number {
    const sum = (this.n * (this.n + 1)) / 2;
    return sum * sum;
  }

  // Difference between the two
  get difference(): number {
    return this.squareOfSum - this.sumOfSquares;
  }
}
