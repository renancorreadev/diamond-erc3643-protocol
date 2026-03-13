import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { baseHtml } from "../templates/base.js";
import type { TestReport, TestSuite } from "../types.js";

export function renderTestReport(report: TestReport, outputDir: string): string {
  const suitesJson = JSON.stringify(
    report.suites.map((s) => ({
      file: s.file,
      contract: s.contract,
      facet: s.facet,
      duration_ms: s.durationMs,
      passed: s.passed,
      failed: s.failed,
      skipped: s.skipped,
      tests: s.tests.map((t) => ({
        name: t.name,
        status: t.status,
        reason: t.reason,
        duration_ms: t.durationMs,
        gas: t.gas,
        mean_gas: t.meanGas,
        median_gas: t.medianGas,
        kind: t.kind,
      })),
    }))
  );

  const suiteOptions = report.suites
    .map((s) => `<option value="${s.contract}">${s.contract}</option>`)
    .join("\n");

  const content = `
<div x-data="{
  search: '',
  statusFilter: 'all',
  suiteFilter: 'all',
  sortBy: 'name',
  sortAsc: true,
  get filteredSuites() {
    return suites.filter(s => {
      if (this.suiteFilter !== 'all' && s.contract !== this.suiteFilter) return false;
      return true;
    }).map(s => ({
      ...s,
      tests: s.tests.filter(t => {
        if (this.statusFilter !== 'all' && t.status !== this.statusFilter) return false;
        if (this.search && !t.name.toLowerCase().includes(this.search.toLowerCase())) return false;
        return true;
      }).sort((a, b) => {
        let va = a[this.sortBy], vb = b[this.sortBy];
        if (typeof va === 'string') return this.sortAsc ? va.localeCompare(vb) : vb.localeCompare(va);
        return this.sortAsc ? (va||0) - (vb||0) : (vb||0) - (va||0);
      })
    })).filter(s => s.tests.length > 0);
  }
}" x-cloak>

  <!-- Summary Cards -->
  <div class="grid grid-cols-2 md:grid-cols-5 gap-4 mb-8">
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold text-white">${report.totalTests}</div>
      <div class="text-sm text-gray-400 mt-1">Total Tests</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold status-pass">${report.totalPassed}</div>
      <div class="text-sm text-gray-400 mt-1">Passed</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold status-fail">${report.totalFailed}</div>
      <div class="text-sm text-gray-400 mt-1">Failed</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold status-skip">${report.totalSkipped}</div>
      <div class="text-sm text-gray-400 mt-1">Skipped</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold text-accent">${report.passRate.toFixed(1)}%</div>
      <div class="text-sm text-gray-400 mt-1">Pass Rate</div>
    </div>
  </div>

  <!-- Filters -->
  <div class="glass rounded-xl p-4 mb-6 flex flex-wrap gap-4 items-center">
    <input type="text" x-model="search" placeholder="Search tests..."
      class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none flex-1 min-w-[200px]">
    <select x-model="statusFilter" class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none">
      <option value="all">All Status</option>
      <option value="success">Passed</option>
      <option value="failure">Failed</option>
      <option value="skipped">Skipped</option>
    </select>
    <select x-model="suiteFilter" class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none">
      <option value="all">All Suites</option>
      ${suiteOptions}
    </select>
  </div>

  <!-- Test Suites -->
  <template x-for="suite in filteredSuites" :key="suite.contract">
    <div class="glass rounded-xl mb-4 overflow-hidden">
      <div class="px-4 py-3 border-b border-dark-600 flex items-center justify-between">
        <div class="flex items-center gap-3">
          <span class="font-semibold text-white" x-text="suite.contract"></span>
          <span class="text-xs text-gray-500" x-text="suite.file"></span>
        </div>
        <div class="flex items-center gap-3 text-sm">
          <span class="status-pass" x-text="suite.passed + ' passed'"></span>
          <template x-if="suite.failed > 0">
            <span class="status-fail" x-text="suite.failed + ' failed'"></span>
          </template>
          <span class="text-gray-500" x-text="suite.duration_ms.toFixed(0) + 'ms'"></span>
        </div>
      </div>
      <table class="w-full text-sm">
        <thead>
          <tr class="text-gray-500 text-xs uppercase border-b border-dark-700">
            <th class="px-4 py-2 text-left cursor-pointer hover:text-white" @click="sortBy='name'; sortAsc=!sortAsc">Test</th>
            <th class="px-4 py-2 text-center w-20">Status</th>
            <th class="px-4 py-2 text-right w-24 cursor-pointer hover:text-white" @click="sortBy='gas'; sortAsc=!sortAsc">Gas</th>
            <th class="px-4 py-2 text-right w-24 cursor-pointer hover:text-white" @click="sortBy='duration_ms'; sortAsc=!sortAsc">Duration</th>
            <th class="px-4 py-2 text-center w-20">Kind</th>
          </tr>
        </thead>
        <tbody>
          <template x-for="test in suite.tests" :key="test.name">
            <tr class="border-b border-dark-700/50 hover:bg-dark-700/30">
              <td class="px-4 py-2 font-mono text-xs" x-text="test.name"></td>
              <td class="px-4 py-2 text-center">
                <span x-show="test.status==='success'" class="status-pass">PASS</span>
                <span x-show="test.status==='failure'" class="status-fail">FAIL</span>
                <span x-show="test.status==='skipped'" class="status-skip">SKIP</span>
              </td>
              <td class="px-4 py-2 text-right font-mono text-xs"
                :class="test.gas_level==='low'?'gas-heatmap-low':test.gas_level==='mid'?'gas-heatmap-mid':test.gas_level==='high'?'gas-heatmap-high':''"
                x-text="test.gas ? test.gas.toLocaleString() : '-'"></td>
              <td class="px-4 py-2 text-right text-gray-400 text-xs" x-text="test.duration_ms.toFixed(1)+'ms'"></td>
              <td class="px-4 py-2 text-center">
                <span x-show="test.kind.type==='Standard'" class="text-xs text-gray-500">unit</span>
                <span x-show="test.kind.type==='Fuzz'" class="text-xs text-purple-400" x-text="'fuzz('+test.kind.runs+')'"></span>
                <span x-show="test.kind.type==='Invariant'" class="text-xs text-blue-400" x-text="'inv('+test.kind.runs+')'"></span>
              </td>
            </tr>
          </template>
        </tbody>
      </table>
    </div>
  </template>
</div>

<script>
const suites = ${suitesJson};
const allGas = suites.flatMap(s => s.tests.map(t => t.gas)).filter(g => g != null).sort((a,b) => a-b);
const p25 = allGas[Math.floor(allGas.length*0.25)]||0;
const p75 = allGas[Math.floor(allGas.length*0.75)]||0;
suites.forEach(s => s.tests.forEach(t => {
  if (t.gas == null) t.gas_level = 'none';
  else if (t.gas <= p25) t.gas_level = 'low';
  else if (t.gas <= p75) t.gas_level = 'mid';
  else t.gas_level = 'high';
}));
</script>`;

  const html = baseHtml({
    title: "Test Report",
    generatedAt: report.generatedAt,
    gitRef: report.gitRef,
    gitBranch: report.gitBranch,
    content,
  });

  mkdirSync(outputDir, { recursive: true });
  const outputPath = join(outputDir, "test-report.html");
  writeFileSync(outputPath, html);
  return outputPath;
}
