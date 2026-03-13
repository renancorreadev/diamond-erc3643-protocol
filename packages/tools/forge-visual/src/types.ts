// ── Test Report ──

export interface TestSuite {
  file: string;
  contract: string;
  facet: string | null;
  durationMs: number;
  tests: TestCase[];
  passed: number;
  failed: number;
  skipped: number;
}

export interface TestCase {
  name: string;
  status: "success" | "failure" | "skipped";
  reason: string | null;
  durationMs: number;
  gas: number | null;
  meanGas: number | null;
  medianGas: number | null;
  kind: TestKind;
}

export type TestKind =
  | { type: "Standard" }
  | { type: "Fuzz"; runs: number }
  | { type: "Invariant"; runs: number; calls: number; reverts: number };

export interface TestReport {
  suites: TestSuite[];
  totalTests: number;
  totalPassed: number;
  totalFailed: number;
  totalSkipped: number;
  totalDurationMs: number;
  passRate: number;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

// ── Gas Report ──

export interface FunctionGas {
  name: string;
  min: number;
  avg: number;
  median: number;
  max: number;
  calls: number;
}

export interface ContractGasReport {
  contract: string;
  functions: FunctionGas[];
}

export interface GasHistoryEntry {
  timestamp: string;
  gitRef: string | null;
  gitBranch: string | null;
  reports: ContractGasReport[];
}

export interface GasHistory {
  entries: GasHistoryEntry[];
}

export interface GasDiff {
  function: string;
  contract: string;
  previousAvg: number;
  currentAvg: number;
  delta: number;
  deltaPct: number;
  alert: boolean;
}

export interface GasReportData {
  contracts: ContractGasReport[];
  diffs: GasDiff[];
  hasPrevious: boolean;
  threshold: number;
  alertsCount: number;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

// ── Coverage ──

export interface FileCoverage {
  path: string;
  facet: string | null;
  category: string;
  linesHit: number;
  linesTotal: number;
  branchesHit: number;
  branchesTotal: number;
  linePct: number;
  branchPct: number;
  uncoveredLines: number[];
  level: "high" | "medium" | "low";
}

export interface CoverageData {
  files: FileCoverage[];
  totalLinesHit: number;
  totalLines: number;
  totalBranchesHit: number;
  totalBranches: number;
  linePct: number;
  branchPct: number;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

// ── Dashboard ──

export interface DashboardData {
  testReport: TestReport | null;
  gasReport: GasReportData | null;
  coverage: CoverageData | null;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}
