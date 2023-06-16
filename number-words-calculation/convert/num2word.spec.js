const { num2wordEN, num2wordES } = require("./num2word.js");

describe("Convert to words", () => {
  describe("English", () => {
    it("convert small numbers", () => {
      expect(num2wordEN(0)).toEqual("zero");
      expect(num2wordEN(8)).toEqual("eight");
      expect(num2wordEN(14)).toEqual("fourteen");
      expect(num2wordEN(18)).toEqual("eighteen");
    });

    it("convert teens numbers", () => {
      expect(num2wordEN(24)).toEqual("twenty-four");
      expect(num2wordEN(35)).toEqual("thirty-five");
    });
    it("convert big numbers", () => {
      expect(num2wordEN(140)).toEqual("one hundred forty");
      expect(num2wordEN(973)).toEqual("nine hundred seventy-three");
    });
  });

  describe("Spanish", () => {
    it("convert small numbers", () => {
      expect(num2wordES(0)).toEqual("cero");
      expect(num2wordES(8)).toEqual("ocho");
      expect(num2wordES(14)).toEqual("catorce");
      expect(num2wordES(18)).toEqual("dieciocho");
    });

    it("convert teens numbers", () => {
      expect(num2wordES(24)).toEqual("veinticuatro");
    });

    it("convert dozens numbers", () => {
      expect(num2wordES(35)).toEqual("treinta y cinco");
      expect(num2wordES(42)).toEqual("cuarenta y dos");
    });

    it("convert big numbers", () => {
      expect(num2wordES(159)).toEqual("ciento cincuenta y nueve");
      expect(num2wordES(973)).toEqual("novecientos setenta y tres");
    });
  });
});
