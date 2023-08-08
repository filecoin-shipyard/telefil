package telefil_test

import (
	"context"
	"testing"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/masih/telefil"
)

func Test(t *testing.T) {
	const timeout = 10 * time.Second
	subject, err := telefil.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("StateDealProviderCollateralBounds", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)
		resp, err := subject.StateDealProviderCollateralBounds(ctx, 1, true)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(t.Name(), resp)
	})
	t.Run("ChainGetGenesis", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)
		resp, err := subject.ChainGetGenesis(ctx)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(t.Name(), resp)
	})
	t.Run("ChainHead", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)
		resp, err := subject.ChainHead(ctx)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(t.Name(), resp)
	})
	t.Run("StateListMiners", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)
		resp, err := subject.StateListMiners(ctx)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(t.Name(), resp[0:10])
	})
	t.Run("StateMinerInfo", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)
		addr, err := address.NewFromString("f01953925")
		if err != nil {
			t.Fatal(err)
		}
		resp, err := subject.StateMinerInfo(ctx, addr)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(t.Name(), resp)
	})
	t.Run("StateMinerPower", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		t.Cleanup(cancel)
		addr, err := address.NewFromString("f01953925")
		if err != nil {
			t.Fatal(err)
		}
		resp, err := subject.StateMinerPower(ctx, addr)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(t.Name(), resp)
	})
}
