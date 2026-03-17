import { truncateAddress } from "@/lib/format";

interface ProtocolEvent {
  txHash: string;
  block: number;
  logIndex: number;
  eventType: string;
  tokenId: string | null;
  address: string | null;
  data: string;
}

const EVENT_CONFIG: Record<string, { label: string; color: string; icon: string }> = {
  asset_registered:          { label: "Asset Registered",       color: "text-indigo-400",  icon: "+" },
  asset_config_updated:      { label: "Config Updated",         color: "text-indigo-400",  icon: "~" },
  uri_updated:               { label: "URI Updated",            color: "text-indigo-400",  icon: "~" },
  identity_bound:            { label: "Identity Bound",         color: "text-cyan-400",    icon: "ID" },
  identity_unbound:          { label: "Identity Unbound",       color: "text-gray-400",    icon: "ID" },
  wallet_frozen:             { label: "Wallet Frozen",          color: "text-blue-400",    icon: "*" },
  asset_frozen:              { label: "Asset Frozen",           color: "text-blue-400",    icon: "*" },
  partial_freeze:            { label: "Partial Freeze",         color: "text-blue-400",    icon: "#" },
  lockup_set:                { label: "Lockup Set",             color: "text-blue-400",    icon: "T" },
  role_granted:              { label: "Role Granted",           color: "text-emerald-400", icon: "R" },
  role_revoked:              { label: "Role Revoked",           color: "text-orange-400",  icon: "R" },
  emergency_pause:           { label: "Emergency Pause",        color: "text-red-400",     icon: "!" },
  protocol_unpaused:         { label: "Protocol Unpaused",      color: "text-green-400",   icon: ">" },
  asset_paused:              { label: "Asset Paused",           color: "text-red-400",     icon: "||" },
  asset_unpaused:            { label: "Asset Unpaused",         color: "text-green-400",   icon: ">" },
  wallet_recovered:          { label: "Wallet Recovered",       color: "text-amber-400",   icon: "W" },
  snapshot_created:          { label: "Snapshot Created",       color: "text-purple-400",  icon: "S" },
  dividend_created:          { label: "Dividend Created",       color: "text-emerald-400", icon: "$" },
  dividend_claimed:          { label: "Dividend Claimed",       color: "text-emerald-400", icon: "$" },
  compliance_module_added:   { label: "Module Added",           color: "text-indigo-400",  icon: "M" },
  compliance_module_removed: { label: "Module Removed",         color: "text-orange-400",  icon: "M" },
  group_created:             { label: "Group Created",          color: "text-purple-400",  icon: "G" },
  unit_minted:               { label: "Unit Minted",            color: "text-emerald-400", icon: "U" },
};

function getConfig(eventType: string) {
  return EVENT_CONFIG[eventType] ?? { label: eventType, color: "text-gray-400", icon: "?" };
}

interface ActivityFeedProps {
  events: ProtocolEvent[];
  isLoading?: boolean;
  title?: string;
  compact?: boolean;
}

export function ActivityFeed({ events, isLoading, title = "Recent Activity", compact }: ActivityFeedProps) {
  if (isLoading) {
    return (
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h3 className="text-lg font-semibold text-white mb-4">{title}</h3>
        <div className="space-y-3">
          {[1, 2, 3].map((i) => (
            <div key={i} className="flex gap-3 animate-pulse">
              <div className="h-8 w-8 rounded-lg bg-white/10 shrink-0" />
              <div className="flex-1">
                <div className="h-4 w-32 rounded bg-white/10 mb-1" />
                <div className="h-3 w-48 rounded bg-white/10" />
              </div>
            </div>
          ))}
        </div>
      </div>
    );
  }

  if (!events || events.length === 0) {
    return (
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h3 className="text-lg font-semibold text-white mb-4">{title}</h3>
        <p className="text-sm text-gray-500">No events recorded yet.</p>
      </div>
    );
  }

  return (
    <div className="rounded-xl bg-white/5 border border-white/10 p-6">
      <h3 className="text-lg font-semibold text-white mb-4">{title}</h3>
      <div className="space-y-1">
        {events.map((evt) => {
          const cfg = getConfig(evt.eventType);
          return (
            <div
              key={`${evt.txHash}-${evt.logIndex}`}
              className="flex items-center gap-3 rounded-lg px-3 py-2 transition-colors hover:bg-white/[0.03]"
            >
              <div className={`flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-white/5 border border-white/10 text-xs font-bold ${cfg.color}`}>
                {cfg.icon}
              </div>
              <div className="flex-1 min-w-0">
                <div className="flex items-center gap-2">
                  <span className={`text-sm font-medium ${cfg.color}`}>{cfg.label}</span>
                  {evt.tokenId && (
                    <span className="text-xs text-gray-600 font-mono">#{evt.tokenId}</span>
                  )}
                </div>
                {!compact && evt.address && (
                  <p className="text-xs text-gray-500 font-mono truncate">
                    {truncateAddress(evt.address)}
                  </p>
                )}
              </div>
              <div className="text-right shrink-0">
                <p className="text-xs text-gray-600 font-mono">
                  {evt.block.toLocaleString()}
                </p>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}
