package chunk

import (
	"context"
	"time"
)

// ChannelChunk 从 channel 中读取指定数量的数据
func ChannelChunk[T any](src chan T, size int) []T {
	result := make([]T, 0, size)
	i := 0
	for data := range src {
		result = append(result, data)
		i++
		if i == size {
			break
		}
	}

	return result
}

// ChannelChunkTimeout 从 channel 中读取指定数量的数据
func ChannelChunkTimeout[T any](src chan T, size int, timeout time.Duration) []T {
	result := make([]T, 0, size)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case data, ok := <-src:
			if !ok {
				return result
			}

			result = append(result, data)
			if len(result) == size {
				return result
			}
		case <-ctx.Done():
			return result
		}
	}
}
