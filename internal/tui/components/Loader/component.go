package loader

import (
	"math"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/yashranjan1/relay/internal/tui/styles"
)

type TickMsg time.Time

type Loader struct {
	start     time.Time
	speed     time.Duration
	message   string
	direction int
	index     int
	width     int
}

func NewLoader(speed time.Duration, width int) *Loader {
	return &Loader{
		start:     time.Now(),
		speed:     speed,
		direction: 1,
		width:     width,
		message:   " Generating commit",
	}
}

func (l *Loader) SetMessage(message string) {
	l.message = message
}

func (l *Loader) Tick() tea.Cmd {
	return tea.Tick(l.speed, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (l *Loader) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case TickMsg:
		l.index += l.direction
		if l.index >= l.width-1 {
			l.direction = -1
		} else if l.index <= 0 {
			l.direction = 1
		}
		return l.Tick()
	}
	return nil
}

func (l *Loader) View() string {
	var result strings.Builder

	for i := 0; i < l.width; i++ {
		ch := "■"

		dist := float64(abs(i-l.index)) / float64(l.width)
		sin := math.Sin((1.0 - dist) * math.Pi / 2)
		alpha := math.Pow(sin, 8)

		char := styles.GetLoaderCharStyle(alpha)(ch)

		result.WriteString(char)

	}
	result.WriteString(styles.LoaderMessageStyle.Render(l.message))
	return result.String()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
