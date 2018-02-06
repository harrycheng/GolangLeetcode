package library

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > len(m.musics) {
		return nil, errors.New("index error!")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (index int, music *MusicEntry) {
	if len(m.musics) == 0 {
		return 0, nil
	}

	for index, m := range m.musics {
		if m.Name == name {
			return index, &m
		}
	}

	return 0, nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
	fmt.Println(music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index > len(m.musics)-1 {
		return nil
	}
	removedMusic := &m.musics[index]
	if index < len(m.musics)-1 { // index (0, len - 1)
		m.musics = append(m.musics[:index], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	index, music := m.Find(name)
	if music == nil {
		fmt.Println("not found...", name)
		return nil
	}

	return m.Remove(index)
}
