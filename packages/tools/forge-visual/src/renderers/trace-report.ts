import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { baseHtml } from "../templates/base.js";
import type { TraceTest, TraceNode } from "../parsers/trace.js";
import { simplifyArgs } from "../parsers/trace.js";

interface TraceReportOpts {
  tests: TraceTest[];
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

export function renderTraceReport(opts: TraceReportOpts, outputDir: string): string {
  const testsJson = JSON.stringify(
    opts.tests.map((t) => ({
      name: t.name,
      status: t.status,
      gas: t.gas,
      nodes: flattenNodes(t.nodes),
    }))
  );

  const totalTests = opts.tests.length;
  const passed = opts.tests.filter((t) => t.status === "pass").length;
  const failed = opts.tests.filter((t) => t.status === "fail").length;

  const content = `
<div x-data="{
  tests: [],
  search: '',
  filterKind: 'all',
  hideVM: true,
  hideReturns: true,
  selectedTest: null,
  init() {
    this.tests = window.__traceData;
    if (this.tests.length > 0) this.selectedTest = this.tests[0].name;
  },
  get currentTest() {
    return this.tests.find(t => t.name === this.selectedTest);
  },
  get filteredNodes() {
    if (!this.currentTest) return [];
    return this.currentTest.nodes.filter(n => {
      if (this.hideVM && n.contract === 'VM') return false;
      if (this.hideReturns && (n.kind === 'RETURN' || n.kind === 'STOP')) return false;
      if (this.filterKind !== 'all' && n.kind !== this.filterKind) return false;
      if (this.search && !n.label.toLowerCase().includes(this.search.toLowerCase())) return false;
      return true;
    });
  }
}" x-init="init()" x-cloak>

  <!-- Summary -->
  <div class="grid grid-cols-3 gap-4 mb-6">
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold text-white">${totalTests}</div>
      <div class="text-sm text-gray-400">Traced Tests</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold status-pass">${passed}</div>
      <div class="text-sm text-gray-400">Passed</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold status-fail">${failed}</div>
      <div class="text-sm text-gray-400">Failed</div>
    </div>
  </div>

  <!-- Test Selector -->
  <div class="glass rounded-xl p-4 mb-6 flex flex-wrap gap-4 items-center">
    <select x-model="selectedTest"
      class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none flex-1 min-w-[300px]">
      <template x-for="t in tests" :key="t.name">
        <option :value="t.name" x-text="(t.status === 'pass' ? '✓ ' : '✗ ') + t.name + ' (' + t.gas.toLocaleString() + ' gas)'"></option>
      </template>
    </select>
    <input type="text" x-model="search" placeholder="Search calls..."
      class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none min-w-[200px]">
    <select x-model="filterKind" class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none">
      <option value="all">All Types</option>
      <option value="CALL">CALL</option>
      <option value="DELEGATECALL">DELEGATECALL</option>
      <option value="STATICCALL">STATICCALL</option>
      <option value="EMIT">EMIT</option>
      <option value="REVERT">REVERT</option>
    </select>
    <label class="flex items-center gap-2 text-sm text-gray-400 cursor-pointer">
      <input type="checkbox" x-model="hideVM" class="accent-accent"> Hide VM
    </label>
    <label class="flex items-center gap-2 text-sm text-gray-400 cursor-pointer">
      <input type="checkbox" x-model="hideReturns" class="accent-accent"> Hide Returns
    </label>
  </div>

  <!-- Trace Tree -->
  <div class="glass rounded-xl overflow-hidden">
    <div class="px-4 py-3 border-b border-dark-600 flex items-center justify-between">
      <span class="font-semibold text-white" x-text="selectedTest || 'Select a test'"></span>
      <span class="text-xs text-gray-500" x-text="currentTest ? currentTest.gas.toLocaleString() + ' gas' : ''"></span>
    </div>
    <div class="p-2 font-mono text-xs leading-6 overflow-x-auto">
      <template x-for="(node, idx) in filteredNodes" :key="idx">
        <div class="flex items-start hover:bg-dark-700/30 rounded px-2"
             :style="'padding-left: ' + (node.depth * 20) + 'px'">
          <!-- Kind badge -->
          <span class="inline-block w-24 shrink-0 text-center rounded px-1 mr-2"
            :class="{
              'bg-blue-500/20 text-blue-300': node.kind === 'CALL',
              'bg-purple-500/20 text-purple-300': node.kind === 'DELEGATECALL',
              'bg-green-500/20 text-green-300': node.kind === 'STATICCALL',
              'bg-yellow-500/20 text-yellow-300': node.kind === 'EMIT',
              'bg-red-500/20 text-red-300': node.kind === 'REVERT',
              'bg-gray-500/20 text-gray-400': node.kind === 'RETURN' || node.kind === 'STOP',
              'bg-gray-600/20 text-gray-500': node.kind === 'VM' || node.kind === 'OTHER',
            }"
            x-text="node.kind"></span>

          <!-- Gas -->
          <span class="w-20 shrink-0 text-right mr-3"
            :class="node.gas > 50000 ? 'text-gas-high' : node.gas > 10000 ? 'text-gas-mid' : 'text-gray-500'"
            x-text="node.gas ? node.gas.toLocaleString() : ''"></span>

          <!-- Label -->
          <span :class="node.isRevert ? 'text-red-400' : node.isEmit ? 'text-yellow-300' : 'text-gray-200'">
            <span class="text-gray-500" x-text="node.contract ? node.contract + '.' : ''"></span>
            <span class="text-white font-semibold" x-text="node.func"></span>
            <span class="text-gray-500" x-text="node.shortArgs ? '(' + node.shortArgs + ')' : ''"></span>
            <span class="text-gray-600 ml-2" x-text="node.returnData ? '→ ' + node.returnData : ''"></span>
          </span>
        </div>
      </template>
    </div>
  </div>

  <!-- Gas Flame Chart -->
  <template x-if="currentTest">
    <div class="mt-6">
      <h3 class="text-lg font-semibold text-white mb-3">Gas Flame Chart</h3>
      <div class="glass rounded-xl p-4 space-y-1">
        <template x-for="(node, idx) in currentTest.nodes.filter(n => n.gas > 0 && n.kind !== 'RETURN' && n.kind !== 'STOP' && n.contract !== 'VM')" :key="idx">
          <div class="flex items-center gap-2">
            <span class="text-xs text-gray-500 w-40 truncate text-right" x-text="node.contract + '.' + node.func"></span>
            <div class="flex-1 h-5 bg-dark-700 rounded overflow-hidden">
              <div class="h-full rounded transition-all"
                :class="node.isRevert ? 'bg-red-500/60' : node.kind === 'DELEGATECALL' ? 'bg-purple-500/60' : node.kind === 'STATICCALL' ? 'bg-green-500/60' : 'bg-blue-500/60'"
                :style="'width: ' + Math.max(1, (node.gas / currentTest.gas) * 100) + '%'"></div>
            </div>
            <span class="text-xs text-gray-400 w-16 text-right" x-text="node.gas.toLocaleString()"></span>
          </div>
        </template>
      </div>
    </div>
  </template>
</div>

<script>
window.__traceData = ${testsJson};
</script>`;

  const html = baseHtml({
    title: "Trace Report",
    generatedAt: opts.generatedAt,
    gitRef: opts.gitRef,
    gitBranch: opts.gitBranch,
    content,
  });

  mkdirSync(outputDir, { recursive: true });
  const outputPath = join(outputDir, "trace-report.html");
  writeFileSync(outputPath, html);
  return outputPath;
}

/** Flatten the tree into a flat list with depth info for rendering */
function flattenNodes(nodes: TraceNode[]): FlatNode[] {
  const result: FlatNode[] = [];
  function walk(node: TraceNode) {
    result.push({
      depth: node.depth,
      gas: node.gas,
      kind: node.kind,
      contract: node.contract,
      func: node.func,
      shortArgs: simplifyArgs(node.args).slice(0, 80),
      returnData: node.returnData ? simplifyArgs(node.returnData).slice(0, 60) : null,
      isRevert: node.isRevert,
      isEmit: node.isEmit,
      label: `${node.contract}.${node.func}(${simplifyArgs(node.args)})`,
    });
    for (const child of node.children) {
      walk(child);
    }
  }
  for (const n of nodes) walk(n);
  return result;
}

interface FlatNode {
  depth: number;
  gas: number;
  kind: string;
  contract: string;
  func: string;
  shortArgs: string;
  returnData: string | null;
  isRevert: boolean;
  isEmit: boolean;
  label: string;
}
