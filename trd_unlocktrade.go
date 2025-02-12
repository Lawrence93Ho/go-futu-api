package futuapi

import (
	"context"

	"github.com/Lawrence93Ho/go-futu-api/pb/trdcommon"
	"github.com/Lawrence93Ho/go-futu-api/pb/trdunlocktrade"
	"github.com/Lawrence93Ho/go-futu-api/protocol"
)

const (
	ProtoIDTrdUnlockTrade = 2005 //Trd_UnlockTrade	解锁或锁定交易
)

// 解锁交易
func (api *FutuAPI) UnlockTrade(ctx context.Context, unlock bool, pwd string, firm trdcommon.SecurityFirm) error {
	req := trdunlocktrade.Request{
		C2S: &trdunlocktrade.C2S{
			Unlock: &unlock,
		},
	}
	if pwd != "" {
		req.C2S.PwdMD5 = &pwd
	}
	if firm != 0 {
		req.C2S.SecurityFirm = (*int32)(&firm)
	}
	ch := make(trdunlocktrade.ResponseChan)
	if err := api.get(ProtoIDTrdUnlockTrade, &req, ch); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return ErrChannelClosed
		}
		return protocol.Error(resp)
	}
}
