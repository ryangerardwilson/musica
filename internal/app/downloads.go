package app

import (
	"fmt"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ryangerardwilson/musica/internal/download"
	"github.com/ryangerardwilson/musica/internal/store"
)

var enqueueDownload = func(playlistDir, filename, rawURL string) error {
	_, err := download.Enqueue(playlistDir, filename, rawURL)
	return err
}

var cancelDownload = download.Cancel

func (m *model) syncDownloads() {
	jobs, err := download.List()
	if err != nil {
		return
	}
	finished := newlyFinished(m.downloads, jobs)
	index := m.downloadIndex
	if index < 0 || index >= len(jobs) {
		index = runningIndex(jobs)
	}
	m.downloads = jobs
	if index >= len(jobs) {
		index = len(jobs) - 1
	}
	if index < 0 {
		index = 0
	}
	m.downloadIndex = index
	if finished && m.currentPlaylist != "" {
		m.reloadTracks()
	}
}

func newlyFinished(previous, next []store.Job) bool {
	if len(previous) == 0 {
		return false
	}
	was := make(map[string]string, len(previous))
	for _, job := range previous {
		was[job.ID] = job.Status
	}
	for _, job := range next {
		if job.Status == "done" && was[job.ID] != "" && was[job.ID] != "done" {
			return true
		}
	}
	return false
}

func runningIndex(jobs []store.Job) int {
	for i, job := range jobs {
		if job.Status == "running" {
			return i
		}
	}
	return 0
}

func (m model) renderDownloads() string {
	playing := ""
	if m.player != nil {
		if path, _ := m.player.State(); path != "" {
			playing = filepath.Base(path)
		}
	}
	return download.Render(m.downloads, m.downloadIndex, m.visualizerFrame, m.width, playing)
}

func (m model) downloadStatus() string {
	for _, job := range m.downloads {
		if job.Status != "running" && job.Status != "queued" {
			continue
		}
		if job.Status == "running" && job.Percent > 0 {
			return fmt.Sprintf("downloading %s %.0f%%", job.Title, job.Percent)
		}
		return "downloading " + job.Title
	}
	return ""
}

func (m model) updateDownloadKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.showDownloads = false
		return m, nil
	case "j", "down":
		m.moveDownload(1)
		return m, nil
	case "k", "up":
		m.moveDownload(-1)
		return m, nil
	case "x":
		m.cancelSelectedDownload()
		return m, nil
	case "y":
		m.beginYouTubeDownload()
		return m, nil
	default:
		return m, nil
	}
}

func (m *model) moveDownload(delta int) {
	if len(m.downloads) == 0 {
		return
	}
	m.downloadIndex += delta
	if m.downloadIndex < 0 {
		m.downloadIndex = 0
	}
	if m.downloadIndex >= len(m.downloads) {
		m.downloadIndex = len(m.downloads) - 1
	}
}

func (m *model) cancelSelectedDownload() {
	if m.downloadIndex < 0 || m.downloadIndex >= len(m.downloads) {
		m.status = "no download selected"
		return
	}
	job := m.downloads[m.downloadIndex]
	if err := cancelDownload(job.ID); err != nil {
		m.status = err.Error()
		return
	}
	m.downloads[m.downloadIndex].Status = "cancelled"
	m.status = "cancelled " + job.Title
}
