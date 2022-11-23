package mmpaymkttransfers

import (
	"github.com/chanxuehong/wechat/mch/core"
)

// 查询代金券批次信息.
func QueryCouponStock(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(clt.APIBaseURL()+"/mmpaymkttransfers/query_coupon_stock", req)
}
