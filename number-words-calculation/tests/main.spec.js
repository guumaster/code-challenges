const { promisify } = require("util");
const exec = promisify(require("child_process").exec);

describe("main", () => {
  describe("When input is empty", () => {
    it("should return an empty string", async () => {
      const { stdout } = await exec("./calculate");
      expect(stdout).toEqual(expect.stringContaining("empty input"));
    });
  });

  describe("when number is out of range", () => {
    it.each`
      invalidInput
      ${"1000"}
      ${"-1"}
      ${"novecientos mas cuatrocientos"}
    `(
      "should return an error message for input '$invalidInput'",
      async (invalidInput) => {
        const { stdout } = await exec(`./calculate "${invalidInput}"`);
        expect(stdout).toEqual(expect.stringContaining("invalid string"));
      }
    );
  });

  describe("When input string is invalid", () => {
    it.each`
      invalidInput
      ${"onehundred"}
      ${"fortyfour"}
      ${"one thousand"}
      ${"minus one"}
      ${"mil"}
      ${"menos uno"}
    `(
      "should return an error message for '$invalidInput'",
      async (invalidInput) => {
        const { stdout } = await exec(`./calculate "${invalidInput}"`);
        expect(stdout).toEqual(expect.stringContaining("invalid string"));
      }
    );
  });

  describe("when input is valid", () => {
    it.each`
      input                                   | result
      ${"5"}                                  | ${"5"}
      ${"0 + 1"}                              | ${"1"}
      ${"five"}                               | ${"five"}
      ${"zero plus one"}                      | ${"one"}
      ${"nine hundred ninety-eight plus one"} | ${"nine hundred ninety-nine"}
      ${"cinco"}                              | ${"cinco"}
      ${"cero mas uno"}                       | ${"uno"}
      ${"novecientos noventa y ocho mas uno"} | ${"novecientos noventa y nueve"}
    `("should return '$result' for '$input'", async ({ input, result }) => {
      const { stdout } = await exec(`./calculate "${input}"`);
      expect(stdout).toEqual(expect.stringContaining(result));
    });
  });
});
