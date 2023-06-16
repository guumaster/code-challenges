const { calculateSum } = require("./calculate.js");

describe("Calculate", () => {
  describe("invalid input", () => {
    it("handle empty input", () => {
      expect(() => calculateSum("")).toThrow("empty input");
      expect(() => calculateSum("   ")).toThrow("empty input");
    });
    it("handle partial input", () => {
      expect(() => calculateSum("5 + ")).toThrow("missing operand");
      expect(() => calculateSum("plus one ")).toThrow("missing operand");
      expect(() => calculateSum("ocho mas")).toThrow("missing operand");
    });

    it("handle invalid input range", () => {
      const invalidError = "some number is out of range (0-999)";
      expect(() => calculateSum("1000")).toThrow(invalidError);
      expect(() => calculateSum("-1")).toThrow(invalidError);
      expect(() => calculateSum("novecientos mas cuatrocientos")).toThrow(
        invalidError
      );
    });

    it("handle invalid input strings", () => {
      const invalidError = "invalid string";
      expect(() => calculateSum("minus one")).toThrow(invalidError);
      expect(() => calculateSum("mil")).toThrow(invalidError);
      expect(() => calculateSum("menos uno")).toThrow(invalidError);
      expect(() => calculateSum("one thousand")).toThrow(invalidError);
      expect(() => calculateSum("onehundred")).toThrow(invalidError);
      expect(() => calculateSum("fortyfour")).toThrow(invalidError);
      expect(() => calculateSum("one thousand")).toThrow(invalidError);
    });
  });

  describe("single valid input", () => {
    it("handle single valid number", () => {
      expect(calculateSum("1")).toEqual("1");
      expect(calculateSum("300")).toEqual("300");
    });
    it("handle single valid words", () => {
      expect(calculateSum("zero")).toEqual("zero");
      expect(calculateSum("three hundred")).toEqual("three hundred");
      expect(calculateSum("ciento cincuenta")).toEqual("ciento cincuenta");
    });
    it("handle valid words with 'y'", () => {
      expect(calculateSum("treinta y nueve")).toEqual("treinta y nueve");
      expect(calculateSum("cincuenta y uno")).toEqual("cincuenta y uno");
    });
  });
  describe("handle numbers input", () => {
    it("sum valid numbers", () => {
      expect(calculateSum("1 + 1")).toEqual("2");
      expect(calculateSum("122 + 110")).toEqual("232");
      expect(calculateSum("998 + 1")).toEqual("999");
    });
  });

  describe("handle valid operations input", () => {
    it("sum valid numbers", () => {
      expect(calculateSum("one plus five")).toEqual("six");
      expect(
        calculateSum("one hundred nine plus three hundred seventy-five")
      ).toEqual("four hundred eighty-four");
      expect(calculateSum("novecientos noventa y ocho mas uno")).toEqual(
        "novecientos noventa y nueve"
      );
    });
  });
});
