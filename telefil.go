package telefil

import (
	"context"
	"encoding/base64"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/ybbus/jsonrpc/v3"
)

const (
	methodFilStateMinerInfo                    = `Filecoin.StateMinerInfo`
	methodFilChainHead                         = `Filecoin.ChainHead`
	methodFilChainGetGenesis                   = `Filecoin.ChainGetGenesis`
	methodFilStateDealProviderCollateralBounds = `Filecoin.StateDealProviderCollateralBounds`
	methodFilStateListMiners                   = `Filecoin.StateListMiners`
	methodFilStateMinerPower                   = `Filecoin.StateMinerPower`
)

type (
	Telefil struct {
		client jsonrpc.RPCClient
	}
	ChainHead struct {
		Height int64 `json:"Height"`
	}
	ChainGenesis struct {
		Blocks []struct {
			Timestamp int64 `json:"Timestamp"`
		} `json:"Blocks"`
	}
	StateDealProviderCollateralBounds struct {
		Min big.Int `json:"Min"`
		Max big.Int `json:"Max"`
	}
	StateMinerPower struct {
		HasMinPower bool  `json:"HasMinPower"`
		MinerPower  Claim `json:"MinerPower"`
		TotalPower  Claim `json:"TotalPower"`
	}
	Claim struct {
		QualityAdjPower big.Int `json:"QualityAdjPower"`
		RawBytePower    big.Int `json:"RawBytePower"`
	}
)

func New(o ...Option) (*Telefil, error) {
	opts, err := newOptions(o...)
	if err != nil {
		return nil, err
	}
	return &Telefil{
		client: jsonrpc.NewClient(opts.api),
	}, nil
}

func (f *Telefil) ChainHead(ctx context.Context) (*ChainHead, error) {
	switch resp, err := f.client.Call(ctx, methodFilChainHead); {
	case err != nil:
		return nil, err
	case resp.Error != nil:
		return nil, resp.Error
	default:
		var c ChainHead
		if err := resp.GetObject(&c); err != nil {
			return nil, err
		}
		return &c, nil
	}
}

func (f *Telefil) ChainGetGenesis(ctx context.Context) (*ChainGenesis, error) {
	switch resp, err := f.client.Call(ctx, methodFilChainGetGenesis); {
	case err != nil:
		return nil, err
	case resp.Error != nil:
		return nil, resp.Error
	default:
		var c ChainGenesis
		if err := resp.GetObject(&c); err != nil {
			return nil, err
		}
		return &c, nil
	}
}

func (f *Telefil) StateDealProviderCollateralBounds(ctx context.Context, pieceSize abi.PaddedPieceSize, verified bool) (*StateDealProviderCollateralBounds, error) {
	switch resp, err := f.client.Call(ctx, methodFilStateDealProviderCollateralBounds, pieceSize, verified, nil); {
	case err != nil:
		return nil, err
	case resp.Error != nil:
		return nil, resp.Error
	default:
		var b StateDealProviderCollateralBounds
		if err := resp.GetObject(&b); err != nil {
			return nil, err
		}
		return &b, nil
	}
}

func (f *Telefil) StateMinerInfo(ctx context.Context, sp address.Address) (*peer.AddrInfo, error) {
	switch resp, err := f.client.Call(ctx, methodFilStateMinerInfo, sp.String(), nil); {
	case err != nil:
		return nil, err
	case resp.Error != nil:
		return nil, resp.Error
	default:
		var mi struct {
			PeerId     string   `json:"PeerId"`
			Multiaddrs []string `json:"Multiaddrs"`
		}
		if err = resp.GetObject(&mi); err != nil {
			return nil, err
		}
		if len(mi.PeerId) == 0 && len(mi.Multiaddrs) == 0 {
			return nil, nil
		}
		var ai peer.AddrInfo
		if len(mi.PeerId) != 0 {
			var err error
			if ai.ID, err = peer.Decode(mi.PeerId); err != nil {
				return nil, err
			}
		}
		if len(mi.Multiaddrs) != 0 {
			ai.Addrs = make([]multiaddr.Multiaddr, 0, len(mi.Multiaddrs))
			for _, s := range mi.Multiaddrs {
				mb, err := base64.StdEncoding.DecodeString(s)
				if err != nil {
					return nil, err
				}
				addr, err := multiaddr.NewMultiaddrBytes(mb)
				if err != nil {
					return nil, err
				}
				ai.Addrs = append(ai.Addrs, addr)
			}
		}
		return &ai, nil
	}
}

func (f *Telefil) StateMinerPower(ctx context.Context, sp address.Address) (*StateMinerPower, error) {
	switch resp, err := f.client.Call(ctx, methodFilStateMinerPower, sp, nil); {
	case err != nil:
		return nil, err
	case resp.Error != nil:
		return nil, resp.Error
	default:
		var mp StateMinerPower
		if err = resp.GetObject(&mp); err != nil {
			return nil, err
		}
		return &mp, nil
	}
}

func (f *Telefil) StateListMiners(ctx context.Context) ([]address.Address, error) {
	switch resp, err := f.client.Call(ctx, methodFilStateListMiners, nil); {
	case err != nil:
		return nil, err
	case resp.Error != nil:
		return nil, resp.Error
	default:
		var mids []address.Address
		if err = resp.GetObject(&mids); err != nil {
			return nil, err
		}
		return mids, nil
	}
}
