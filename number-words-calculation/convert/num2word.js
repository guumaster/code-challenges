const units_es = [
  "",
  "uno",
  "dos",
  "tres",
  "cuatro",
  "cinco",
  "seis",
  "siete",
  "ocho",
  "nueve",
];
const tens_es = [
  "",
  "",
  "veinte",
  "treinta",
  "cuarenta",
  "cincuenta",
  "sesenta",
  "setenta",
  "ochenta",
  "noventa",
];
const teens_es = [
  "diez",
  "once",
  "doce",
  "trece",
  "catorce",
  "quince",
  "dieciséis",
  "diecisiete",
  "dieciocho",
  "diecinueve",
];
const hundreds_es = [
  "cien",
  "ciento",
  "doscientos",
  "trescientos",
  "cuatrocientos",
  "quinientos",
  "seiscientos",
  "setecientos",
  "ochocientos",
  "novecientos",
];
const units_en = [
  "",
  "one",
  "two",
  "three",
  "four",
  "five",
  "six",
  "seven",
  "eight",
  "nine",
];
const tens_en = [
  "",
  "ten",
  "twenty",
  "thirty",
  "forty",
  "fifty",
  "sixty",
  "seventy",
  "eighty",
  "ninety",
];
const teens_en = [
  "ten",
  "eleven",
  "twelve",
  "thirteen",
  "fourteen",
  "fifteen",
  "sixteen",
  "seventeen",
  "eighteen",
  "nineteen",
];

function num2wordES(num) {
  if (num === 0) {
    return "cero";
  }

  let words = "";

  if (num >= 100) {
    words += hundreds_es[Math.floor(num / 100)];
    num %= 100;
    words += " ";
  }

  if (num > 20 && num < 30) {
    num %= 10;
    if (num > 0) {
      words += "veinti";
    }
  }

  if (num >= 30) {
    words += tens_es[Math.floor(num / 10)];
    num %= 10;
    if (num > 0) {
      words += " y ";
    }
  }

  if (num > 0 && num < 10) {
    words += units_es[num];
  } else if (num >= 10) {
    words += teens_es[num - 10];
  }

  return words.trim();
}

function num2wordEN(num) {
  if (num === 0) {
    return "zero";
  }

  let words = "";

  if (num >= 100) {
    words += units_en[Math.floor(num / 100)] + " hundred";
    num %= 100;
    if (num > 0) {
      words += " ";
    }
  }

  if (num >= 20) {
    words += tens_en[Math.floor(num / 10)];
    num %= 10;
    if (num > 0) {
      words += "-";
    }
  }

  if (num > 0 && num < 10) {
    words += units_en[num];
  } else if (num >= 10) {
    words += teens_en[num - 10];
  }

  return words.trim();
}

module.exports = {
  num2wordEN,
  num2wordES,
};
