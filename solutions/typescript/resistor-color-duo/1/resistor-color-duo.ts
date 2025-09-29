const COLORS = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
];


export function decodedValue(colors:string[]) : number{
  if(!colors || colors.length < 2){
    throw new Error("At least two colors are required");
  }

  const first = COLORS.indexOf(colors[0].toLowerCase())
  const second = COLORS.indexOf(colors[1].toLowerCase())


  if(first === -1 || second === -1){
    throw new Error("Invalid color provided");
  }

  return first * 10 + second;
}
