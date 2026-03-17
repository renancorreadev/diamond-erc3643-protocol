import { useQuery } from "@tanstack/react-query";

const INDEXER_URL = import.meta.env.VITE_INDEXER_URL ?? "";
const GRAPHQL_ENDPOINT = INDEXER_URL ? `${INDEXER_URL}/graphql` : "";

async function gqlFetch<T>(query: string, variables?: Record<string, unknown>): Promise<T> {
  const response = await fetch(GRAPHQL_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query, variables }),
  });

  if (!response.ok) {
    throw new Error(`Indexer request failed: ${response.status}`);
  }

  const json = await response.json();

  if (json.errors) {
    throw new Error(json.errors[0]?.message ?? "GraphQL error");
  }

  return json.data as T;
}

// ---------- Tokens ----------

export interface IndexerToken {
  id: string;
  totalSupply: string;
  holderCount: number;
}

interface TokensResponse {
  tokens: IndexerToken[];
}

const TOKENS_QUERY = `
  query Tokens {
    tokens {
      id
      totalSupply
      holderCount
    }
  }
`;

export function useIndexerTokens() {
  return useQuery<IndexerToken[]>({
    queryKey: ["indexer", "tokens"],
    queryFn: async () => {
      const data = await gqlFetch<TokensResponse>(TOKENS_QUERY);
      return data.tokens;
    },
    enabled: !!GRAPHQL_ENDPOINT,
  });
}

// ---------- Token Detail ----------

export interface IndexerTokenDetail {
  id: string;
  totalSupply: string;
  holderCount: number;
  holders: { address: string; balance: string }[];
  events: IndexerEvent[];
}

interface TokenDetailResponse {
  token: IndexerTokenDetail | null;
}

const TOKEN_DETAIL_QUERY = `
  query Token($id: String!) {
    token(id: $id) {
      id
      totalSupply
      holderCount
      holders(first: 50) {
        address
        balance
      }
      events(first: 20) {
        txHash
        block
        logIndex
        from
        to
        tokenId
        amount
        eventType
      }
    }
  }
`;

export function useIndexerToken(tokenId: string | undefined) {
  return useQuery<IndexerTokenDetail | null>({
    queryKey: ["indexer", "token", tokenId],
    queryFn: async () => {
      const data = await gqlFetch<TokenDetailResponse>(TOKEN_DETAIL_QUERY, { id: tokenId });
      return data.token;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!tokenId,
  });
}

// ---------- Events ----------

export interface IndexerEvent {
  txHash: string;
  block: number;
  logIndex: number;
  from: string;
  to: string;
  tokenId: string;
  amount: string;
  eventType: string;
}

interface EventsResponse {
  events: IndexerEvent[];
}

const EVENTS_QUERY = `
  query Events($first: Int!) {
    events(first: $first) {
      txHash
      block
      logIndex
      from
      to
      tokenId
      amount
      eventType
    }
  }
`;

export function useIndexerEvents(first: number = 20) {
  return useQuery<IndexerEvent[]>({
    queryKey: ["indexer", "events", first],
    queryFn: async () => {
      const data = await gqlFetch<EventsResponse>(EVENTS_QUERY, { first });
      return data.events;
    },
    enabled: !!GRAPHQL_ENDPOINT,
  });
}

// ---------- Holder ----------

export interface IndexerHolder {
  address: string;
  balance: string;
}

interface HolderResponse {
  holder: IndexerHolder | null;
}

const HOLDER_QUERY = `
  query Holder($tokenId: String!, $address: String!) {
    holder(tokenId: $tokenId, address: $address) {
      address
      balance
    }
  }
`;

export function useIndexerHolder(tokenId: string | undefined, address: string | undefined) {
  return useQuery<IndexerHolder | null>({
    queryKey: ["indexer", "holder", tokenId, address],
    queryFn: async () => {
      const data = await gqlFetch<HolderResponse>(HOLDER_QUERY, { tokenId, address });
      return data.holder;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!tokenId && !!address,
  });
}

// ---------- Status ----------

export interface IndexerStatus {
  lastBlock: number;
  tokenCount: number;
}

interface StatusResponse {
  status: IndexerStatus;
}

const STATUS_QUERY = `
  query Status {
    status {
      lastBlock
      tokenCount
    }
  }
`;

export function useIndexerStatus() {
  return useQuery<IndexerStatus>({
    queryKey: ["indexer", "status"],
    queryFn: async () => {
      const data = await gqlFetch<StatusResponse>(STATUS_QUERY);
      return data.status;
    },
    enabled: !!GRAPHQL_ENDPOINT,
    refetchInterval: 10_000,
  });
}

// ---------- Assets ----------

export interface IndexerAsset {
  id: string;
  name: string;
  symbol: string;
  issuer: string;
  profileId: number;
  uri: string | null;
  paused: boolean;
  totalSupply: string;
  holderCount: number;
  registeredAt: number;
}

interface AssetsResponse {
  assets: IndexerAsset[];
}

const ASSETS_QUERY = `
  query Assets {
    assets {
      id
      name
      symbol
      issuer
      profileId
      uri
      paused
      totalSupply
      holderCount
      registeredAt
    }
  }
`;

export function useIndexerAssets() {
  return useQuery<IndexerAsset[]>({
    queryKey: ["indexer", "assets"],
    queryFn: async () => {
      const data = await gqlFetch<AssetsResponse>(ASSETS_QUERY);
      return data.assets;
    },
    enabled: !!GRAPHQL_ENDPOINT,
  });
}

// ---------- Asset Detail ----------

export interface IndexerAssetDetail extends IndexerAsset {
  holders: { address: string; balance: string }[];
  events: IndexerProtocolEvent[];
}

interface AssetDetailResponse {
  asset: IndexerAssetDetail | null;
}

const ASSET_DETAIL_QUERY = `
  query Asset($id: String!) {
    asset(id: $id) {
      id
      name
      symbol
      issuer
      profileId
      uri
      paused
      totalSupply
      holderCount
      registeredAt
      holders(first: 50) {
        address
        balance
      }
      events(first: 20) {
        txHash
        block
        logIndex
        eventType
        tokenId
        address
        data
      }
    }
  }
`;

export function useIndexerAsset(id: string | undefined) {
  return useQuery<IndexerAssetDetail | null>({
    queryKey: ["indexer", "asset", id],
    queryFn: async () => {
      const data = await gqlFetch<AssetDetailResponse>(ASSET_DETAIL_QUERY, { id });
      return data.asset;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!id,
  });
}

// ---------- Identity ----------

export interface IndexerIdentity {
  wallet: string;
  identity: string;
  country: number;
  boundAt: number;
}

interface IdentityResponse {
  identity: IndexerIdentity | null;
}

const IDENTITY_QUERY = `
  query Identity($wallet: String!) {
    identity(wallet: $wallet) {
      wallet
      identity
      country
      boundAt
    }
  }
`;

export function useIndexerIdentity(wallet: string | undefined) {
  return useQuery<IndexerIdentity | null>({
    queryKey: ["indexer", "identity", wallet],
    queryFn: async () => {
      const data = await gqlFetch<IdentityResponse>(IDENTITY_QUERY, { wallet });
      return data.identity;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!wallet,
  });
}

// ---------- Identities ----------

interface IdentitiesResponse {
  identities: IndexerIdentity[];
}

const IDENTITIES_QUERY = `
  query Identities($country: Int, $first: Int!) {
    identities(country: $country, first: $first) {
      wallet
      identity
      country
      boundAt
    }
  }
`;

export function useIndexerIdentities(country?: number, first: number = 50) {
  return useQuery<IndexerIdentity[]>({
    queryKey: ["indexer", "identities", country, first],
    queryFn: async () => {
      const data = await gqlFetch<IdentitiesResponse>(IDENTITIES_QUERY, { country, first });
      return data.identities;
    },
    enabled: !!GRAPHQL_ENDPOINT,
  });
}

// ---------- Freezes ----------

export interface IndexerFreezeRecord {
  wallet: string;
  tokenId: string | null;
  frozen: boolean;
  frozenAmount: string | null;
  lockupExpiry: number | null;
}

interface FreezesResponse {
  freezes: IndexerFreezeRecord[];
}

const FREEZES_QUERY = `
  query Freezes($wallet: String!) {
    freezes(wallet: $wallet) {
      wallet
      tokenId
      frozen
      frozenAmount
      lockupExpiry
    }
  }
`;

export function useIndexerFreezes(wallet: string | undefined) {
  return useQuery<IndexerFreezeRecord[]>({
    queryKey: ["indexer", "freezes", wallet],
    queryFn: async () => {
      const data = await gqlFetch<FreezesResponse>(FREEZES_QUERY, { wallet });
      return data.freezes;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!wallet,
  });
}

// ---------- Protocol Events ----------

export interface IndexerProtocolEvent {
  txHash: string;
  block: number;
  logIndex: number;
  eventType: string;
  tokenId: string | null;
  address: string | null;
  data: string;
}

interface ProtocolEventsResponse {
  protocolEvents: IndexerProtocolEvent[];
}

const PROTOCOL_EVENTS_QUERY = `
  query ProtocolEvents($first: Int!, $eventType: String, $tokenId: String, $address: String) {
    protocolEvents(first: $first, eventType: $eventType, tokenId: $tokenId, address: $address) {
      txHash
      block
      logIndex
      eventType
      tokenId
      address
      data
    }
  }
`;

export interface ProtocolEventFilters {
  first?: number;
  eventType?: string;
  tokenId?: string;
  address?: string;
}

export function useIndexerProtocolEvents(filters?: ProtocolEventFilters) {
  const first = filters?.first ?? 50;
  return useQuery<IndexerProtocolEvent[]>({
    queryKey: ["indexer", "protocolEvents", filters],
    queryFn: async () => {
      const data = await gqlFetch<ProtocolEventsResponse>(PROTOCOL_EVENTS_QUERY, {
        first,
        eventType: filters?.eventType,
        tokenId: filters?.tokenId,
        address: filters?.address,
      });
      return data.protocolEvents;
    },
    enabled: !!GRAPHQL_ENDPOINT,
  });
}

// ---------- Portfolio ----------

export interface IndexerPortfolioHolding {
  tokenId: string;
  balance: string;
}

interface PortfolioResponse {
  portfolio: IndexerPortfolioHolding[];
}

const PORTFOLIO_QUERY = `
  query Portfolio($address: String!) {
    portfolio(address: $address) {
      tokenId
      balance
    }
  }
`;

export function useIndexerPortfolio(address: string | undefined) {
  return useQuery<IndexerPortfolioHolding[]>({
    queryKey: ["indexer", "portfolio", address],
    queryFn: async () => {
      const data = await gqlFetch<PortfolioResponse>(PORTFOLIO_QUERY, { address });
      return data.portfolio;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!address,
  });
}
