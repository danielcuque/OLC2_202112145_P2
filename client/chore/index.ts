import { generateDot } from "./compiler/ast";
import Statements from "./compiler/statements/models";
import symbols from "./compiler/symbols";
import errors from "./compiler/error";
import logs from "./compiler/logs";
import compile from "./compiler";
import { TypeWiseParser } from "./grammar/Parser";

export const initApp = (code: string) => {
  let codeStatements: Statements[] = [];
  symbols.length = 0;
  errors.length = 0;
  logs.length = 0;

  let parser = new TypeWiseParser();
  let ast: string | undefined = "";
  parser.parseError = function (err, hash) {
    errors.push({
      msg: `Se esperaba ${hash.expected.join(", ")} pero se encontró ${
        hash.token
      }`,
      type: "Sintáctico",
      token: {
        col: hash.loc.last_column,
        line: hash.loc.first_line,
      },
    });
  };
  try {
    codeStatements = parser.parse(code);
    if (codeStatements.length > 0) {
      compile(codeStatements);
      ast = "digraph G {\n";
      ast += generateDot(codeStatements);
      ast += "}";
    }
  } catch (error) {}
  return {
    symbols,
    logs,
    errors,
    ast,
  };
};

export default {};
