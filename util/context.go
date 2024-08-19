package util

import (
	"context"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
)

var AttributesJWT = []string{"user_id", "email", "exp", "user_type", "store_id", "type"}

const (
	KeyContextUserId   goutil.KeyContext = "user_id"
	KeyContextEmail    goutil.KeyContext = "email"
	KeyContextExp      goutil.KeyContext = "exp"
	KeyContextUserType goutil.KeyContext = "user_type"
	KeyContextStoreId  goutil.KeyContext = "store_id"
	KeyContextType     goutil.KeyContext = "type"
)

type AttributesInContext struct {
	UserId   int64
	Email    string
	Exp      time.Time
	UserType int
	StoreID  int
	Type     string
}

func GetContext(ctx context.Context) AttributesInContext {

	var payload AttributesInContext

	if ctx != nil {

		if ctx.Value(KeyContextUserId) != nil {
			payload.UserId = int64(ctx.Value(KeyContextUserId).(float64))
		}

		if ctx.Value(KeyContextEmail) != nil {
			payload.Email = ctx.Value(KeyContextEmail).(string)
		}

		if ctx.Value(KeyContextExp) != nil {
			payload.Exp = time.Unix(int64(ctx.Value(KeyContextExp).(float64)), 0)
		}

		if ctx.Value(KeyContextUserType) != nil {
			payload.UserType = int(ctx.Value(KeyContextUserType).(float64))
		}

		if ctx.Value(KeyContextStoreId) != nil {
			payload.StoreID = int(ctx.Value(KeyContextStoreId).(float64))
		}

		if ctx.Value(KeyContextType) != nil {
			payload.Type = ctx.Value(KeyContextType).(string)
		}

	}

	return payload

}

type endpointInformation struct {
	Method   string
	Endpoint string
}

func getEndpointInformationFromContext(ctx context.Context) endpointInformation {

	var payload endpointInformation

	if ctx != nil {

		if ctx.Value(goutil.KeyMethod) != nil {
			payload.Method = ctx.Value(goutil.KeyMethod).(string)
		}

		if ctx.Value(goutil.KeyEndpoint) != nil {
			payload.Endpoint = ctx.Value(goutil.KeyEndpoint).(string)
		}

	}

	return payload

}
