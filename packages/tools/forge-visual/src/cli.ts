#!/usr/bin/env node
import { program } from "commander";
import chalk from "chalk";
import { execSync } from "node:child_process";
import { ForgeRunner } from "./runner.js";
import { HistoryStore } from "./history.js";
import { detectProject } from "./detect.js";
import { parseForgeTestJson } from "./parsers/test-json.js";
import { parseGasReport } from "./parsers/gas-report.js";
import { parseLcov } from "./parsers/lcov.js";
import { parseTraceOutput } from "./parsers/trace.js";
import { buildSelectorMap } from "./parsers/diamond-inspect.js";
import { renderTestReport } from "./renderers/test-report.js";
import { renderGasReport } from "./renderers/gas-report.js";
import { renderCoverageReport } from "./renderers/coverage-report.js";
import { renderDashboard } from "./renderers/dashboard.js";
import { renderTraceReport } from "./renderers/trace-report.js";
import { renderUnifiedReport } from "./renderers/unified-report.js";
import type { UnifiedReportData } from "./renderers/unified-report.js";
import type { TestReport, GasReportData, CoverageData, DashboardData } from "./types.js";

const now = () => new Date().toISOString().replace("T", " ").slice(0, 19) + " UTC";

function openBrowser(path: string) {
  try {
    if (process.platform === "darwin") execSync(`open "${path}"`);
    else if (process.platform === "linux") execSync(`xdg-open "${path}"`);
  } catch { /* ignore */ }
}

program
  .name("forge-visual")
  .description("Interactive HTML reports for Foundry test, gas, and coverage data")
  .version("0.1.0")
  .option("-p, --project <path>", "Path to Foundry project root")
  .option("-o, --output <path>", "Output directory for reports", "reports")
  .option("--open", "Open report in browser");

program
  .command("test")
  .description("Run tests and generate interactive HTML report")
  .argument("[forge-args...]", "Additional args for forge test")
  .action(async (forgeArgs: string[]) => {
    const opts = program.opts();
    const runner = new ForgeRunner(opts.project);

    console.log(chalk.dim("Running forge test --json..."));
    const json = runner.runTests(forgeArgs);
    const suites = parseForgeTestJson(json);

    const report = buildTestReport(suites, runner);
    printTestSummary(report);

    const path = renderTestReport(report, opts.output);
    console.log(`${chalk.dim("Report:")} ${chalk.bold(path)}`);
    if (opts.open) openBrowser(path);
  });

program
  .command("gas")
  .description("Run gas analysis with history tracking")
  .option("-t, --threshold <pct>", "Gas increase alert threshold (%)", "10")
  .argument("[forge-args...]", "Additional args for forge test")
  .action(async (forgeArgs: string[], cmdOpts: { threshold: string }) => {
    const opts = program.opts();
    const runner = new ForgeRunner(opts.project);
    const threshold = parseFloat(cmdOpts.threshold);

    console.log(chalk.dim("Running forge test --gas-report..."));
    const text = runner.runGasReport(forgeArgs);
    const contracts = parseGasReport(text);

    const store = new HistoryStore(opts.output);
    const diffs = store.diff(contracts, threshold);
    const alertsCount = diffs.filter((d) => d.alert).length;

    store.append({
      timestamp: new Date().toISOString(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
      reports: contracts,
    });

    const data: GasReportData = {
      contracts,
      diffs,
      hasPrevious: diffs.length > 0,
      threshold,
      alertsCount,
      generatedAt: now(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
    };

    const totalFuncs = contracts.reduce((s, c) => s + c.functions.length, 0);
    console.log(`\n${chalk.cyan.bold("GAS")} ${contracts.length} contracts, ${totalFuncs} functions analyzed`);
    if (alertsCount > 0) {
      console.log(`${chalk.red.bold("ALERT")} ${alertsCount} functions exceeded ${threshold}% threshold`);
    }

    const path = renderGasReport(data, opts.output);
    console.log(`${chalk.dim("Report:")} ${chalk.bold(path)}`);
    if (opts.open) openBrowser(path);
  });

program
  .command("coverage")
  .description("Generate coverage visual dashboard")
  .option("-l, --lcov <path>", "Path to lcov.info")
  .action(async (cmdOpts: { lcov?: string }) => {
    const opts = program.opts();
    const runner = new ForgeRunner(opts.project);

    let lcovContent: string;
    if (cmdOpts.lcov) {
      console.log(`${chalk.dim("Reading:")} ${cmdOpts.lcov}`);
      lcovContent = runner.readLcov(cmdOpts.lcov);
    } else {
      console.log(chalk.dim("Running forge coverage --report lcov --ir-minimum..."));
      lcovContent = runner.runCoverage();
    }

    const data = buildCoverageData(lcovContent, runner);

    console.log(
      `\n${chalk.cyan.bold("COV")} Lines: ${data.linePct.toFixed(1)}% (${data.totalLinesHit}/${data.totalLines}) | Branches: ${data.branchPct.toFixed(1)}% (${data.totalBranchesHit}/${data.totalBranches})`
    );

    const path = renderCoverageReport(data, opts.output);
    console.log(`${chalk.dim("Report:")} ${chalk.bold(path)}`);
    if (opts.open) openBrowser(path);
  });

program
  .command("trace")
  .description("Run tests with -vvvv and generate interactive trace visualization")
  .option("-m, --match <pattern>", "Match test name pattern (--match-test)")
  .argument("[forge-args...]", "Additional args for forge test")
  .action(async (forgeArgs: string[], cmdOpts: { match?: string }) => {
    const opts = program.opts();
    const runner = new ForgeRunner(opts.project);

    const args = ["-vvvv", ...forgeArgs];
    if (cmdOpts.match) args.push("--match-test", cmdOpts.match);

    console.log(chalk.dim(`Running forge test ${args.join(" ")}...`));
    const text = runner.runTrace(args);

    const tests = parseTraceOutput(text);
    const passed = tests.filter((t) => t.status === "pass").length;
    const failed = tests.filter((t) => t.status === "fail").length;

    console.log(
      `\n${chalk.cyan.bold("TRACE")} ${tests.length} tests traced (${chalk.green(passed)} passed${failed > 0 ? `, ${chalk.red(failed)} failed` : ""})`
    );

    const path = renderTraceReport(
      {
        tests,
        generatedAt: now(),
        gitRef: runner.gitRef(),
        gitBranch: runner.gitBranch(),
      },
      opts.output
    );

    console.log(`${chalk.dim("Report:")} ${chalk.bold(path)}`);
    if (opts.open) openBrowser(path);
  });

program
  .command("summary")
  .description("Generate combined dashboard from all data")
  .argument("[forge-args...]", "Additional args for forge test")
  .action(async (forgeArgs: string[]) => {
    const opts = program.opts();
    const runner = new ForgeRunner(opts.project);

    // Tests
    console.log(chalk.dim("Running forge test --json..."));
    const json = runner.runTests(forgeArgs);
    const suites = parseForgeTestJson(json);
    const testReport = buildTestReport(suites, runner);
    printTestSummary(testReport);
    renderTestReport(testReport, opts.output);

    // Gas
    console.log(chalk.dim("Running forge test --gas-report..."));
    const gasText = runner.runGasReport(forgeArgs);
    const contracts = parseGasReport(gasText);
    const store = new HistoryStore(opts.output);
    const diffs = store.diff(contracts, 10);
    store.append({
      timestamp: new Date().toISOString(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
      reports: contracts,
    });
    const gasData: GasReportData = {
      contracts,
      diffs,
      hasPrevious: diffs.length > 0,
      threshold: 10,
      alertsCount: diffs.filter((d) => d.alert).length,
      generatedAt: now(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
    };
    renderGasReport(gasData, opts.output);

    // Coverage
    let coverageData: CoverageData | null = null;
    try {
      console.log(chalk.dim("Running forge coverage --report lcov --ir-minimum..."));
      const lcov = runner.runCoverage();
      coverageData = buildCoverageData(lcov, runner);
      renderCoverageReport(coverageData, opts.output);
    } catch (e) {
      console.log(`${chalk.yellow.bold("WARN")} Coverage skipped: ${e}`);
    }

    // Dashboard
    const dashboard: DashboardData = {
      testReport,
      gasReport: gasData,
      coverage: coverageData,
      generatedAt: now(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
    };

    const path = renderDashboard(dashboard, opts.output);
    console.log(`\n${chalk.green.bold("DONE")} All reports generated in ${opts.output}`);
    console.log(`${chalk.dim("Dashboard:")} ${chalk.bold(path)}`);
    if (opts.open) openBrowser(path);
  });

program
  .command("report")
  .description("Generate unified SPA report with all data (recommended)")
  .option("--skip-coverage", "Skip coverage collection (slow)")
  .option("--skip-trace", "Skip trace collection")
  .option("-m, --match <pattern>", "Match test name pattern for traces")
  .argument("[forge-args...]", "Additional args for forge test")
  .action(async (forgeArgs: string[], cmdOpts: { skipCoverage?: boolean; skipTrace?: boolean; match?: string }) => {
    const opts = program.opts();
    const runner = new ForgeRunner(opts.project);

    // Detect project type
    const project = detectProject(runner.projectRoot);
    const badge = project.type === "diamond" ? chalk.magenta("Diamond")
      : project.type === "uups" ? chalk.blue("UUPS")
      : chalk.gray("Standard");
    console.log(`${chalk.dim("Project:")} ${badge} ${project.facets.length > 0 ? chalk.dim(`(${project.facets.length} facets)`) : ""}`);

    // Tests
    console.log(chalk.dim("Running forge test --json..."));
    const json = runner.runTests(forgeArgs);
    const suites = parseForgeTestJson(json);
    const testReport = buildTestReport(suites, runner);
    printTestSummary(testReport);

    // Gas
    console.log(chalk.dim("Running forge test --gas-report..."));
    const gasText = runner.runGasReport(forgeArgs);
    const contracts = parseGasReport(gasText);
    const store = new HistoryStore(opts.output);
    const diffs = store.diff(contracts, 10);
    store.append({
      timestamp: new Date().toISOString(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
      reports: contracts,
    });
    const gasData: GasReportData = {
      contracts,
      diffs,
      hasPrevious: diffs.length > 0,
      threshold: 10,
      alertsCount: diffs.filter((d) => d.alert).length,
      generatedAt: now(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
    };
    const totalFuncs = contracts.reduce((s, c) => s + c.functions.length, 0);
    console.log(`${chalk.cyan.bold("GAS")} ${contracts.length} contracts, ${totalFuncs} functions`);

    // Coverage
    let coverageData: CoverageData | null = null;
    if (!cmdOpts.skipCoverage) {
      try {
        console.log(chalk.dim("Running forge coverage --report lcov --ir-minimum..."));
        const lcov = runner.runCoverage();
        coverageData = buildCoverageData(lcov, runner);
        console.log(`${chalk.cyan.bold("COV")} Lines: ${coverageData.linePct.toFixed(1)}% | Branches: ${coverageData.branchPct.toFixed(1)}%`);
      } catch (e) {
        console.log(`${chalk.yellow.bold("WARN")} Coverage skipped: ${e}`);
      }
    }

    // Traces
    let traces: import("./parsers/trace.js").TraceTest[] | null = null;
    if (!cmdOpts.skipTrace) {
      try {
        const traceArgs = ["-vvvv", ...forgeArgs];
        if (cmdOpts.match) traceArgs.push("--match-test", cmdOpts.match);
        console.log(chalk.dim(`Running forge test ${traceArgs.join(" ")}...`));
        const traceText = runner.runTrace(traceArgs);
        traces = parseTraceOutput(traceText);
        const tracePassed = traces.filter((t) => t.status === "pass").length;
        console.log(`${chalk.cyan.bold("TRACE")} ${traces.length} tests traced (${tracePassed} passed)`);
      } catch (e) {
        console.log(`${chalk.yellow.bold("WARN")} Traces skipped: ${e}`);
      }
    }

    // Diamond: Selector Map
    let selectorMap: import("./renderers/selector-map.js").SelectorMapData | null = null;
    if (project.type === "diamond" && project.facets.length > 0) {
      try {
        console.log(chalk.dim(`Inspecting ${project.facets.length} facets...`));
        selectorMap = buildSelectorMap(project.facets, runner.projectRoot);
        console.log(`${chalk.magenta.bold("DIAMOND")} ${selectorMap.totalSelectors} selectors, ${selectorMap.collisions.length} collisions`);
      } catch (e) {
        console.log(`${chalk.yellow.bold("WARN")} Selector map skipped: ${e}`);
      }
    }

    // Unified report
    const reportData: UnifiedReportData = {
      project,
      testReport,
      gasReport: gasData,
      coverage: coverageData,
      traces,
      selectorMap,
      generatedAt: now(),
      gitRef: runner.gitRef(),
      gitBranch: runner.gitBranch(),
    };

    const path = renderUnifiedReport(reportData, opts.output);
    console.log(`\n${chalk.green.bold("DONE")} Unified report generated`);
    console.log(`${chalk.dim("Report:")} ${chalk.bold(path)}`);
    if (opts.open) openBrowser(path);
  });

program.parse();

// ── Helpers ──

function buildTestReport(suites: ReturnType<typeof parseForgeTestJson>, runner: ForgeRunner): TestReport {
  const totalTests = suites.reduce((s, suite) => s + suite.tests.length, 0);
  const totalPassed = suites.reduce((s, suite) => s + suite.passed, 0);
  const totalFailed = suites.reduce((s, suite) => s + suite.failed, 0);
  const totalSkipped = suites.reduce((s, suite) => s + suite.skipped, 0);
  const totalDurationMs = suites.reduce((s, suite) => s + suite.durationMs, 0);
  const passRate = totalTests > 0 ? (totalPassed / totalTests) * 100 : 0;

  return {
    suites,
    totalTests,
    totalPassed,
    totalFailed,
    totalSkipped,
    totalDurationMs,
    passRate,
    generatedAt: now(),
    gitRef: runner.gitRef(),
    gitBranch: runner.gitBranch(),
  };
}

function buildCoverageData(lcovContent: string, runner: ForgeRunner): CoverageData {
  const files = parseLcov(lcovContent);
  const totalLinesHit = files.reduce((s, f) => s + f.linesHit, 0);
  const totalLines = files.reduce((s, f) => s + f.linesTotal, 0);
  const totalBranchesHit = files.reduce((s, f) => s + f.branchesHit, 0);
  const totalBranches = files.reduce((s, f) => s + f.branchesTotal, 0);

  return {
    files,
    totalLinesHit,
    totalLines,
    totalBranchesHit,
    totalBranches,
    linePct: totalLines > 0 ? (totalLinesHit / totalLines) * 100 : 100,
    branchPct: totalBranches > 0 ? (totalBranchesHit / totalBranches) * 100 : 100,
    generatedAt: now(),
    gitRef: runner.gitRef(),
    gitBranch: runner.gitBranch(),
  };
}

function printTestSummary(r: TestReport) {
  const label = r.totalFailed > 0 ? chalk.red.bold("FAIL") : chalk.green.bold("PASS");
  console.log(
    `\n${label} ${r.totalTests} tests: ${chalk.green(r.totalPassed)} passed, ${r.totalFailed > 0 ? chalk.red(r.totalFailed) : r.totalFailed} failed, ${r.totalSkipped} skipped (${(r.totalDurationMs / 1000).toFixed(1)}s)`
  );
}
