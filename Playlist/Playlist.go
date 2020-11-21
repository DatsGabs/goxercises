package main

import "fmt"

type song struct {
	name     string
	artist   string
	previous *song
	next     *song
}

type playlist struct {
	name       string
	head       *song
	tail       *song
	nowPlaying *song
}

func (p *playlist) addSong(name, artist string) error {
	song := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = song
	} else {
		current := p.tail
		current.next = song
		song.previous = current
	}
	p.tail = song
	return nil
}

func (p playlist) showAllSongs() error {
	current := p.head
	if current == nil {
		fmt.Println("The playlist is empty")
		return nil
	}

	fmt.Println(current.name)
	for current.next != nil {
		current = current.next
		fmt.Println(current.name)
	}
	return nil
}

func (p *playlist) nextSong() error {
	current := p.nowPlaying
	if current.next != nil {
		p.nowPlaying = current.next
	} else {
		p.nowPlaying = p.head
	}
	return nil
}

func (p *playlist) previousSong() error {
	current := p.nowPlaying
	if current.previous != nil {
		p.nowPlaying = current.previous
	} else {
		p.nowPlaying = p.tail
	}
	return nil
}

func (p *playlist) startSong(name string) error {
	current := p.head

	if current == nil {
		fmt.Println("The playlist is empty")
		return nil
	}

	if current.name == name {
		p.nowPlaying = current
		return nil
	}
	for current.next != nil {
		current = current.next
		if current.name == name {
			p.nowPlaying = current
			return nil
		}
	}

	fmt.Println("Couldn't find the song")
	return nil
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func main() {
	playlist := createPlaylist("Songs")
	playlist.addSong("OKUPA", "WOS")
	playlist.addSong("CANGURO", "WOS")
	playlist.addSong("Human", "Rag'n'bone man")
	playlist.startSong("CANGURO")
	fmt.Println(playlist.nowPlaying)
}
