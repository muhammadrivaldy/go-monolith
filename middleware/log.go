package middleware

import (
	"backend/logs"
	"backend/util"
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

func WrapUseCase(ctx context.Context, req interface{}, f func() (interface{}, util.Error)) (res interface{}, errs util.Error) {

	logs.Logging.Info(ctx, wrapDataIntoJSON(req), zap.String("type", "request"))

	res, errs = f()
	if errs.IsError() {
		return res, errs
	}

	logs.Logging.Info(ctx, wrapDataIntoJSON(res), zap.String("type", "response"))

	return res, util.ErrorMapping(nil)

}

func wrapDataIntoJSON(data interface{}) string {

	if data != nil {

		result, _ := json.Marshal(data)

		var dataMap map[string]interface{}

		json.Unmarshal(result, &dataMap)

		for _, i := range []string{"password", "token", "refresh_token"} {
			_, exist := dataMap[i]
			if exist {
				dataMap[i] = "*****"
			}
		}

		result, _ = json.Marshal(dataMap)

		return string(result)

	}

	return ""

}
