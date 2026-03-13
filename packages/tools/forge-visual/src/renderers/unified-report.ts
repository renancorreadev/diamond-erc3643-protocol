import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { spaHtml } from "../templates/spa-shell.js";
import type { SpaTab } from "../templates/spa-shell.js";
import type { ProjectInfo } from "../detect.js";
import type { TestReport, GasReportData, CoverageData } from "../types.js";
import type { TraceTest } from "../parsers/trace.js";
import { simplifyArgs } from "../parsers/trace.js";
import type { SelectorMapData } from "./selector-map.js";
import { renderDashboardContent, renderDashboardScript } from "./content/dashboard-content.js";
import { renderTestContent, renderTestScript } from "./content/test-content.js";
import { renderGasContent } from "./content/gas-content.js";
import { renderCoverageContent } from "./content/coverage-content.js";
import { renderTraceContent, renderTraceScript, flattenTraceNodes } from "./content/trace-content.js";
import { renderSelectorMapContent, renderSelectorMapScript } from "./content/selector-map-content.js";

export interface UnifiedReportData {
  project: ProjectInfo;
  testReport: TestReport | null;
  gasReport: GasReportData | null;
  coverage: CoverageData | null;
  traces: TraceTest[] | null;
  selectorMap: SelectorMapData | null;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

export function renderUnifiedReport(data: UnifiedReportData, outputDir: string): string {
  const tabs: SpaTab[] = [];
  const scripts: string[] = [];

  // Dashboard (always)
  tabs.push({
    id: "dashboard",
    label: "Dashboard",
    icon: "\u25A6",
    content: renderDashboardContent(data),
  });
  scripts.push(renderDashboardScript(data));

  // Tests
  if (data.testReport) {
    tabs.push({
      id: "tests",
      label: "Tests",
      icon: "\u2713",
      content: renderTestContent(data.testReport),
    });
    scripts.push(renderTestScript(data.testReport));
  }

  // Gas
  if (data.gasReport) {
    tabs.push({
      id: "gas",
      label: "Gas",
      icon: "\u26A1",
      content: renderGasContent(data.gasReport),
    });
  }

  // Coverage
  if (data.coverage) {
    tabs.push({
      id: "coverage",
      label: "Coverage",
      icon: "\u25CE",
      content: renderCoverageContent(data.coverage, data.project),
    });
  }

  // Traces
  if (data.traces && data.traces.length > 0) {
    tabs.push({
      id: "traces",
      label: "Traces",
      icon: "\u21C9",
      content: renderTraceContent(data.traces),
    });
    scripts.push(renderTraceScript(data.traces));
  }

  // Diamond: Selector Map
  if (data.selectorMap && data.project.type === "diamond") {
    tabs.push({
      id: "selectors",
      label: "Selectors",
      icon: "\u2394",
      content: renderSelectorMapContent(data.selectorMap),
      diamond: true,
    });
    scripts.push(renderSelectorMapScript(data.selectorMap));
  }

  const html = spaHtml({
    projectType: data.project.type,
    facets: data.project.facets,
    generatedAt: data.generatedAt,
    gitRef: data.gitRef,
    gitBranch: data.gitBranch,
    tabs,
    scripts: scripts.join("\n"),
  });

  mkdirSync(outputDir, { recursive: true });
  const outputPath = join(outputDir, "report.html");
  writeFileSync(outputPath, html);
  return outputPath;
}
