import { existsSync, readdirSync, readFileSync } from "node:fs";
import { join, basename } from "node:path";

export type ProjectType = "diamond" | "uups" | "simple";

export interface ProjectInfo {
  type: ProjectType;
  facets: string[];
}

const DIAMOND_MARKERS = [
  "DiamondCutFacet.sol",
  "DiamondLoupeFacet.sol",
  "LibDiamond.sol",
];

export function detectProject(projectRoot: string): ProjectInfo {
  const srcDir = findSrcDir(projectRoot);
  if (!srcDir) return { type: "simple", facets: [] };

  // Diamond detection: marker files or facets/ directory
  const allFiles = walkSol(srcDir);
  const names = allFiles.map((f) => basename(f));

  const hasDiamondMarker = DIAMOND_MARKERS.some((m) => names.includes(m));
  const facetsDir = allFiles.filter((f) => f.includes("/facets/"));

  if (hasDiamondMarker || facetsDir.length >= 2) {
    const facets = facetsDir
      .map((f) => basename(f, ".sol"))
      .filter((n) => !n.startsWith("I")); // exclude interfaces
    return { type: "diamond", facets };
  }

  // UUPS detection: grep for UUPSUpgradeable import
  const hasUUPS = allFiles.some((f) => {
    try {
      const content = readFileSync(f, "utf-8");
      return content.includes("UUPSUpgradeable");
    } catch {
      return false;
    }
  });

  if (hasUUPS) return { type: "uups", facets: [] };

  return { type: "simple", facets: [] };
}

function findSrcDir(projectRoot: string): string | null {
  // Check foundry.toml for custom src dir
  const tomlPath = join(projectRoot, "foundry.toml");
  if (existsSync(tomlPath)) {
    try {
      const content = readFileSync(tomlPath, "utf-8");
      const match = content.match(/^\s*src\s*=\s*["']([^"']+)["']/m);
      if (match) {
        const dir = join(projectRoot, match[1]);
        if (existsSync(dir)) return dir;
      }
    } catch { /* ignore */ }
  }

  const defaultSrc = join(projectRoot, "src");
  return existsSync(defaultSrc) ? defaultSrc : null;
}

function walkSol(dir: string): string[] {
  const results: string[] = [];
  try {
    for (const entry of readdirSync(dir, { withFileTypes: true })) {
      const full = join(dir, entry.name);
      if (entry.isDirectory()) {
        results.push(...walkSol(full));
      } else if (entry.name.endsWith(".sol")) {
        results.push(full);
      }
    }
  } catch { /* ignore permission errors */ }
  return results;
}
