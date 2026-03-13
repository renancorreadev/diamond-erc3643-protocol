import type { ProjectType } from "../detect.js";

export interface SpaShellOpts {
  projectType: ProjectType;
  facets: string[];
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
  tabs: SpaTab[];
  scripts: string;
}

export interface SpaTab {
  id: string;
  label: string;
  icon: string;
  content: string;
  diamond?: boolean; // only shown for Diamond projects
}

const TYPE_BADGES: Record<ProjectType, { label: string; color: string }> = {
  diamond: { label: "Diamond", color: "bg-purple-500/20 text-purple-300 border-purple-500/30" },
  uups: { label: "UUPS", color: "bg-blue-500/20 text-blue-300 border-blue-500/30" },
  simple: { label: "Standard", color: "bg-gray-500/20 text-gray-300 border-gray-500/30" },
};

export function spaHtml(opts: SpaShellOpts): string {
  const badge = TYPE_BADGES[opts.projectType];
  const visibleTabs = opts.tabs.filter(
    (t) => !t.diamond || opts.projectType === "diamond"
  );

  const sidebarLinks = visibleTabs
    .map(
      (t) => `
        <button @click="navigate('${t.id}')"
          :class="tab === '${t.id}' ? 'bg-accent/10 text-accent border-l-2 border-accent' : 'text-gray-400 hover:text-gray-200 hover:bg-dark-700/50 border-l-2 border-transparent'"
          class="w-full flex items-center gap-3 px-4 py-2.5 text-sm transition-all">
          <span class="text-base">${t.icon}</span>
          <span x-show="!collapsed">${t.label}</span>
        </button>`
    )
    .join("\n");

  const sections = visibleTabs
    .map(
      (t) => `
      <section x-show="tab === '${t.id}'" x-transition:enter="transition ease-out duration-200"
        x-transition:enter-start="opacity-0 translate-y-1" x-transition:enter-end="opacity-100 translate-y-0">
        ${t.content}
      </section>`
    )
    .join("\n");

  return `<!DOCTYPE html>
<html lang="en" class="dark">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>forge-visual — Report</title>
  <script src="https://cdn.tailwindcss.com"></script>
  <script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"></script>
  <script>
    tailwind.config = {
      darkMode: 'class',
      theme: {
        extend: {
          colors: {
            dark: { 900: '#0a0a0f', 800: '#12121a', 700: '#1a1a2e', 600: '#252540' },
            accent: { DEFAULT: '#6366f1', light: '#818cf8' },
            gas: { low: '#10b981', mid: '#f59e0b', high: '#ef4444' },
          }
        }
      }
    }
  </script>
  <style>
    body { background: #0a0a0f; }
    .glass { background: rgba(18, 18, 26, 0.8); backdrop-filter: blur(12px); border: 1px solid rgba(99, 102, 241, 0.15); }
    .gas-heatmap-low { background: rgba(16, 185, 129, 0.15); color: #6ee7b7; }
    .gas-heatmap-mid { background: rgba(245, 158, 11, 0.15); color: #fcd34d; }
    .gas-heatmap-high { background: rgba(239, 68, 68, 0.15); color: #fca5a5; }
    .status-pass { color: #10b981; }
    .status-fail { color: #ef4444; }
    .status-skip { color: #6b7280; }
    .coverage-high { color: #10b981; }
    .coverage-medium { color: #f59e0b; }
    .coverage-low { color: #ef4444; }
    .bar-high { background: #10b981; }
    .bar-medium { background: #f59e0b; }
    .bar-low { background: #ef4444; }
    [x-cloak] { display: none !important; }
    ::-webkit-scrollbar { width: 6px; height: 6px; }
    ::-webkit-scrollbar-track { background: transparent; }
    ::-webkit-scrollbar-thumb { background: #252540; border-radius: 3px; }
    ::-webkit-scrollbar-thumb:hover { background: #6366f1; }
  </style>
</head>
<body class="text-gray-200 min-h-screen" x-data="{
  tab: window.location.hash.slice(1) || '${visibleTabs[0]?.id ?? "dashboard"}',
  collapsed: false,
  projectType: '${opts.projectType}',
  init() {
    window.addEventListener('hashchange', () => {
      this.tab = window.location.hash.slice(1) || '${visibleTabs[0]?.id ?? "dashboard"}';
    });
  },
  navigate(t) {
    this.tab = t;
    window.location.hash = t;
  }
}" x-cloak>

  <div class="flex min-h-screen">
    <!-- Sidebar -->
    <aside class="glass sticky top-0 h-screen flex flex-col border-r border-dark-600 transition-all duration-300 z-50"
      :class="collapsed ? 'w-14' : 'w-56'">

      <!-- Logo -->
      <div class="px-4 py-4 border-b border-dark-600 flex items-center gap-2">
        <span class="text-accent font-bold text-lg">fv</span>
        <span x-show="!collapsed" class="text-sm text-gray-400 font-medium">forge-visual</span>
        <button @click="collapsed=!collapsed" class="ml-auto text-gray-500 hover:text-gray-300 text-xs">
          <span x-text="collapsed ? '>' : '<'"></span>
        </button>
      </div>

      <!-- Project badge -->
      <div class="px-4 py-3 border-b border-dark-600" x-show="!collapsed">
        <span class="inline-block text-xs px-2 py-0.5 rounded border ${badge.color}">${badge.label}</span>
        ${opts.facets.length > 0 ? `<span class="text-xs text-gray-500 ml-1">${opts.facets.length} facets</span>` : ""}
      </div>

      <!-- Nav links -->
      <nav class="flex-1 py-2 space-y-0.5 overflow-y-auto">
        ${sidebarLinks}
      </nav>

      <!-- Footer -->
      <div class="px-4 py-3 border-t border-dark-600 text-xs text-gray-600" x-show="!collapsed">
        ${opts.gitBranch ? `<div>${opts.gitBranch}</div>` : ""}
        ${opts.gitRef ? `<div class="font-mono">${opts.gitRef}</div>` : ""}
        <div class="mt-1">${opts.generatedAt}</div>
      </div>
    </aside>

    <!-- Main content -->
    <main class="flex-1 min-w-0 px-6 py-6 max-w-7xl">
      ${sections}
    </main>
  </div>

  ${opts.scripts}
</body>
</html>`;
}
