import { execSync } from "node:child_process";
import type { SelectorEntry, SelectorMapData } from "../renderers/selector-map.js";

/**
 * Runs `forge inspect <Facet> methods` for each facet and builds a selector map.
 */
export function buildSelectorMap(facets: string[], projectRoot: string): SelectorMapData {
  const entries: SelectorEntry[] = [];

  for (const facet of facets) {
    try {
      const output = execSync(`forge inspect ${facet} methods --json`, {
        cwd: projectRoot,
        encoding: "utf-8",
        maxBuffer: 10 * 1024 * 1024,
        stdio: ["pipe", "pipe", "pipe"],
      });

      // Output is JSON: { "functionName(args)": "0x12345678", ... }
      const methods = JSON.parse(output) as Record<string, string>;
      for (const [signature, selector] of Object.entries(methods)) {
        entries.push({
          selector: selector.startsWith("0x") ? selector : `0x${selector}`,
          signature,
          facet,
        });
      }
    } catch {
      // Skip facets that fail to compile/inspect
    }
  }

  // Sort by facet then selector
  entries.sort((a, b) => a.facet.localeCompare(b.facet) || a.selector.localeCompare(b.selector));

  // Detect collisions
  const selectorMap = new Map<string, string[]>();
  for (const e of entries) {
    const list = selectorMap.get(e.selector) ?? [];
    if (!list.includes(e.facet)) list.push(e.facet);
    selectorMap.set(e.selector, list);
  }

  const collisions = [...selectorMap.entries()]
    .filter(([, facets]) => facets.length > 1)
    .map(([selector, facets]) => ({ selector, facets }));

  const facetCount = new Set(entries.map((e) => e.facet)).size;

  return {
    entries,
    collisions,
    facetCount,
    totalSelectors: entries.length,
  };
}
