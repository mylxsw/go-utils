package chunk

import (
	"testing"
	"time"

	"github.com/mylxsw/go-utils/assert"
)

func TestChannelChunk(t *testing.T) {
	src := make(chan int)
	go func() {
		defer close(src)
		for i := 0; i < 5; i++ {
			src <- i
			t.Logf("write %d", i)
		}
	}()

	assert.Equal(t, []int{0, 1}, ChannelChunk(src, 2))
	assert.Equal(t, []int{2, 3}, ChannelChunk(src, 2))
	assert.Equal(t, []int{4}, ChannelChunk(src, 2))
	assert.Equal(t, []int{}, ChannelChunk(src, 2))
}

func TestChannelChunkTimeout(t *testing.T) {
	src := make(chan int)
	go func() {
		defer close(src)
		for i := 0; i < 5; i++ {
			src <- i
			t.Logf("write %d", i)
			time.Sleep(2 * time.Millisecond)
		}
	}()

	assert.Equal(t, []int{0, 1}, ChannelChunkTimeout(src, 2, 6*time.Millisecond))
	assert.Equal(t, []int{2}, ChannelChunkTimeout(src, 2, 3*time.Millisecond))
	assert.Equal(t, []int{3, 4}, ChannelChunkTimeout(src, 2, 5*time.Millisecond))
	assert.Equal(t, []int{}, ChannelChunkTimeout(src, 2, 1*time.Millisecond))
}
