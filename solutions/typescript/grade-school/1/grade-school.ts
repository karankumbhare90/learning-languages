export class GradeSchool {
  private _roster: Record<number, string[]> = {};

  add(student: string, grade: number): void {
    // Remove student from any existing grade (if re-added to another)
    for (const g in this._roster) {
      this._roster[g] = this._roster[g].filter(name => name !== student);
    }

    if (!this._roster[grade]) {
      this._roster[grade] = [];
    }

    // Add student and keep names sorted alphabetically
    this._roster[grade].push(student);
    this._roster[grade].sort();
  }

  roster(): Record<number, string[]> {
    // Deep clone so outside code can’t mutate internal data
    const copy: Record<number, string[]> = {};
    for (const [grade, students] of Object.entries(this._roster)) {
      copy[Number(grade)] = [...students];
    }
    return copy;
  }

  grade(grade: number): string[] {
    return this._roster[grade] ? [...this._roster[grade]] : [];
  }
}
