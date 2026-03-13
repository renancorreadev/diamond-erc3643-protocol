# ── Stage 1: Build RocksDB + Go binary ──
FROM golang:1.24-bookworm AS builder

# Install build tools and RocksDB compile dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    cmake \
    libsnappy-dev \
    liblz4-dev \
    libzstd-dev \
    libbz2-dev \
    zlib1g-dev \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Build RocksDB 9.7.4 from source (compatible with grocksdb v1.10.7)
ENV ROCKSDB_VERSION=9.7.4
RUN wget -qO- https://github.com/facebook/rocksdb/archive/refs/tags/v${ROCKSDB_VERSION}.tar.gz | tar xz \
    && cd rocksdb-${ROCKSDB_VERSION} \
    && make shared_lib -j$(nproc) \
    && make install-shared INSTALL_PATH=/usr/local \
    && cd / && rm -rf rocksdb-${ROCKSDB_VERSION} \
    && ldconfig

WORKDIR /src

# Cache Go modules
COPY go.mod go.sum ./
ENV CGO_ENABLED=1
ENV CGO_CFLAGS="-I/usr/local/include"
ENV CGO_LDFLAGS="-L/usr/local/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd"
RUN go mod download

# Copy source and build
COPY . .
RUN go build -ldflags="-s -w" -o /indexer ./cmd/indexer/

# ── Stage 2: Runtime ──
FROM debian:bookworm-slim

# Install runtime libs
RUN apt-get update && apt-get install -y --no-install-recommends \
    libsnappy1v5 \
    liblz4-1 \
    libzstd1 \
    libbz2-1.0 \
    zlib1g \
    libstdc++6 \
    ca-certificates \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Copy RocksDB shared lib from builder
COPY --from=builder /usr/local/lib/librocksdb.so* /usr/local/lib/
RUN ldconfig

# Create non-root user
RUN useradd -r -s /bin/false indexer

WORKDIR /app

# Copy binary
COPY --from=builder /indexer /app/indexer

# Data volume
RUN mkdir -p /app/data && chown indexer:indexer /app/data
VOLUME ["/app/data"]

USER indexer

ENV HTTP_LISTEN=:8080
ENV DATA_DIR=/app/data

EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=3s --start-period=10s \
  CMD wget -qO- http://localhost:8080/health || exit 1

ENTRYPOINT ["/app/indexer"]
