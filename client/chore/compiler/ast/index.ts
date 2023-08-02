import {
  FunctionBlock,
  ExpAssignment,
  FunctionCall,
  CycleControl,
  Declaration,
  Expression,
  Condition,
  Switch,
  Main,
  Assignment,
} from "../statements";
import Statement from "../statements";

const getRandom = () => {
  return Math.round(Math.random() * 10000);
};

let g = "";
let parentNode = "";

export const generateDot = (statements: Statement[]) => {
  parentNode = "Program";
  g += `"${parentNode}" [label="Program"];\n`;
  statements.forEach((statement) => {
    const rand = getRandom();
    graphAStatement(statement, rand, parentNode);
  });
  const grafo = g;
  g = "";
  return grafo;
};

const graphAStatement = (
  statement: Statement,
  rand: number,
  parentNodeG?: string
): void => {
  if (statement instanceof FunctionBlock) {
    graphAFunction(statement, rand, parentNodeG);
  } else if (statement instanceof FunctionCall) {
    graphAFunctionCall(statement, rand, parentNodeG);
  } else if (statement instanceof Assignment) {
    graphAAssignment(statement, rand, parentNodeG);
  } else if (statement instanceof Declaration) {
    graphADeclaration(statement, rand, parentNodeG);
  } else if (statement instanceof CycleControl) {
    graphACycleControl(statement, rand, parentNodeG);
  } else if (statement instanceof Condition) {
    graphACondition(statement, rand, parentNodeG);
  } else if (statement instanceof Switch) {
    graphASwitch(statement, rand, parentNodeG);
  } else if (statement instanceof ExpAssignment) {
    graphAExpAssignment(statement, rand, parentNodeG);
  } else if (statement instanceof Expression) {
    graphAExpression(parentNodeG as string, rand, statement);
  } else if (statement instanceof Main) {
    graphAMain(statement, rand);
  }
};

const graphAExpression = (
  expNode: string,
  rand: number,
  exp?: Expression
): void => {
  if (exp?.props.operator) {
    graphAExpression(expNode, rand, exp.props.left);
    const oper = `"${exp.props.operator}_${rand}"`;
    g += `${oper} [label="${exp.props.operator}"];\n`;
    g += `${expNode} -> ${oper};\n`;
    graphAExpression(expNode, rand, exp.props.right);
  } else if (exp?.props.value) {
    const valRand = getRandom();
    const valNode = `"Value_${rand}_${valRand}"`;
    g += `${valNode} [label="Valor"];\n`;
    const valNodeVal = `"${exp.props.value.props.value}_${rand}_${valRand}"`;
    g += `${valNodeVal} [label="${exp.props.value.props.value}"];\n`;
    g += `${expNode} -> ${valNode};\n`;
    g += `${valNode} -> ${valNodeVal};\n`;
  }
};

const graphAFunction = (
  func: FunctionBlock,
  rand: number,
  parentNodeG?: string
): void => {
  const funcNode = `"${func.name}_${rand}"`;
  g += `${funcNode} [label="Función ${func.props.id}"];\n`;
  g += `${parentNodeG} -> ${funcNode};\n`;
  func.props.content?.forEach((statement) => {
    const rand = getRandom();
    const node = `"${statement.name}_${rand}"`;
    g += `${node} [label="${statement.constructor.name}"];\n`;
    g += `${funcNode} -> ${node};\n`;
    graphAStatement(statement, rand);
  });
  if (func.props.params.length > 0) {
    // Creamos un nodo llamado "Parámetros" y lo conectamos con el nodo de la función para despues conectar los parámetros con este nodo
    const paramsNode = `"Parámetros_${rand}"`;
    g += `${paramsNode} [label="Parámetros"];\n`;
    g += `${funcNode} -> ${paramsNode};\n`;

    func.props.params?.forEach((param) => {
      const rand = getRandom();
      const node = `"${param.type}_${rand}"`;
      g += `${node} [label="Parámetro ${param.id}"];\n`;
      g += `${paramsNode} -> ${node};\n`;
    });
  }
};
const graphAFunctionCall = (
  call: FunctionCall,
  rand: number,
  parentNodeI?: string
): void => {
  // Creamos el nodo de la función a la que se va a llamar
  const callNode = `"${call.name}_${rand}"`;

  // Creamos la etiqueta del nodo de la función a la que se va a llamar
  g += `${callNode} [label="Llamada a la función ${call.props.id}"];\n`;

  if (parentNodeI) {
    g += `${parentNodeI} -> ${callNode};\n`;
  }

  // Si el nodo padre no es nulo, lo conectamos con el nodo de la función a la que se va a llamar
  const callnameNode = `"${call.props?.id + "_" ?? ""}${rand}"`;

  // Creamos la etiqueta del nodo de la función a la que se va a llamar
  g += `${callnameNode} [label="${call.props?.id ?? "Función"}"];\n`;

  // Si el nodo padre no es nulo, lo conectamos con el nodo de la función a la que se va a llamar
  const par = `"${rand}_("`;
  g += `${par} [label="("];\n`;
  g += `${callNode} -> ${par};\n`;
  g += `${callNode} -> ${callnameNode};\n`;

  // Recorremos los parámetros de la función
  call.props.params?.forEach((arg) => {
    graphAStatement(arg, rand, callNode);
  });
  const par2 = `"${rand}_)"`;
  g += `${par2} [label=")"];\n`;
  g += `${callNode} -> ${par2};\n`;
};
const graphAAssignment = (
  assign: Assignment,
  rand: number,
  parentNodeI?: string
): void => {
  // Creamos el nodo de la asignación
  const assignNode = `"${assign.name}_${rand}"`;

  // Creamos la etiqueta del nodo de la asignación
  g += `${assignNode} [label="Asignación a ${assign.id}"];\n`;

  if (parentNodeI) g += `${parentNodeI} -> ${assignNode};\n`;

  const randA = getRandom();
  const node = `"${assign.id}_${randA}"`;
  g += `${node} [label="${assign.id}"];\n`;
  g += `${assignNode} -> ${node};\n`;

  //Creamos un nodo con signo igual
  const equal = `"${rand}_="`;
  g += `${equal} [label="="];\n`;
  g += `${assignNode} -> ${equal};\n`;

  graphAExpression(assignNode, randA, assign.props.exp);
  if (assign.value) graphAStatement(assign.value, rand);
};
const graphADeclaration = (
  dec: Declaration,
  rand: number,
  parentNodeG?: string
): void => {
  // Creamos el nodo de la declaración
  const decNode = `"${dec.name}_${rand}"`;

  // Creamos la etiqueta del nodo de la declaración
  g += `${decNode} [label="Declaración de tipo ${dec.props.type}"];\n`;

  if (parentNodeG) g += `${parentNodeG} -> ${decNode};\n`;

  dec.props.assignments.forEach((assign) => {
    const rand = getRandom();
    const node = `"${assign.constructor.name}_${rand}"`;
    g += `${node} [label="${assign.constructor.name}"];\n`;
    g += `${decNode} -> ${node};\n`;
    graphAStatement(assign, rand, decNode);
  });
};
const graphACycleControl = (
  cycle: CycleControl,
  rand: number,
  parentNodeG?: string
): void => {
  const cycleNode = `"${cycle.name}_${rand}"`;
  g += `${cycleNode} [label="${cycle.name}"];\n`;

  if (parentNodeG) g += `${parentNodeG} -> ${cycleNode};\n`;

  graphAExpression(cycleNode, rand, cycle.props.condition);
  cycle.props.body.forEach((statement) => {
    const rand = getRandom();
    graphAStatement(statement, rand, cycleNode);
  });
};

const graphACondition = (
  condition: Condition,
  rand: number,
  parentNodeG?: string
): void => {
  const conditionNode = `"${condition.name}_${rand}"`;
  g += `${conditionNode} [label="${condition.name}"];\n`;

  if (parentNodeG) g += `${parentNodeG} -> ${conditionNode};\n`;

  graphAExpression(conditionNode, rand, condition.getConditionExp().exp);
  condition.getConditionExp().body.forEach((statement) => {
    const rand = getRandom();
    graphAStatement(statement, rand, conditionNode);
  });
};
const graphASwitch = (
  switchStatement: Switch,
  rand: number,
  parentNodeG?: string
): void => {
  const switchNode = `"${switchStatement.name}_${rand}"`;
  g += `${switchNode} [label="${switchStatement.name}"];\n`;

  if (parentNodeG) g += `${parentNodeG} -> ${switchNode};\n`;

  graphAExpression(switchNode, rand, switchStatement.props.value);

  switchStatement.props.cases?.forEach((caseStatement) => {
    const rand = getRandom();
    const caseNode = `"${caseStatement.constructor.name}_${rand}"`;
    g += `${caseNode} [label="${caseStatement.constructor.name}"];\n`;
    g += `${switchNode} -> ${caseNode};\n`;
    caseStatement.body.forEach((statement) => {
      const rand = getRandom();
      graphAStatement(statement, rand, caseNode);
    });
  });
};

const graphAMain = (main: Main, rand: number): void => {
  // Creamos el nodo main
  const mainNode = `"${main.name}_${rand}"`;
  // Creamos la etiqueta del nodo main
  g += `${mainNode} [label="${main.name}"];\n`;

  const call = main.functionCall;

  // Creamos el nodo de la función a la que se va a llamar
  const callNode = `"${call.name}_${rand}"`;

  // Creamos la etiqueta del nodo de la función a la que se va a llamar
  g += `${callNode} [label="Llamada a la función ${call.props.id}"];\n`;

  const callnameNode = `"${call.props?.id + "_" ?? ""}${rand}"`;
  g += `${callnameNode} [label="${call.props?.id ?? "Función"}"];\n`;
  const par = `"${rand}_("`;
  g += `${par} [label="("];\n`;
  g += `${callNode} -> ${par};\n`;
  g += `${callNode} -> ${callnameNode};\n`;
  call.props.params?.forEach((arg) => {
    const rand = getRandom();
    graphAStatement(arg, rand, callNode);
  });
  const par2 = `"${rand}_)"`;
  g += `${par2} [label=")"];\n`;
  g += `${callNode} -> ${par2};\n`;
  g += `${mainNode} -> ${callNode};\n`;
  g += `${parentNode} -> ${mainNode};\n`;
};
const graphAExpAssignment = (
  exp: ExpAssignment,
  rand: number,
  parentNodeG?: string
): void => {
  const expNode = `"${exp.name}_${rand}"`;
  g += `${expNode} [label="${exp.name}"];\n`;
  if (parentNodeG) g += `${parentNodeG} -> ${expNode};\n`;
  graphAExpression(expNode, rand, exp.props.exp);
};
