export function hey(message: unknown): string {
  if (typeof message !== 'string') return "Whatever.";

  const trimmed = message.trim();

  // Silence
  if (trimmed.length === 0) {
    return "Fine. Be that way!";
  }

  const isQuestion = trimmed.endsWith('?');
  const hasLetters = /[a-zA-Z]/.test(trimmed);
  const isYelling = hasLetters && trimmed === trimmed.toUpperCase();

  if (isYelling && isQuestion) {
    return "Calm down, I know what I'm doing!";
  } else if (isYelling) {
    return "Whoa, chill out!";
  } else if (isQuestion) {
    return "Sure.";
  } else {
    return "Whatever.";
  }
}
