package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat, progress int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV......")

	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(" . ", p.progress)
		p.progress += 10
	}

	fmt.Println("Finished Playing WAV", source)
}
