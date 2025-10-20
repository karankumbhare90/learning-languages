export class Matrix {
  private _rows: number[][];

  constructor(matrixString: string) {
    this._rows = matrixString
      .trim()
      .split('\n')
      .map(row =>
        row
          .trim()
          .split(/\s+/)
          .map(num => Number(num))
      );
  }

  get rows(): number[][] {
    return this._rows;
  }

  get columns(): number[][] {
    return this._rows[0].map((_, colIndex) =>
      this._rows.map(row => row[colIndex])
    );
  }
}
