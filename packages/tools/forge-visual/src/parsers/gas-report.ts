import type { ContractGasReport, FunctionGas } from "../types.js";

export function parseGasReport(text: string): ContractGasReport[] {
  const reports: ContractGasReport[] = [];
  let currentContract: string | null = null;
  let currentFunctions: FunctionGas[] = [];
  let inFunctionSection = false;

  for (const line of text.split("\n")) {
    const trimmed = line.trim();

    if (!trimmed || trimmed.startsWith("+-") || trimmed.startsWith("|--")) {
      continue;
    }

    if (!trimmed.startsWith("|")) {
      inFunctionSection = false;
      continue;
    }

    const cells = trimmed
      .split("|")
      .map((s) => s.trim())
      .filter((s) => s.length > 0);

    if (cells.length === 0) continue;

    // Contract header: "src/.../Facet.sol:FacetName Contract"
    if (cells.length >= 1 && /[Cc]ontract\s*$/.test(cells[0])) {
      if (currentContract && currentFunctions.length > 0) {
        reports.push({ contract: currentContract, functions: currentFunctions });
      }

      const header = cells[0].replace(/\s*[Cc]ontract\s*$/, "").trim();
      const colonIdx = header.lastIndexOf(":");
      currentContract = colonIdx >= 0 ? header.slice(colonIdx + 1) : header;
      currentFunctions = [];
      inFunctionSection = false;
      continue;
    }

    if (cells[0] === "Function Name") {
      inFunctionSection = true;
      continue;
    }

    if (cells[0] === "Deployment Cost" || cells[0] === "Deployment Size") {
      continue;
    }

    if (inFunctionSection && cells.length >= 6 && currentContract) {
      const func = parseFunctionRow(cells);
      if (func) currentFunctions.push(func);
    }
  }

  if (currentContract && currentFunctions.length > 0) {
    reports.push({ contract: currentContract, functions: currentFunctions });
  }

  return reports;
}

function parseFunctionRow(cells: string[]): FunctionGas | null {
  const min = parseInt(cells[1].replace(/,/g, ""), 10);
  const avg = parseInt(cells[2].replace(/,/g, ""), 10);
  const median = parseInt(cells[3].replace(/,/g, ""), 10);
  const max = parseInt(cells[4].replace(/,/g, ""), 10);
  const callCount = parseInt(cells[5].replace(/,/g, ""), 10);

  if ([min, avg, median, max, callCount].some(isNaN)) return null;

  return { name: cells[0], min, avg, median, max, calls: callCount };
}
