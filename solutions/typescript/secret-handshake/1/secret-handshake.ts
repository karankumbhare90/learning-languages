export function commands(number:number) {
  const actions = [];

  if (number & 1) actions.push('wink');
  if (number & 2) actions.push('double blink');
  if (number & 4) actions.push('close your eyes');
  if (number & 8) actions.push('jump');

  // If 5th bit (16) is set, reverse the sequence
  if (number & 16) actions.reverse();

  return actions;
}
