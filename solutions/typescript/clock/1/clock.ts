export class Clock {

  private minutes : number;
  
  constructor(hour: number, minute: number = 0) {
    this.minutes = ((hour * 60 + minute) % (24 * 60) + (24 * 60)) % (24 * 60)
  }

  public toString(): string {
    const hours = Math.floor(this.minutes / 60)
    const mins = this.minutes % 60
    return `${hours.toString().padStart(2, '0')}:${mins.toString().padStart(2, '0')}`
  }

  public plus(minutes: number): Clock {
    return new Clock(0, this.minutes + minutes);
  }

  public minus(minutes: number): Clock {
    return new Clock(0, this.minutes - minutes);
  }

  public equals(other: Clock): boolean {
    return this.minutes === other.minutes
  }
}
