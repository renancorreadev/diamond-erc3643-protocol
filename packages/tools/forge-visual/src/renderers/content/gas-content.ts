import type { GasReportData } from "../../types.js";

export function renderGasContent(data: GasReportData): string {
  const alertRows = data.diffs
    .filter((d) => d.alert)
    .map(
      (d) => `
      <div class="flex items-center justify-between text-sm py-1 border-b border-dark-700/50">
        <span class="font-mono text-xs">${d.contract}.${d.function}</span>
        <div class="flex items-center gap-4">
          <span class="text-gray-500">${d.previousAvg} &rarr; ${d.currentAvg}</span>
          <span class="text-gas-high font-bold">+${d.deltaPct.toFixed(1)}%</span>
        </div>
      </div>`
    )
    .join("");

  const contractTables = data.contracts
    .map((c) => {
      const rows = c.functions
        .map(
          (f) => `
        <tr class="border-b border-dark-700/50 hover:bg-dark-700/30">
          <td class="px-4 py-2 font-mono text-xs">${f.name}</td>
          <td class="px-4 py-2 text-right font-mono text-xs gas-heatmap-low">${f.min.toLocaleString()}</td>
          <td class="px-4 py-2 text-right font-mono text-xs gas-heatmap-mid">${f.avg.toLocaleString()}</td>
          <td class="px-4 py-2 text-right font-mono text-xs">${f.median.toLocaleString()}</td>
          <td class="px-4 py-2 text-right font-mono text-xs gas-heatmap-high">${f.max.toLocaleString()}</td>
          <td class="px-4 py-2 text-right text-gray-400 text-xs">${f.calls}</td>
        </tr>`
        )
        .join("");

      return `
      <div class="glass rounded-xl mb-4 overflow-hidden">
        <div class="px-4 py-3 border-b border-dark-600">
          <span class="font-semibold text-white">${c.contract}</span>
          <span class="text-xs text-gray-500 ml-2">${c.functions.length} functions</span>
        </div>
        <table class="w-full text-sm">
          <thead>
            <tr class="text-gray-500 text-xs uppercase border-b border-dark-700">
              <th class="px-4 py-2 text-left">Function</th>
              <th class="px-4 py-2 text-right">Min</th>
              <th class="px-4 py-2 text-right">Avg</th>
              <th class="px-4 py-2 text-right">Median</th>
              <th class="px-4 py-2 text-right">Max</th>
              <th class="px-4 py-2 text-right">Calls</th>
            </tr>
          </thead>
          <tbody>${rows}</tbody>
        </table>
      </div>`;
    })
    .join("");

  const diffTable = data.hasPrevious
    ? `
    <h2 class="text-xl font-bold text-white mt-8 mb-4">Gas Diff vs Previous Run</h2>
    <div class="glass rounded-xl overflow-hidden">
      <table class="w-full text-sm">
        <thead>
          <tr class="text-gray-500 text-xs uppercase border-b border-dark-700">
            <th class="px-4 py-2 text-left">Contract</th>
            <th class="px-4 py-2 text-left">Function</th>
            <th class="px-4 py-2 text-right">Previous</th>
            <th class="px-4 py-2 text-right">Current</th>
            <th class="px-4 py-2 text-right">Delta</th>
          </tr>
        </thead>
        <tbody>
          ${data.diffs
            .map(
              (d) => `
          <tr class="border-b border-dark-700/50 hover:bg-dark-700/30">
            <td class="px-4 py-2 text-xs text-gray-400">${d.contract}</td>
            <td class="px-4 py-2 font-mono text-xs">${d.function}</td>
            <td class="px-4 py-2 text-right font-mono text-xs text-gray-400">${d.previousAvg.toLocaleString()}</td>
            <td class="px-4 py-2 text-right font-mono text-xs">${d.currentAvg.toLocaleString()}</td>
            <td class="px-4 py-2 text-right font-mono text-xs font-bold ${d.deltaPct > 0 ? "text-gas-high" : d.deltaPct < 0 ? "text-gas-low" : "text-gray-400"}">
              ${d.deltaPct > 0 ? "+" : ""}${d.deltaPct.toFixed(1)}%
            </td>
          </tr>`
            )
            .join("")}
        </tbody>
      </table>
    </div>`
    : "";

  const alertBanner =
    data.alertsCount > 0
      ? `
    <div class="glass rounded-xl p-4 mb-6 border-l-4 border-gas-high" x-data="{ show: false }">
      <div class="flex items-center gap-2">
        <span class="text-gas-high font-bold text-lg">${data.alertsCount}</span>
        <span class="text-gray-300">functions exceeded ${data.threshold}% gas increase threshold</span>
        <button @click="show=!show" class="ml-auto text-sm text-accent hover:underline" x-text="show?'Hide':'Show'"></button>
      </div>
      <div x-show="show" x-transition class="mt-3 space-y-1">${alertRows}</div>
    </div>`
      : "";

  return `<div>${alertBanner}${contractTables}${diffTable}</div>`;
}
