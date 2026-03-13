import type { FileCoverage } from "../types.js";

export function parseLcov(content: string): FileCoverage[] {
  const files: FileCoverage[] = [];
  let currentPath: string | null = null;
  let linesHit = 0;
  let linesTotal = 0;
  let branchesHit = 0;
  let branchesTotal = 0;
  let uncovered: number[] = [];

  for (const line of content.split("\n")) {
    const trimmed = line.trim();

    if (trimmed.startsWith("SF:")) {
      currentPath = trimmed.slice(3);
      linesHit = 0;
      linesTotal = 0;
      branchesHit = 0;
      branchesTotal = 0;
      uncovered = [];
    } else if (trimmed.startsWith("DA:")) {
      const parts = trimmed.slice(3).split(",");
      const lineNum = parseInt(parts[0], 10);
      const count = parseInt(parts[1], 10);
      if (count === 0) uncovered.push(lineNum);
    } else if (trimmed.startsWith("LF:")) {
      linesTotal = parseInt(trimmed.slice(3), 10);
    } else if (trimmed.startsWith("LH:")) {
      linesHit = parseInt(trimmed.slice(3), 10);
    } else if (trimmed.startsWith("BRF:")) {
      branchesTotal = parseInt(trimmed.slice(4), 10);
    } else if (trimmed.startsWith("BRH:")) {
      branchesHit = parseInt(trimmed.slice(4), 10);
    } else if (trimmed === "end_of_record" && currentPath) {
      const linePct = linesTotal > 0 ? (linesHit / linesTotal) * 100 : 100;
      const branchPct = branchesTotal > 0 ? (branchesHit / branchesTotal) * 100 : 100;

      files.push({
        path: currentPath,
        facet: extractFacet(currentPath),
        category: extractCategory(currentPath),
        linesHit,
        linesTotal,
        branchesHit,
        branchesTotal,
        linePct,
        branchPct,
        uncoveredLines: uncovered,
        level: linePct >= 80 ? "high" : linePct >= 50 ? "medium" : "low",
      });

      currentPath = null;
    }
  }

  return files;
}

function extractFacet(path: string): string | null {
  const parts = path.split("/");
  const filename = parts[parts.length - 1];
  return filename?.replace(".sol", "") ?? null;
}

function extractCategory(path: string): string {
  if (path.includes("/facets/")) {
    if (path.includes("/token/")) return "Token";
    if (path.includes("/compliance/")) return "Compliance";
    if (path.includes("/identity/")) return "Identity";
    if (path.includes("/rwa/")) return "RWA";
    if (path.includes("/core/") || path.includes("/DiamondCut") || path.includes("/DiamondLoupe") || path.includes("/Ownership")) return "Core";
    if (path.includes("/security/")) return "Security";
    return "Facets";
  }
  if (path.includes("/compliance/modules/")) return "Compliance Modules";
  if (path.includes("/storage/") || path.includes("/libraries/")) return "Libraries";
  return "Other";
}
