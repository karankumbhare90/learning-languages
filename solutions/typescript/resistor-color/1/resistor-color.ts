 
  
  export const COLORS: string[] = [
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

 export const colorCode = (color?: string) => {
   if (!color) {
     throw new Error("Color must be provided");
   }
   return COLORS.indexOf(color.toLowerCase())
  }