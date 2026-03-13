import type { CoverageData, FileCoverage } from "../../types.js";
import type { ProjectInfo } from "../../detect.js";

function level(pct: number): string {
  return pct >= 80 ? "high" : pct >= 50 ? "medium" : "low";
}

export function renderCoverageContent(data: CoverageData, project: ProjectInfo): string {
  // Group by category
  const categories = new Map<string, FileCoverage[]>();
  for (const file of data.files) {
    const list = categories.get(file.category) ?? [];
    list.push(file);
    categories.set(file.category, list);
  }

  // Facet heatmap for Diamond projects
  const facetHeatmap =
    project.type === "diamond"
      ? renderFacetHeatmap(data.files.filter((f) => f.facet))
      : "";

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

  return `
  <div>
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
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
    ${facetHeatmap}
    ${categoryBlocks}
  </div>`;
}

function renderFacetHeatmap(facetFiles: FileCoverage[]): string {
  if (facetFiles.length === 0) return "";

  const tiles = facetFiles
    .sort((a, b) => (b.linePct || 0) - (a.linePct || 0))
    .map(
      (f) => `
      <div class="rounded-lg p-3 text-center border border-dark-600 hover:border-accent/50 transition-colors"
        style="background: rgba(${f.linePct >= 80 ? "16,185,129" : f.linePct >= 50 ? "245,158,11" : "239,68,68"}, ${Math.max(0.05, f.linePct / 300)})">
        <div class="text-xs font-mono text-white truncate">${f.facet}</div>
        <div class="text-lg font-bold coverage-${f.level} mt-1">${f.linePct.toFixed(0)}%</div>
        <div class="text-xs text-gray-500">${f.branchPct.toFixed(0)}% branches</div>
      </div>`
    )
    .join("");

  return `
    <div class="mb-6">
      <h3 class="text-sm font-semibold text-white mb-3">Facet Coverage Heatmap</h3>
      <div class="grid grid-cols-3 md:grid-cols-5 lg:grid-cols-7 gap-2">
        ${tiles}
      </div>
    </div>`;
}
