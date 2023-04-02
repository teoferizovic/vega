package controller

import (
	"context"

	"fmt"

	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/teo/vega/database"
)

func AddScore(p map[string]interface{}) (map[string]interface{}, error) {

	ctx := context.TODO()

	nickname := p["nickname"].(string)

	steps := p["steps"].(float64)

	//Validate data here in a production environment
	err := database.Redis.ZAdd(ctx, "app_users", &redis.Z{

		Score: steps,

		Member: nickname,
	}).Err()

	if err != nil {

		return nil, err

	}

	rank := database.Redis.ZRank(ctx, "app_users", p["nickname"].(string))

	if err != nil {

		return nil, err

	}

	response := map[string]interface{}{

		"data": map[string]interface{}{

			"nickname": p["nickname"].(string),

			"rank": rank.Val(),
		},
	}

	return response, nil

}

func GetScores(p map[string]interface{}) (map[string]interface{}, error) {

	ctx := context.TODO()

	start, err := strconv.ParseInt(fmt.Sprint(p["start"]), 10, 64)

	if err != nil {

		return nil, err

	}

	stop, err := strconv.ParseInt(fmt.Sprint(p["stop"]), 10, 64)

	if err != nil {

		return nil, err

	}

	total, err := database.Redis.ZCount(ctx, "app_users", "-inf", "+inf").Result() //int64

	if err != nil {

		return nil, err

	}

	scores, err := database.Redis.ZRevRangeWithScores(ctx, "app_users", start, stop).Result() //highest to lowest score

	if err != nil {

		return nil, err

	}

	data := []map[string]interface{}{}

	for _, z := range scores {

		record := map[string]interface{}{}

		rank := database.Redis.ZRank(ctx, "app_users", z.Member.(string))

		if err != nil {

			return nil, err

		}

		record["nickname"] = z.Member.(string)

		record["score"] = z.Score

		record["rank"] = rank.Val()

		data = append(data, record)

	}

	countPerRequest := stop - start + 1

	if stop == -1 {

		countPerRequest = total

	}

	response := map[string]interface{}{

		"data": data,

		"meta": map[string]interface{}{

			"start": start,

			"stop": stop,

			"per_request": countPerRequest,

			"total": total,
		},
	}

	return response, nil

}
