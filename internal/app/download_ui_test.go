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

func TestDownloadViewHidesFinishedJobs(t *testing.T) {
	m := model{
		width: 80,
		downloads: visibleDownloads([]store.Job{
			{ID: "1", Title: "Now Serving", Status: "done", Percent: 100},
			{ID: "2", Title: "Experience", Status: "running", Percent: 40, Speed: "1MiB/s", ETA: "00:03"},
			{ID: "3", Title: "Broken", Status: "failed", Error: "no audio"},
		}),
		downloadIndex: 0,
	}
	view := m.renderDownloads()
	if strings.Contains(view, "Now Serving") || strings.Contains(view, "in the cassette") {
		t.Fatalf("finished download stayed visible: %s", view)
	}
	for _, want := range []string{"Experience", "Broken"} {
		if !strings.Contains(view, want) {
			t.Fatalf("view %q missing %q", view, want)
		}
	}
}

func TestCompletionStillReloadsFromTheFullList(t *testing.T) {
	previous := []store.Job{{ID: "1", Title: "Experience", Status: "running"}}
	next := []store.Job{{ID: "1", Title: "Experience", Status: "done"}}
	if !newlyFinished(previous, next) {
		t.Fatal("completion was not noticed")
	}
	if len(visibleDownloads(next)) != 0 {
		t.Fatal("finished download stayed in the visible list")
	}
}
