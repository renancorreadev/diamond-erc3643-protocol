package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/store"
)

// NewSchema builds the GraphQL schema backed by a RocksDB store.
func NewSchema(db *store.Store) (graphql.Schema, error) {
	// ── Types ──────────────────────────────────────────────────

	holderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Holder",
		Fields: graphql.Fields{
			"address": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"balance": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	transferEventType := graphql.NewObject(graphql.ObjectConfig{
		Name: "TransferEvent",
		Fields: graphql.Fields{
			"txHash":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"block":     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"logIndex":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"from":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"to":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"tokenId":   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"amount":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"eventType": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	protocolEventType := graphql.NewObject(graphql.ObjectConfig{
		Name: "ProtocolEvent",
		Fields: graphql.Fields{
			"txHash":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"block":     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"logIndex":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"eventType": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"tokenId":   &graphql.Field{Type: graphql.String},
			"address":   &graphql.Field{Type: graphql.String},
			"data":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	tokenType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Token",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(*store.TokenMeta).TokenID, nil
				},
			},
			"totalSupply": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"holderCount": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"holders": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(holderType))),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 50,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					meta := p.Source.(*store.TokenMeta)
					holders, err := db.GetHolders(meta.TokenID)
					if err != nil {
						return nil, err
					}
					limit := p.Args["first"].(int)
					if limit > 0 && limit < len(holders) {
						holders = holders[:limit]
					}
					return holders, nil
				},
			},
			"events": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(transferEventType))),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 20,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					meta := p.Source.(*store.TokenMeta)
					limit := p.Args["first"].(int)
					events, err := db.GetTokenEvents(meta.TokenID, limit)
					if err != nil {
						return nil, err
					}
					return events, nil
				},
			},
		},
	})

	assetType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Asset",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(*assetView).TokenID, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(*assetView).Name, nil
				},
			},
			"symbol": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(*assetView).Symbol, nil
				},
			},
			"issuer": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(*assetView).Issuer, nil
				},
			},
			"profileId": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int(p.Source.(*assetView).ProfileID), nil
				},
			},
			"uri":    &graphql.Field{Type: graphql.String},
			"paused": &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean)},
			"totalSupply": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(*assetView).TotalSupply, nil
				},
			},
			"holderCount": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int(p.Source.(*assetView).HolderCount), nil
				},
			},
			"registeredAt": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int(p.Source.(*assetView).RegisteredAt), nil
				},
			},
			"holders": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(holderType))),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 50},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					av := p.Source.(*assetView)
					holders, err := db.GetHolders(av.TokenID)
					if err != nil {
						return nil, err
					}
					limit := p.Args["first"].(int)
					if limit > 0 && limit < len(holders) {
						holders = holders[:limit]
					}
					return holders, nil
				},
			},
			"events": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(protocolEventType))),
				Args: graphql.FieldConfigArgument{
					"first":     &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 20},
					"eventType": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					av := p.Source.(*assetView)
					limit := p.Args["first"].(int)
					var et string
					if v, ok := p.Args["eventType"].(string); ok {
						et = v
					}
					events, err := db.GetProtocolEvents(limit, et, av.TokenID, "")
					if err != nil {
						return nil, err
					}
					return events, nil
				},
			},
		},
	})

	identityRecordType := graphql.NewObject(graphql.ObjectConfig{
		Name: "IdentityRecord",
		Fields: graphql.Fields{
			"wallet":   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"identity": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"country":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"boundAt":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		},
	})

	freezeRecordType := graphql.NewObject(graphql.ObjectConfig{
		Name: "FreezeRecord",
		Fields: graphql.Fields{
			"wallet":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"tokenId":      &graphql.Field{Type: graphql.String},
			"frozen":       &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean)},
			"frozenAmount": &graphql.Field{Type: graphql.String},
			"lockupExpiry": &graphql.Field{Type: graphql.Int},
		},
	})

	portfolioHoldingType := graphql.NewObject(graphql.ObjectConfig{
		Name: "PortfolioHolding",
		Fields: graphql.Fields{
			"tokenId": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"balance": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	indexerStatusType := graphql.NewObject(graphql.ObjectConfig{
		Name: "IndexerStatus",
		Fields: graphql.Fields{
			"lastBlock":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"tokenCount": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		},
	})

	// ── Queries ────────────────────────────────────────────────

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// ── Existing queries ──
			"token": &graphql.Field{
				Type: tokenType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					meta, err := db.GetTokenMeta(id)
					if err != nil {
						return nil, err
					}
					if meta.TotalSupply == "0" && meta.HolderCount == 0 {
						return nil, nil
					}
					return meta, nil
				},
			},
			"tokens": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(tokenType))),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return db.GetAllTokens()
				},
			},
			"holder": &graphql.Field{
				Type: holderType,
				Args: graphql.FieldConfigArgument{
					"tokenId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"address": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					tokenID := p.Args["tokenId"].(string)
					addr := p.Args["address"].(string)
					bal, err := db.GetHolderBalance(tokenID, addr)
					if err != nil {
						return nil, err
					}
					return store.Holder{Address: addr, Balance: bal}, nil
				},
			},
			"events": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(transferEventType))),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 50,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					limit := p.Args["first"].(int)
					return db.GetRecentEvents(limit)
				},
			},
			"status": &graphql.Field{
				Type: graphql.NewNonNull(indexerStatusType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					cursor, _ := db.GetCursor()
					tokens, _ := db.GetAllTokens()
					return map[string]interface{}{
						"lastBlock":  cursor,
						"tokenCount": len(tokens),
					}, nil
				},
			},

			// ── Assets ──
			"asset": &graphql.Field{
				Type: assetType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					return buildAssetView(db, id)
				},
			},
			"assets": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(assetType))),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					configs, err := db.GetAllAssets()
					if err != nil {
						return nil, err
					}
					var views []*assetView
					for _, cfg := range configs {
						av, err := buildAssetView(db, cfg.TokenID)
						if err != nil || av == nil {
							continue
						}
						views = append(views, av)
					}
					return views, nil
				},
			},

			// ── Identity ──
			"identity": &graphql.Field{
				Type: identityRecordType,
				Args: graphql.FieldConfigArgument{
					"wallet": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					wallet := p.Args["wallet"].(string)
					return db.GetIdentity(wallet)
				},
			},
			"identities": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(identityRecordType))),
				Args: graphql.FieldConfigArgument{
					"country": &graphql.ArgumentConfig{Type: graphql.Int},
					"first":   &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 50},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					limit := p.Args["first"].(int)
					if country, ok := p.Args["country"].(int); ok {
						return db.GetIdentitiesByCountry(uint16(country), limit)
					}
					return db.GetAllIdentities(limit)
				},
			},

			// ── Freeze ──
			"freezes": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(freezeRecordType))),
				Args: graphql.FieldConfigArgument{
					"wallet": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					wallet := p.Args["wallet"].(string)
					return db.GetFreezes(wallet)
				},
			},
			"frozenWallets": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(freezeRecordType))),
				Args: graphql.FieldConfigArgument{
					"tokenId": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var tokenID string
					if v, ok := p.Args["tokenId"].(string); ok {
						tokenID = v
					}
					return db.GetFrozenWallets(tokenID)
				},
			},

			// ── Protocol Events ──
			"protocolEvents": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(protocolEventType))),
				Args: graphql.FieldConfigArgument{
					"first":     &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 50},
					"eventType": &graphql.ArgumentConfig{Type: graphql.String},
					"tokenId":   &graphql.ArgumentConfig{Type: graphql.String},
					"address":   &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					limit := p.Args["first"].(int)
					var eventType, tokenID, address string
					if v, ok := p.Args["eventType"].(string); ok {
						eventType = v
					}
					if v, ok := p.Args["tokenId"].(string); ok {
						tokenID = v
					}
					if v, ok := p.Args["address"].(string); ok {
						address = v
					}
					return db.GetProtocolEvents(limit, eventType, tokenID, address)
				},
			},

			// ── Portfolio ──
			"portfolio": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(portfolioHoldingType))),
				Args: graphql.FieldConfigArgument{
					"address": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					address := p.Args["address"].(string)
					holders, err := db.GetPortfolio(address)
					if err != nil {
						return nil, err
					}
					// Map from Holder (address=tokenId, balance) to PortfolioHolding shape
					var holdings []map[string]interface{}
					for _, h := range holders {
						holdings = append(holdings, map[string]interface{}{
							"tokenId": h.Address,
							"balance": h.Balance,
						})
					}
					return holdings, nil
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

// assetView merges AssetConfig with TokenMeta for the Asset GraphQL type.
type assetView struct {
	TokenID      string
	Name         string
	Symbol       string
	Issuer       string
	ProfileID    uint32
	URI          string
	Paused       bool
	TotalSupply  string
	HolderCount  uint64
	RegisteredAt uint64
}

func buildAssetView(db *store.Store, tokenID string) (*assetView, error) {
	cfg, err := db.GetAsset(tokenID)
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, nil
	}

	meta, err := db.GetTokenMeta(tokenID)
	if err != nil {
		return nil, err
	}

	return &assetView{
		TokenID:      cfg.TokenID,
		Name:         cfg.Name,
		Symbol:       cfg.Symbol,
		Issuer:       cfg.Issuer,
		ProfileID:    cfg.ProfileID,
		URI:          cfg.URI,
		Paused:       cfg.Paused,
		TotalSupply:  meta.TotalSupply,
		HolderCount:  meta.HolderCount,
		RegisteredAt: cfg.RegisteredAt,
	}, nil
}
