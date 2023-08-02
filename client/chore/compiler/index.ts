import Scope from "./runtime/scope";
import Statement from "./statements/models";
import { FunctionBlock } from "./statements";
import errors from "./error";

const compile = (instructions: Statement[]) => {
  const globalScope = new Scope("Global", "Global");

  instructions.forEach((instruction: Statement) => {
    if (instruction.name === "Function") {
      const functionBlock = instruction as FunctionBlock;
      globalScope.addFunction(
        functionBlock.props.id,
        functionBlock.props.type,
        functionBlock
      );
    } else if (instruction.name === "Declaration") {
      instruction.compile(globalScope);
    } else if (instruction.name === "Assignment") {
      instruction.compile(globalScope);
    }
  });
  const mainInstruction = instructions.filter(
    (instruction) => instruction.name === "Main"
  );
  if (mainInstruction.length === 1) {
    const mainIndex = instructions.findIndex(
      (instruction) => instruction.name === "Main"
    );
    if (mainIndex >= 0) instructions[mainIndex].compile(globalScope);
    else
      errors.push({
        type: "Semántico",
        token: { line: 0, col: 0 },
        msg: "No se ha definido el método main",
      });
  } else {
    errors.push({
      type: "Semántico",
      token: { line: 0, col: 0 },
      msg: "No se puede definir más de un método main",
    });
  }
};

export default compile;
