const map_small_en = {
  zero: 0,
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
  ten: 10,
  eleven: 11,
  twelve: 12,
  thirteen: 13,
  fourteen: 14,
  fifteen: 15,
  sixteen: 16,
  seventeen: 17,
  eighteen: 18,
  nineteen: 19,
  twenty: 20,
  thirty: 30,
  forty: 40,
  fifty: 50,
  sixty: 60,
  seventy: 70,
  eighty: 80,
  ninety: 90,
};

const map_small_es = {
  cero: 0,
  uno: 1,
  dos: 2,
  tres: 3,
  cuatro: 4,
  cinco: 5,
  seis: 6,
  siete: 7,
  ocho: 8,
  nueve: 9,
  diez: 10,
  once: 11,
  doce: 12,
  trece: 13,
  catorce: 14,
  quince: 15,
  dieciséis: 16,
  diecisiete: 17,
  dieciocho: 18,
  diecinueve: 19,
  veinte: 20,
  treinta: 30,
  cuarenta: 40,
  cincuenta: 50,
  sesenta: 60,
  setenta: 70,
  ochenta: 80,
  noventa: 90,
};

const map_hundreds_es = {
  cien: 100,
  ciento: 100,
  doscientos: 200,
  trescientos: 300,
  cuatrocientos: 400,
  quinientos: 500,
  seiscientos: 600,
  setecientos: 700,
  ochocientos: 800,
  novecientos: 900,
};

function word2numES(input) {
  input = input.replace(/\sy\s/, " ").split(" ");
  let n = 0;
  let g = 0;

  for (let word of input) {
    let x = map_small_es[word];
    if (x === undefined && word.match(/veinti/)) {
      g = 20;
      x = map_small_es[word.replace(/veinti/, "")];
    }
    if (x === undefined && word.match(/dieci/)) {
      g = 10;
      x = map_small_es[word.replace(/dieci/, "")];
    }
    if (x !== undefined) {
      g += x;
    } else if (map_hundreds_es[word] !== undefined) {
      n += map_hundreds_es[word];
      g = 0;
    } else {
      throw new Error("Unknown number: " + word);
    }
  }

  return n + g;
}

function word2numEN(input) {
  input = input.replace(/\sand\s/, " ").split(/[\s-]+/);
  let n = 0;
  let g = 0;

  for (let word of input) {
    let x = map_small_en[word];
    if (x !== undefined) {
      g += x;
    } else if (word === "hundred" && g !== 0) {
      g *= 100;
    } else {
      throw new Error("Unknown number: " + word);
    }
  }

  return n + g;
}

module.exports = {
  word2numEN,
  word2numES,
};
