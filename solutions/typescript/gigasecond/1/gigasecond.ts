export class Gigasecond {
  private startDate: Date;

  constructor(date: Date) {
    this.startDate = date;
  }

  public date(): Date {
    const GIGASECOND_MS = 1e9 * 1000; // 1 gigasecond = 1e9 seconds
    return new Date(this.startDate.getTime() + GIGASECOND_MS);
  }
}
