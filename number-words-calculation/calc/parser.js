const {
  ESP_OP,
  ENG_OP,
  MATH_OP,
  num2word,
  word2num,
} = require("../convert/index.js");
const ERRORS = require("./errors.js");

function parseInput(input, lang = "all") {
  if (input.match(/[0-9]+/)) {
    const val = parseInt(input, 10);
    return [inRange(val), val];
  }
  const langs = lang !== "all" ? [lang] : [ENG_OP, ESP_OP];

  let val, validFound;
  const someValid = langs.some((lang) => {
    if (validFound) return true;
    try {
      val = word2num(input, lang);
      validFound = true;
      return true;
    } catch (err) {
      return false;
    }
  });

  if (!someValid) {
    throw new Error(ERRORS.INVALID_INPUT);
  }

  return [inRange(val), val];
}

const inRange = (val) => val >= 0 && val < 1000;

function parseOperation(leftOperand, rightOperand, lang) {
  let numLeft, numRight, validLeft, validRight;
  if (lang == MATH_OP) {
    numLeft = parseInt(leftOperand, 10);
    numRight = parseInt(rightOperand, 10);
    if (!inRange(numLeft) || !inRange(numRight)) {
      throw new Error(ERRORS.OUT_OF_RANGE);
    }
  } else {
    [validLeft, numLeft] = parseInput(leftOperand, lang);
    [validRight, numRight] = parseInput(rightOperand, lang);

    if (!validLeft || !validRight) {
      throw new Error(ERRORS.INVALID_INPUT);
    }
  }

  const sum = numLeft + numRight;
  if (sum >= 1000) {
    throw new Error(ERRORS.OUT_OF_RANGE);
  }

  if (lang !== MATH_OP) {
    return num2word(sum, lang);
  }

  return sum.toString();
}

module.exports = {
  parseInput,
  parseOperation,
};
