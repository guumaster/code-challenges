const { word2numEN, word2numES } = require("./word2num.js");
const { num2wordEN, num2wordES } = require("./num2word.js");

const MATH_OP = "math";
const ESP_OP = "esp";
const ENG_OP = "eng";

function word2num(input, lang) {
  switch (lang) {
    case ENG_OP:
      return word2numEN(input);

    case ESP_OP:
      return word2numES(input);

    default:
      throw new Error("unknown language");
  }
}

function num2word(input, lang) {
  switch (lang) {
    case ENG_OP:
      return num2wordEN(input);

    case ESP_OP:
      return num2wordES(input);

    default:
      throw new Error("unknown language");
  }
}

module.exports = {
  MATH_OP,
  ESP_OP,
  ENG_OP,
  word2num,
  num2word,
};
