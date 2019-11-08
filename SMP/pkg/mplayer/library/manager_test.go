package library

import "testing"

func TestOps ( t *testing.T){
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}

	if mm.len() != 0 {
		t.Error("not empty")
	}

	m0 := &MusicEntry{
		"1", "my heart", "cellion dion","http://qbox.me/1212","MP3"
	}

	mm.Add(m0)

	if mm.len() != 1 {
		t.Error("NewMusicManager Add failed")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("NewMusicManager find  failed")
	}

	m , err := mm.Get(0)
	if m == nil {
		t.Error("NewMusicManager get failed")
	}

	m := mm.Remove(0)
	if m == nil {
		t.Error("NewMusicManager remove failed")
	}
}