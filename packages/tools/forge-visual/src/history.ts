import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import type { ContractGasReport, GasDiff, GasHistory, GasHistoryEntry } from "./types.js";

const MAX_ENTRIES = 50;

export class HistoryStore {
  private readonly path: string;

  constructor(outputDir: string) {
    this.path = join(outputDir, ".forge-visual", "gas-history.json");
  }

  load(): GasHistory {
    if (!existsSync(this.path)) {
      return { entries: [] };
    }
    return JSON.parse(readFileSync(this.path, "utf-8"));
  }

  append(entry: GasHistoryEntry): void {
    const history = this.load();
    history.entries.push(entry);

    if (history.entries.length > MAX_ENTRIES) {
      history.entries.splice(0, history.entries.length - MAX_ENTRIES);
    }

    const dir = dirname(this.path);
    mkdirSync(dir, { recursive: true });
    writeFileSync(this.path, JSON.stringify(history, null, 2));
  }

  latest(): GasHistoryEntry | null {
    const history = this.load();
    return history.entries[history.entries.length - 1] ?? null;
  }

  diff(current: ContractGasReport[], threshold: number): GasDiff[] {
    const previous = this.latest()?.reports;
    if (!previous) return [];

    const diffs: GasDiff[] = [];

    for (const curr of current) {
      for (const func of curr.functions) {
        const prevContract = previous.find((c) => c.contract === curr.contract);
        const prevFunc = prevContract?.functions.find((f) => f.name === func.name);

        if (prevFunc) {
          const delta = func.avg - prevFunc.avg;
          const deltaPct = prevFunc.avg > 0 ? (delta / prevFunc.avg) * 100 : 0;

          diffs.push({
            function: func.name,
            contract: curr.contract,
            previousAvg: prevFunc.avg,
            currentAvg: func.avg,
            delta,
            deltaPct,
            alert: deltaPct > threshold,
          });
        }
      }
    }

    diffs.sort((a, b) => b.deltaPct - a.deltaPct);
    return diffs;
  }
}
