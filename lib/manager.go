package lib

import (
	"app/music/model"
	"errors"
)

type MusicManager struct {
	musics []model.MusicEntry
}

func NewMusuicManager() *MusicManager {
	return &MusicManager{make([]model.MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *model.MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *model.MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}

	return nil
}

func (m *MusicManager) Add(music *model.MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *model.MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	// remove from m.musics
	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = []model.MusicEntry{}
	} else {
		m.musics = m.musics[:index-1]
	}

	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *model.MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for i, music := range m.musics {
		if music.Name == name {
			return m.Remove(i)

		}
	}

	return nil
}
