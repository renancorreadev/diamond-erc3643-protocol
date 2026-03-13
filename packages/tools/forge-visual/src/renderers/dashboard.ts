import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { baseHtml } from "../templates/base.js";
import type { DashboardData } from "../types.js";

function level(pct: number): string {
  return pct >= 80 ? "high" : pct >= 50 ? "medium" : "low";
}

export function renderDashboard(data: DashboardData, outputDir: string): string {
  const { testReport: t, gasReport: g, coverage: c } = data;

  // Top gas consumers
  const topGas = g
    ? g.contracts
        .flatMap((ct) => ct.functions.map((f) => ({ contract: ct.contract, ...f })))
        .sort((a, b) => b.avg - a.avg)
        .slice(0, 10)
    : [];

  const testCards = t
    ? `
    <div class="glass rounded-xl p-6 text-center">
      <div class="text-4xl font-bold text-white mb-1">${t.totalTests}</div>
      <div class="text-sm text-gray-400">Tests</div>
      <div class="mt-2 text-sm">
        <span class="status-pass">${t.totalPassed} passed</span>
        ${t.totalFailed > 0 ? `<span class="status-fail ml-2">${t.totalFailed} failed</span>` : ""}
      </div>
    </div>
    <div class="glass rounded-xl p-6 text-center">
      <div class="text-4xl font-bold ${t.passRate === 100 ? "status-pass" : t.passRate >= 90 ? "text-yellow-400" : "status-fail"}">${t.passRate.toFixed(1)}%</div>
      <div class="text-sm text-gray-400">Pass Rate</div>
    </div>`
    : "";

  const coverageCards = c
    ? `
    <div class="glass rounded-xl p-6 text-center">
      <div class="text-4xl font-bold coverage-${level(c.linePct)}">${c.linePct.toFixed(1)}%</div>
      <div class="text-sm text-gray-400">Line Coverage</div>
    </div>
    <div class="glass rounded-xl p-6 text-center">
      <div class="text-4xl font-bold coverage-${level(c.branchPct)}">${c.branchPct.toFixed(1)}%</div>
      <div class="text-sm text-gray-400">Branch Coverage</div>
    </div>`
    : "";

  const gasAlerts =
    g && g.alertsCount > 0
      ? `
    <div class="glass rounded-xl p-4 border-l-4 border-gas-high">
      <span class="text-gas-high font-bold">${g.alertsCount} gas alerts</span>
      <span class="text-gray-400 ml-2">— functions exceeded ${g.threshold}% threshold</span>
    </div>`
      : "";

  const gasTable =
    topGas.length > 0
      ? `
    <div class="glass rounded-xl overflow-hidden">
      <div class="px-4 py-3 border-b border-dark-600">
        <span class="font-semibold text-white">Top Gas Consumers (by avg)</span>
      </div>
      <table class="w-full text-sm">
        <thead>
          <tr class="text-gray-500 text-xs uppercase border-b border-dark-700">
            <th class="px-4 py-2 text-left">Contract</th>
            <th class="px-4 py-2 text-left">Function</th>
            <th class="px-4 py-2 text-right">Avg Gas</th>
            <th class="px-4 py-2 text-right">Max Gas</th>
          </tr>
        </thead>
        <tbody>
          ${topGas
            .map(
              (f) => `
          <tr class="border-b border-dark-700/50">
            <td class="px-4 py-2 text-xs text-gray-400">${f.contract}</td>
            <td class="px-4 py-2 font-mono text-xs">${f.name}</td>
            <td class="px-4 py-2 text-right font-mono text-xs gas-heatmap-mid">${f.avg.toLocaleString()}</td>
            <td class="px-4 py-2 text-right font-mono text-xs gas-heatmap-high">${f.max.toLocaleString()}</td>
          </tr>`
            )
            .join("")}
        </tbody>
      </table>
    </div>`
      : "";

  const content = `
  <div class="space-y-8">
    <div class="text-center">
      <h1 class="text-3xl font-bold text-white mb-2">Project Dashboard</h1>
      <p class="text-gray-500">Combined view of tests, gas, and coverage</p>
    </div>
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      ${testCards}
      ${coverageCards}
    </div>
    ${gasAlerts}
    ${gasTable}
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <a href="test-report.html" class="glass rounded-xl p-4 hover:border-accent transition-colors text-center">
        <div class="text-accent font-semibold">Test Report</div>
        <div class="text-xs text-gray-500 mt-1">Full test results with gas heatmap</div>
      </a>
      <a href="gas-report.html" class="glass rounded-xl p-4 hover:border-accent transition-colors text-center">
        <div class="text-accent font-semibold">Gas Report</div>
        <div class="text-xs text-gray-500 mt-1">Per-function gas analysis and diff</div>
      </a>
      <a href="coverage-report.html" class="glass rounded-xl p-4 hover:border-accent transition-colors text-center">
        <div class="text-accent font-semibold">Coverage Report</div>
        <div class="text-xs text-gray-500 mt-1">Line and branch coverage per facet</div>
      </a>
    </div>
  </div>`;

  const html = baseHtml({
    title: "Dashboard",
    generatedAt: data.generatedAt,
    gitRef: data.gitRef,
    gitBranch: data.gitBranch,
    content,
  });

  mkdirSync(outputDir, { recursive: true });
  const outputPath = join(outputDir, "dashboard.html");
  writeFileSync(outputPath, html);
  return outputPath;
}
