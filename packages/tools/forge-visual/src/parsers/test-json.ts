import type { TestCase, TestKind, TestSuite } from "../types.js";

interface RawSuiteResult {
  duration: string | { secs: number; nanos: number };
  test_results: Record<string, RawTestResult>;
}

interface RawTestResult {
  status: string;
  reason: string | null;
  duration: string | { secs: number; nanos: number };
  kind: Record<string, unknown>;
}

export function parseForgeTestJson(jsonStr: string): TestSuite[] {
  const raw: Record<string, RawSuiteResult> = JSON.parse(jsonStr);
  const suites: TestSuite[] = [];

  for (const [key, suiteResult] of Object.entries(raw)) {
    const [file, contract] = parseSuiteKey(key);
    const facet = extractFacet(contract);
    const durationMs = parseDuration(suiteResult.duration);

    const tests: TestCase[] = [];

    for (const [testName, result] of Object.entries(suiteResult.test_results)) {
      const status =
        result.status === "Success"
          ? "success"
          : result.status === "Failure"
            ? "failure"
            : "skipped";

      const { kind, gas, meanGas, medianGas } = parseKindAndGas(result.kind);

      tests.push({
        name: testName,
        status: status as TestCase["status"],
        reason: result.reason,
        durationMs: parseDuration(result.duration),
        gas,
        meanGas,
        medianGas,
        kind,
      });
    }

    tests.sort((a, b) => a.name.localeCompare(b.name));

    const passed = tests.filter((t) => t.status === "success").length;
    const failed = tests.filter((t) => t.status === "failure").length;
    const skipped = tests.filter((t) => t.status === "skipped").length;

    suites.push({
      file,
      contract,
      facet,
      durationMs,
      tests,
      passed,
      failed,
      skipped,
    });
  }

  suites.sort((a, b) => a.contract.localeCompare(b.contract));
  return suites;
}

function parseSuiteKey(key: string): [string, string] {
  const idx = key.lastIndexOf(":");
  if (idx === -1) return [key, key];
  return [key.slice(0, idx), key.slice(idx + 1)];
}

function extractFacet(contract: string): string | null {
  const name = contract.replace(/Tests?$/, "");
  return name || null;
}

function parseDuration(d: string | { secs: number; nanos: number }): number {
  if (typeof d === "object" && d !== null) {
    return (d.secs ?? 0) * 1000 + (d.nanos ?? 0) / 1_000_000;
  }
  if (typeof d === "string") {
    return parseDurationString(d);
  }
  return 0;
}

function parseDurationString(s: string): number {
  let totalMs = 0;
  for (const part of s.split(/\s+/)) {
    const match = part.match(/^([\d.]+)(ns|µs|us|ms|s)$/);
    if (!match) continue;
    const val = parseFloat(match[1]);
    switch (match[2]) {
      case "s":  totalMs += val * 1000;       break;
      case "ms": totalMs += val;              break;
      case "µs":
      case "us": totalMs += val / 1000;       break;
      case "ns": totalMs += val / 1_000_000;  break;
    }
  }
  return totalMs;
}

interface KindResult {
  kind: TestKind;
  gas: number | null;
  meanGas: number | null;
  medianGas: number | null;
}

function parseKindAndGas(kindObj: Record<string, unknown>): KindResult {
  if ("Standard" in kindObj) {
    return {
      kind: { type: "Standard" },
      gas: typeof kindObj.Standard === "number" ? kindObj.Standard : null,
      meanGas: null,
      medianGas: null,
    };
  }

  if ("Fuzz" in kindObj) {
    const fuzz = kindObj.Fuzz as Record<string, unknown>;
    const runs = (fuzz.runs as number) ?? 0;
    const meanGas = (fuzz.mean_gas as number) ?? null;
    const medianGas = (fuzz.median_gas as number) ?? null;
    const firstCaseGas =
      (fuzz.first_case as Record<string, unknown>)?.gas as number | undefined;

    return {
      kind: { type: "Fuzz", runs },
      gas: firstCaseGas ?? meanGas,
      meanGas,
      medianGas,
    };
  }

  if ("Invariant" in kindObj) {
    const inv = kindObj.Invariant as Record<string, unknown>;
    return {
      kind: {
        type: "Invariant",
        runs: (inv.runs as number) ?? 0,
        calls: (inv.calls as number) ?? 0,
        reverts: (inv.reverts as number) ?? 0,
      },
      gas: null,
      meanGas: null,
      medianGas: null,
    };
  }

  return { kind: { type: "Standard" }, gas: null, meanGas: null, medianGas: null };
}
