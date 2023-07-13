package main

import (
	"context"
	"time"
)

func IncrementalSum(ctx context.Context, start int, end int, msg chan Number) {
	memo := Number(start)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if start <= end {
				memo = sum(int(memo), 1)
				msg <- memo
			} else {
				return
			}

			time.Sleep(time.Second)
		}
	}
}
