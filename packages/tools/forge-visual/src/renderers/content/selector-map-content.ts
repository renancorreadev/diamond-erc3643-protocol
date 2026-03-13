import type { SelectorMapData } from "../selector-map.js";

export function renderSelectorMapContent(data: SelectorMapData): string {
  const collisionWarning =
    data.collisions.length > 0
      ? `
    <div class="glass rounded-xl p-4 mb-6 border-l-4 border-gas-high">
      <span class="text-gas-high font-bold">${data.collisions.length} selector collisions detected</span>
      <div class="mt-2 space-y-1">
        ${data.collisions
          .map(
            (c) => `
          <div class="text-xs font-mono">
            <span class="text-gray-400">${c.selector}</span>
            <span class="text-gas-high ml-2">${c.facets.join(", ")}</span>
          </div>`
          )
          .join("")}
      </div>
    </div>`
      : "";

  return `
<div x-data="{
  entries: window.__selectorData || [],
  search: '',
  facetFilter: 'all',
  get facets() {
    return [...new Set(this.entries.map(e => e.facet))].sort();
  },
  get filtered() {
    return this.entries.filter(e => {
      if (this.facetFilter !== 'all' && e.facet !== this.facetFilter) return false;
      if (this.search) {
        const q = this.search.toLowerCase();
        return e.selector.toLowerCase().includes(q) || e.signature.toLowerCase().includes(q) || e.facet.toLowerCase().includes(q);
      }
      return true;
    });
  },
  copySelector(sel) {
    navigator.clipboard.writeText(sel);
  }
}">

  <!-- Summary -->
  <div class="grid grid-cols-3 gap-4 mb-6">
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold text-white">${data.totalSelectors}</div>
      <div class="text-sm text-gray-400">Selectors</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold text-purple-300">${data.facetCount}</div>
      <div class="text-sm text-gray-400">Facets</div>
    </div>
    <div class="glass rounded-xl p-4 text-center">
      <div class="text-3xl font-bold ${data.collisions.length > 0 ? "text-gas-high" : "status-pass"}">${data.collisions.length}</div>
      <div class="text-sm text-gray-400">Collisions</div>
    </div>
  </div>

  ${collisionWarning}

  <!-- Filters -->
  <div class="glass rounded-xl p-4 mb-6 flex flex-wrap gap-4 items-center">
    <input type="text" x-model="search" placeholder="Search selector, function, or facet..."
      class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none flex-1 min-w-[250px]">
    <select x-model="facetFilter" class="bg-dark-700 border border-dark-600 rounded-lg px-3 py-2 text-sm focus:border-accent outline-none">
      <option value="all">All Facets</option>
      <template x-for="f in facets" :key="f">
        <option :value="f" x-text="f"></option>
      </template>
    </select>
    <span class="text-xs text-gray-500" x-text="filtered.length + ' of ' + entries.length + ' selectors'"></span>
  </div>

  <!-- Selector Table -->
  <div class="glass rounded-xl overflow-hidden">
    <table class="w-full text-sm">
      <thead>
        <tr class="text-gray-500 text-xs uppercase border-b border-dark-700">
          <th class="px-4 py-2 text-left w-32">Selector</th>
          <th class="px-4 py-2 text-left">Function</th>
          <th class="px-4 py-2 text-left w-48">Facet</th>
        </tr>
      </thead>
      <tbody>
        <template x-for="e in filtered" :key="e.selector + e.facet">
          <tr class="border-b border-dark-700/50 hover:bg-dark-700/30">
            <td class="px-4 py-2 font-mono text-xs">
              <span class="cursor-pointer hover:text-accent" @click="copySelector(e.selector)" x-text="e.selector" title="Click to copy"></span>
            </td>
            <td class="px-4 py-2 font-mono text-xs text-gray-200" x-text="e.signature"></td>
            <td class="px-4 py-2 text-xs">
              <span class="px-2 py-0.5 rounded bg-purple-500/10 text-purple-300 border border-purple-500/20" x-text="e.facet"></span>
            </td>
          </tr>
        </template>
      </tbody>
    </table>
  </div>
</div>`;
}

export function renderSelectorMapScript(data: SelectorMapData): string {
  return `<script>window.__selectorData = ${JSON.stringify(data.entries)};</script>`;
}
