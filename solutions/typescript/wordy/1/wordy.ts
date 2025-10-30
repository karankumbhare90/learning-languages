export const answer = (question: string): number => {
  if (typeof question !== 'string') throw new Error('Unknown operation');

  // Ensure format
  if (!question.startsWith('What is') || !question.endsWith('?')) {
    throw new Error('Unknown operation');
  }

  // Clean and prepare the string
  let expression = question
    .replace('What is', '')
    .replace('?', '')
    .trim();

  // Replace verbal math operations
  expression = expression
    .replace(/plus/g, '+')
    .replace(/minus/g, '-')
    .replace(/multiplied by/g, '*')
    .replace(/divided by/g, '/');

  // Split into tokens
  const tokens = expression.split(' ').filter(Boolean);

  // Validate tokens
  const validTokens = /^-?\d+$|[+\-*/]$/;
  if (!tokens.every((t) => validTokens.test(t))) {
    throw new Error('Unknown operation');
  }

  // Must start with a number
  if (isNaN(parseInt(tokens[0]))) {
    throw new Error('Syntax error');
  }

  // Evaluate left-to-right (no operator precedence)
  let result = parseInt(tokens[0]);

  for (let i = 1; i < tokens.length; i += 2) {
    const operator = tokens[i];
    const operand = parseInt(tokens[i + 1]);

    if (isNaN(operand)) throw new Error('Syntax error');

    switch (operator) {
      case '+':
        result += operand;
        break;
      case '-':
        result -= operand;
        break;
      case '*':
        result *= operand;
        break;
      case '/':
        result /= operand;
        break;
      default:
        throw new Error('Unknown operation');
    }
  }

  return result;
};
