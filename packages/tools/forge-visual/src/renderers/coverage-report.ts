import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { baseHtml } from "../templates/base.js";
import type { CoverageData, FileCoverage } from "../types.js";

function level(pct: number): string {
  return pct >= 80 ? "high" : pct >= 50 ? "medium" : "low";
}

export function renderCoverageReport(data: CoverageData, outputDir: string): string {
  // Group by category
  const categories = new Map<string, FileCoverage[]>();
  for (const file of data.files) {
    const list = categories.get(file.category) ?? [];
    list.push(file);
    categories.set(file.category, list);
  }

  const categoryBlocks = [...categories.entries()]
    .sort(([a], [b]) => a.localeCompare(b))
    .map(([cat, files]) => {
      const rows = files
        .map(
          (f) => `
        <tr class="border-b border-dark-700/50 hover:bg-dark-700/30 cursor-pointer"
            onclick="this.nextElementSibling.classList.toggle('hidden')">
          <td class="px-4 py-2 font-mono text-xs">${f.facet ?? f.path}</td>
          <td class="px-4 py-2 text-right text-xs">
            <span class="coverage-${f.level}">${f.linesHit}/${f.linesTotal}</span>
          </td>
          <td class="px-4 py-2">
            <div class="w-full bg-dark-700 rounded-full h-2">
              <div class="h-2 rounded-full bar-${f.level}" style="width:${f.linePct.toFixed(1)}%"></div>
            </div>
          </td>
          <td class="px-4 py-2 text-right text-xs text-gray-400">${f.branchesHit}/${f.branchesTotal}</td>
        </tr>
        <tr class="hidden">
          <td colspan="4" class="px-4 py-2 bg-dark-900/50">
            <div class="text-xs font-mono text-gray-500">
              <span class="text-gray-400">Path:</span> ${f.path}<br>
              <span class="text-gray-400">Line coverage:</span> <span class="coverage-${f.level}">${f.linePct.toFixed(1)}%</span><br>
              <span class="text-gray-400">Branch coverage:</span> ${f.branchPct.toFixed(1)}%<br>
              ${f.uncoveredLines.length > 0 ? `<span class="text-gas-high">Uncovered lines:</span> ${f.uncoveredLines.join(", ")}` : ""}
            </div>
          </td>
        </tr>`
        )
        .join("");

      return `
      <div class="glass rounded-xl mb-4 overflow-hidden">
        <div class="px-4 py-3 border-b border-dark-600 flex items-center justify-between">
          <span class="font-semibold text-white">${cat}</span>
          <span class="text-xs text-gray-500">${files.length} files</span>
        </div>
        <table class="w-full text-sm">
          <thead>
            <tr class="text-gray-500 text-xs uppercase border-b border-dark-700">
              <th class="px-4 py-2 text-left">File</th>
              <th class="px-4 py-2 text-right w-32">Lines</th>
              <th class="px-4 py-2 w-48">Coverage</th>
              <th class="px-4 py-2 text-right w-32">Branches</th>
            </tr>
          </thead>
          <tbody>${rows}</tbody>
        </table>
      </div>`;
    })
    .join("");

  const lineLevel = level(data.linePct);
  const branchLevel = level(data.branchPct);

  const content = `
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold coverage-${lineLevel}">${data.linePct.toFixed(1)}%</div>
      <div class="text-sm text-gray-400 mt-1">Line Coverage</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold coverage-${branchLevel}">${data.branchPct.toFixed(1)}%</div>
      <div class="text-sm text-gray-400 mt-1">Branch Coverage</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold text-white">${data.totalLinesHit}/${data.totalLines}</div>
      <div class="text-sm text-gray-400 mt-1">Lines Hit</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold text-white">${data.totalBranchesHit}/${data.totalBranches}</div>
      <div class="text-sm text-gray-400 mt-1">Branches Hit</div>
    </div>
  </div>
  ${categoryBlocks}`;

  const html = baseHtml({
    title: "Coverage Report",
    generatedAt: data.generatedAt,
    gitRef: data.gitRef,
    gitBranch: data.gitBranch,
    content,
  });

  mkdirSync(outputDir, { recursive: true });
  const outputPath = join(outputDir, "coverage-report.html");
  writeFileSync(outputPath, html);
  return outputPath;
}
