import { execSync } from "node:child_process";
import { existsSync, readFileSync } from "node:fs";
import { join, resolve } from "node:path";

export class ForgeRunner {
  readonly projectRoot: string;

  constructor(project?: string) {
    this.projectRoot = project
      ? resolve(project)
      : this.findProjectRoot();

    if (!existsSync(join(this.projectRoot, "foundry.toml"))) {
      throw new Error(`foundry.toml not found in ${this.projectRoot}`);
    }
  }

  runTests(extraArgs: string[] = []): string {
    const args = ["test", "--json", ...extraArgs].join(" ");
    return this.exec(`forge ${args}`);
  }

  runGasReport(extraArgs: string[] = []): string {
    const args = ["test", "--gas-report", ...extraArgs].join(" ");
    // Gas report table goes to stdout; merge stderr to capture everything
    return this.execMerged(`forge ${args}`);
  }

  runTrace(extraArgs: string[] = []): string {
    const args = ["test", ...extraArgs].join(" ");
    return this.execMerged(`forge ${args}`);
  }

  runCoverage(): string {
    this.exec("forge coverage --report lcov --ir-minimum");
    const lcovPath = join(this.projectRoot, "lcov.info");
    return readFileSync(lcovPath, "utf-8");
  }

  readLcov(path: string): string {
    return readFileSync(resolve(path), "utf-8");
  }

  gitRef(): string | null {
    try {
      return this.exec("git rev-parse --short HEAD").trim();
    } catch {
      return null;
    }
  }

  gitBranch(): string | null {
    try {
      return this.exec("git rev-parse --abbrev-ref HEAD").trim();
    } catch {
      return null;
    }
  }

  private exec(cmd: string): string {
    return execSync(cmd, {
      cwd: this.projectRoot,
      encoding: "utf-8",
      maxBuffer: 100 * 1024 * 1024,
      stdio: ["pipe", "pipe", "pipe"],
    });
  }

  private execMerged(cmd: string): string {
    return execSync(`${cmd} 2>&1`, {
      cwd: this.projectRoot,
      encoding: "utf-8",
      maxBuffer: 100 * 1024 * 1024,
      shell: true,
      stdio: ["pipe", "pipe", "pipe"],
    });
  }

  private findProjectRoot(): string {
    let dir = process.cwd();
    while (true) {
      if (existsSync(join(dir, "foundry.toml"))) return dir;
      const parent = resolve(dir, "..");
      if (parent === dir) {
        throw new Error("could not find foundry.toml in any parent directory");
      }
      dir = parent;
    }
  }
}
