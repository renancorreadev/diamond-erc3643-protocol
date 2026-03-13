export interface TraceTest {
  name: string;
  status: "pass" | "fail";
  gas: number;
  nodes: TraceNode[];
}

export interface TraceNode {
  depth: number;
  gas: number;
  kind: "CALL" | "DELEGATECALL" | "STATICCALL" | "CREATE" | "EMIT" | "RETURN" | "REVERT" | "STOP" | "VM" | "OTHER";
  contract: string;
  func: string;
  args: string;
  returnData: string | null;
  isRevert: boolean;
  isEmit: boolean;
  children: TraceNode[];
}

/**
 * Parses forge test -vvvv text output into structured trace data.
 *
 * Line patterns:
 *   [PASS] test_Name() (gas: 12345)
 *   [173600] ContractTest::test_Name()
 *     ├─ [26770] Contract::func(args)
 *     │   ├─ emit EventName(args)
 *     │   └─ ← [Return] value
 *     ├─ [111235] Facet::func(args) [delegatecall]
 *     │   └─ ← [Revert] ErrorName(args)
 */
export function parseTraceOutput(text: string): TraceTest[] {
  const tests: TraceTest[] = [];
  const lines = text.split("\n");
  let i = 0;

  while (i < lines.length) {
    const line = lines[i];

    // Match [PASS] or [FAIL] line
    const statusMatch = line.match(/\[(PASS|FAIL)\]\s+(\S+)\s+\(gas:\s*(\d+)\)/);
    if (statusMatch) {
      const status = statusMatch[1].toLowerCase() as "pass" | "fail";
      const name = statusMatch[2];
      const gas = parseInt(statusMatch[3], 10);

      // Collect trace lines until next test or end
      i++;
      // Skip "Traces:" label
      if (i < lines.length && lines[i].trim() === "Traces:") i++;

      const traceLines: string[] = [];
      while (i < lines.length) {
        const tl = lines[i];
        // Stop at next test result, suite summary, or empty line after traces
        if (tl.match(/^\[(PASS|FAIL)\]/) || tl.match(/^Suite result:/)) break;
        if (tl.match(/^Ran \d+ test/)) break;
        traceLines.push(tl);
        i++;
      }

      const nodes = parseTraceLines(traceLines);
      tests.push({ name, status, gas, nodes });
    } else {
      i++;
    }
  }

  return tests;
}

function parseTraceLines(lines: string[]): TraceNode[] {
  const rootNodes: TraceNode[] = [];
  const stack: { depth: number; node: TraceNode }[] = [];

  for (const line of lines) {
    // Skip empty lines
    if (!line.trim()) continue;

    // Count depth by tree characters (│ ├ └)
    const depth = countDepth(line);

    // Strip tree prefix to get content
    const content = stripTreePrefix(line);
    if (!content) continue;

    // Parse the content
    const node = parseTraceLine(content, depth);
    if (!node) continue;

    // Find parent
    while (stack.length > 0 && stack[stack.length - 1].depth >= depth) {
      stack.pop();
    }

    if (stack.length > 0) {
      stack[stack.length - 1].node.children.push(node);
    } else {
      rootNodes.push(node);
    }

    stack.push({ depth, node });
  }

  return rootNodes;
}

function countDepth(line: string): number {
  let depth = 0;
  for (let i = 0; i < line.length; i++) {
    const ch = line[i];
    if (ch === " ") continue;
    if (ch === "│" || ch === "├" || ch === "└" || ch === "─") {
      depth++;
    } else {
      break;
    }
  }
  // Normalize: each level is roughly 4 chars of indent
  return Math.floor(depth / 2);
}

function stripTreePrefix(line: string): string {
  return line.replace(/^[\s│├└─]+/, "").trim();
}

function parseTraceLine(content: string, depth: number): TraceNode | null {
  // Return/Revert/Stop line: ← [Return] value / ← [Revert] Error(args) / ← [Stop]
  const returnMatch = content.match(/^←\s+\[(Return|Revert|Stop)\]\s*(.*)?$/);
  if (returnMatch) {
    const isRevert = returnMatch[1] === "Revert";
    return {
      depth,
      gas: 0,
      kind: isRevert ? "REVERT" : returnMatch[1] === "Stop" ? "STOP" : "RETURN",
      contract: "",
      func: returnMatch[1],
      args: "",
      returnData: returnMatch[2]?.trim() || null,
      isRevert,
      isEmit: false,
      children: [],
    };
  }

  // Emit line: emit EventName(args)
  const emitMatch = content.match(/^emit\s+(\w+)\((.*)?\)$/);
  if (emitMatch) {
    return {
      depth,
      gas: 0,
      kind: "EMIT",
      contract: "",
      func: emitMatch[1],
      args: emitMatch[2] || "",
      returnData: null,
      isRevert: false,
      isEmit: true,
      children: [],
    };
  }

  // Call line: [gas] Contract::func(args) [delegatecall]?
  const callMatch = content.match(/^\[(\d+)\]\s+(\w+)::(\w+)\((.*)?\)(?:\s+\[(\w+)\])?$/);
  if (callMatch) {
    const callKind = callMatch[5]?.toUpperCase() || "CALL";
    return {
      depth,
      gas: parseInt(callMatch[1], 10),
      kind: callKind as TraceNode["kind"],
      contract: callMatch[2],
      func: callMatch[3],
      args: callMatch[4] || "",
      returnData: null,
      isRevert: false,
      isEmit: false,
      children: [],
    };
  }

  // Fallback call: [gas] Contract::fallback(args) [staticcall]?
  const fallbackMatch = content.match(/^\[(\d+)\]\s+(\w+)::fallback\((.*)?\)(?:\s+\[(\w+)\])?$/);
  if (fallbackMatch) {
    const callKind = fallbackMatch[4]?.toUpperCase() || "CALL";
    return {
      depth,
      gas: parseInt(fallbackMatch[1], 10),
      kind: callKind as TraceNode["kind"],
      contract: fallbackMatch[2],
      func: "fallback",
      args: fallbackMatch[3] || "",
      returnData: null,
      isRevert: false,
      isEmit: false,
      children: [],
    };
  }

  // Top-level test: [gas] TestContract::testName()
  const topMatch = content.match(/^\[(\d+)\]\s+(\w+)::(\w+)\(\)$/);
  if (topMatch) {
    return {
      depth,
      gas: parseInt(topMatch[1], 10),
      kind: "CALL",
      contract: topMatch[2],
      func: topMatch[3],
      args: "",
      returnData: null,
      isRevert: false,
      isEmit: false,
      children: [],
    };
  }

  return null;
}

/**
 * Simplifies addresses in trace args for readability.
 * "alice: [0x328809Bc894f92807417D2dAD6b7C998c1aFdac6]" → "alice"
 */
export function simplifyArgs(args: string): string {
  return args.replace(/(\w+):\s*\[0x[a-fA-F0-9]+\]/g, "$1");
}
