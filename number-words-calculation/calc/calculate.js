const ERRORS = require("./errors.js");
const { ESP_OP, ENG_OP, MATH_OP } = require("../convert/index.js");
const { parseInput, parseOperation } = require("./parser.js");

const OPERATORS = {
  "+": MATH_OP,
  mas: ESP_OP,
  plus: ENG_OP,
};

const operatorRE = /^(.*)(?:\s*(\+|mas|plus)\s*)(.*)$/i;

function calculateSum(input) {
  input = input.trim();
  if (input == "") {
    throw new Error(ERRORS.EMPTY_INPUT);
  }

  // Single word input
  if (!containsOperator(input)) {
    const [isValid] = parseInput(input);
    if (isValid) {
      return input;
    } else {
      throw new Error(ERRORS.OUT_OF_RANGE);
    }
  }

  // multi word input
  const match = input.match(operatorRE);
  if (!match) {
    throw new Error(ERRORS.INVALID_INPUT);
  }

  const operatorLang = OPERATORS[match[2].trim().toLowerCase()];
  const leftOperand = match[1].trim();
  const rightOperand = match[3].trim();

  if (!leftOperand || !rightOperand) {
    throw new Error(ERRORS.MISSING_OPERAND);
  }

  if (operatorLang === "") {
    throw new Error(ERRORS.INVALID_OPERATOR);
  }

  return parseOperation(leftOperand, rightOperand, operatorLang);
}

const containsOperator = (input) => input.match(/\s*(\+|mas|plus)\s*/);

module.exports = { calculateSum };
