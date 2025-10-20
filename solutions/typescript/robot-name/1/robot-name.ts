export class Robot {
  private static usedNames = new Set<string>();
  private _name: string | null = null;

  // Public getter for the robot's name
  public get name(): string {
    if (!this._name) {
      this._name = Robot.generateUniqueName();
    }
    return this._name;
  }

  // Reset robot to factory state
  public resetName(): void {
    this._name = null;
  }

  // Clear all used names
  public static releaseNames(): void {
    Robot.usedNames.clear();
  }

  // Generate a unique random name
  private static generateUniqueName(): string {
    let newName: string;

    do {
      newName = Robot.generateRandomName();
    } while (Robot.usedNames.has(newName)); // ensure uniqueness

    Robot.usedNames.add(newName);
    return newName;
  }

  // Generate a random name like "AB123"
  private static generateRandomName(): string {
    const letters = String.fromCharCode(
      65 + Math.floor(Math.random() * 26)
    ) + String.fromCharCode(
      65 + Math.floor(Math.random() * 26)
    );

    const digits = String(Math.floor(Math.random() * 1000)).padStart(3, '0');

    return letters + digits;
  }
}
