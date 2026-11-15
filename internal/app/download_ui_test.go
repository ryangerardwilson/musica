package app

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ryangerardwilson/musica/internal/store"
	"github.com/ryangerardwilson/musica/internal/visualizer"
)

func TestDownloadViewAnimatesAndCancelLeavesPlayback(t *testing.T) {
	cancelled := ""
	previous := cancelDownload
	cancelDownload = func(id string) error {
		cancelled = id
		return nil
	}
	defer func() { cancelDownload = previous }()

	player := &fakeAudioPlayer{path: "/music/other.mp3", playing: true}
	m := model{
		width:         80,
		height:        24,
		showDownloads: true,
		player:        player,
		visualizer:    visualizer.New(4),
		downloads: []store.Job{
			{ID: "a", Title: "Blue Monday", Status: "running", Percent: 40, Speed: "1MiB/s", ETA: "00:03"},
			{ID: "b", Title: "Other", Status: "queued"},
		},
	}
	first := m.renderDownloads()
	m.visualizerFrame = 1
	second := m.renderDownloads()
	if first == second {
		t.Fatalf("download view did not animate:\n%s", first)
	}
	if !strings.Contains(first, "Blue Monday") || !strings.Contains(first, "playing  other.mp3") {
		t.Fatalf("view = %s", first)
	}

	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'x'}})
	updated := next.(model)
	if cancelled != "a" {
		t.Fatalf("cancelled %q", cancelled)
	}
	if updated.downloads[1].Status != "queued" {
		t.Fatal("cancel changed a different download")
	}
	if !player.IsPlaying() {
		t.Fatal("cancel stopped playback")
	}
	if !strings.Contains(updated.status, "cancelled Blue Monday") {
		t.Fatalf("status = %q", updated.status)
	}
}

func TestSpaceWhileDownloadingDoesNotCancel(t *testing.T) {
	player := &fakeAudioPlayer{path: "/music/other.mp3", playing: true}
	m := model{
		player:     player,
		visualizer: visualizer.New(4),
		downloads:  []store.Job{{ID: "a", Title: "Blue Monday", Status: "running"}},
	}
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{' '}})
	updated := next.(model)
	if updated.downloads[0].Status != "running" {
		t.Fatal("playback cancelled the download")
	}
	if !player.IsPaused() {
		t.Fatal("space did not pause playback")
	}
}

func TestUOpensDownloadView(t *testing.T) {
	next, _ := model{}.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'u'}})
	updated := next.(model)
	if !updated.showDownloads {
		t.Fatal("u did not open downloads")
	}
	view := updated.View()
	if !strings.Contains(view, "downloads") {
		t.Fatalf("view = %s", view)
	}
}
