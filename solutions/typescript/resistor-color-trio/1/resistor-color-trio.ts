type Color =
  | "black"
  | "brown"
  | "red"
  | "orange"
  | "yellow"
  | "green"
  | "blue"
  | "violet"
  | "grey"
  | "white";

export function decodedResistorValue(colors: [Color, Color, Color]): string {
  const colorCodes: Record<Color, number> = {
    black: 0,
    brown: 1,
    red: 2,
    orange: 3,
    yellow: 4,
    green: 5,
    blue: 6,
    violet: 7,
    grey: 8,
    white: 9,
  };

  const [first, second, third] = colors;
  const baseValue = parseInt(`${colorCodes[first]}${colorCodes[second]}`);
  const multiplier = Math.pow(10, colorCodes[third]);
  const ohmsValue = baseValue * multiplier;

  if (ohmsValue >= 1_000_000_000) {
    return `${ohmsValue / 1_000_000_000} gigaohms`;
  } else if (ohmsValue >= 1_000_000) {
    return `${ohmsValue / 1_000_000} megaohms`;
  } else if (ohmsValue >= 1_000) {
    return `${ohmsValue / 1_000} kiloohms`;
  } else {
    return `${ohmsValue} ohms`;
  }
}
