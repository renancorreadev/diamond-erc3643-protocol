import { useMemo } from 'react';
import { Link } from 'react-router-dom';
import { useAccount, useReadContract, useReadContracts } from 'wagmi';

import { DIAMOND_ADDRESS, diamondAbi } from '@/config/contracts';
import { formatTokenAmount } from '@/lib/format';
import { useIndexerPortfolio, useIndexerAssets } from '@/hooks/use-indexer';

const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;

type AssetConfig = {
  name: string;
  symbol: string;
  uri: string;
  supplyCap: bigint;
  totalSupply: bigint;
  identityProfileId: number;
  complianceModule: string;
  issuer: string;
  paused: boolean;
  exists: boolean;
  allowedCountries: number[];
};

export default function PortfolioPage() {
  const { address, isConnected } = useAccount();

  // Indexer: fast portfolio + asset metadata
  const { data: indexerPortfolio, isLoading: indexerLoading } = useIndexerPortfolio(address);
  const { data: indexerAssets } = useIndexerAssets();

  // Build asset lookup from indexer
  const assetMap = useMemo(() => {
    const map = new Map<string, { name: string; symbol: string }>();
    if (indexerAssets) {
      for (const a of indexerAssets) {
        map.set(a.id, { name: a.name, symbol: a.symbol });
      }
    }
    return map;
  }, [indexerAssets]);

  // On-chain fallback: read all token IDs + balances + configs
  const { data: tokenIds, isLoading: idsLoading } = useReadContract({
    ...diamond,
    functionName: 'getRegisteredTokenIds',
  });

  const registeredIds = (tokenIds as bigint[]) ?? [];

  const balanceContracts = useMemo(() => {
    if (!isConnected || !address || registeredIds.length === 0) return [];
    return registeredIds.map((id) => ({
      ...diamond,
      functionName: 'balanceOf' as const,
      args: [address, id] as const,
    }));
  }, [tokenIds, isConnected, address]);

  const configContracts = useMemo(() => {
    if (registeredIds.length === 0) return [];
    return registeredIds.map((id) => ({
      ...diamond,
      functionName: 'getAssetConfig' as const,
      args: [id] as const,
    }));
  }, [tokenIds]);

  const { data: balancesData, isLoading: balancesLoading } = useReadContracts({
    contracts: balanceContracts,
    query: { enabled: balanceContracts.length > 0 },
  });

  const { data: configsData, isLoading: configsLoading } = useReadContracts({
    contracts: configContracts,
    query: { enabled: configContracts.length > 0 },
  });

  // Use indexer data if available, otherwise wait for on-chain
  const hasIndexerData = !!indexerPortfolio && indexerPortfolio.length > 0;
  const onChainLoading = idsLoading || balancesLoading || configsLoading;
  const isLoading = hasIndexerData ? false : (indexerLoading || onChainLoading);

  if (!isConnected) {
    return (
      <div className="min-h-screen bg-[#0a0a0f] p-8">
        <h1 className="mb-8 text-3xl font-bold text-white">Portfolio</h1>
        <div className="rounded-xl bg-white/5 border border-white/10 p-6 text-center">
          <p className="text-gray-400 text-lg">Connect your wallet to view your portfolio.</p>
        </div>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="min-h-screen bg-[#0a0a0f] p-8">
        <h1 className="mb-8 text-3xl font-bold text-white">Portfolio</h1>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {[1, 2, 3].map((i) => (
            <div key={i} className="rounded-xl bg-white/5 border border-white/10 p-6 animate-pulse">
              <div className="h-4 w-24 rounded bg-white/10 mb-3" />
              <div className="h-6 w-16 rounded bg-white/10 mb-2" />
              <div className="h-4 w-32 rounded bg-white/10" />
            </div>
          ))}
        </div>
      </div>
    );
  }

  // Merge: prefer on-chain data when available, fall back to indexer
  const onChainTokens = registeredIds.map((id, i) => {
    const balance = balancesData?.[i]?.result as bigint | undefined;
    const config = configsData?.[i]?.result as AssetConfig | undefined;
    return { id: id.toString(), balance: balance ?? 0n, name: config?.name, symbol: config?.symbol };
  });

  const onChainWithBalance = onChainTokens.filter((t) => t.balance > 0n);

  // If on-chain data is ready, use it; otherwise use indexer
  const tokensWithBalance = onChainWithBalance.length > 0
    ? onChainWithBalance.map((t) => ({
        id: t.id,
        balance: t.balance,
        name: t.name ?? assetMap.get(t.id)?.name,
        symbol: t.symbol ?? assetMap.get(t.id)?.symbol,
      }))
    : (indexerPortfolio ?? []).map((h) => ({
        id: h.tokenId,
        balance: BigInt(h.balance),
        name: assetMap.get(h.tokenId)?.name,
        symbol: assetMap.get(h.tokenId)?.symbol,
      }));

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Portfolio</h1>

      {tokensWithBalance.length === 0 ? (
        <div className="rounded-xl bg-white/5 border border-white/10 p-6 text-center">
          <p className="text-gray-400 text-lg">No tokens found in your wallet.</p>
          <p className="text-gray-500 text-sm mt-2">
            You don&apos;t hold any registered security tokens yet.
          </p>
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {tokensWithBalance.map((token) => (
            <Link key={token.id} to={`/portfolio/${token.id}`}>
              <div className="rounded-xl bg-white/5 border border-white/10 p-6 transition-colors hover:border-indigo-400/50 hover:bg-white/[0.08]">
                <div className="flex items-center justify-between mb-3">
                  <h3 className="text-lg font-semibold text-indigo-400">
                    {token.name ?? `Token #${token.id}`}
                  </h3>
                  <span className="text-xs font-mono text-gray-500 bg-white/5 px-2 py-1 rounded">
                    {token.symbol ?? '???'}
                  </span>
                </div>
                <p className="text-2xl font-bold text-white">
                  {formatTokenAmount(token.balance)}
                </p>
                <p className="text-sm text-gray-500 mt-1">
                  Token ID: {token.id}
                </p>
              </div>
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}
