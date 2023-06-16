const { word2numEN, word2numES } = require("./word2num.js");

describe("Convert to numbers", () => {
  describe("English", () => {
    it("convert basic numbers", () => {
      expect(word2numEN("zero")).toEqual(0);
      expect(word2numEN("one")).toEqual(1);
    });

    it("convert big numbers", () => {
      expect(word2numEN("one hundred and forty")).toEqual(140);
      expect(word2numEN("one hundred forty")).toEqual(140);
      expect(word2numEN("nine hundred seventy-three")).toEqual(973);
    });
  });

  describe("Spanish", () => {
    it("convert small numbers", () => {
      expect(word2numES("cero")).toEqual(0);
      expect(word2numES("catorce")).toEqual(14);
      expect(word2numES("diez y ocho")).toEqual(18);
      expect(word2numES("diez")).toEqual(10);
      expect(word2numES("dieciocho")).toEqual(18);
    });

    it("convert teens numbers", () => {
      expect(word2numES("veinticuatro")).toEqual(24);
      expect(word2numES("veinte y cuatro")).toEqual(24);
    });

    it("convert dozens numbers", () => {
      expect(word2numES("treinta y cinco")).toEqual(35);
    });

    it("convert big numbers", () => {
      expect(word2numES("ciento cincuenta y nueve")).toEqual(159);
      expect(word2numES("novecientos setenta y tres")).toEqual(973);
    });
  });
});
