package parsers

import (
	"encoding/xml"
)

type DJPlaylists struct {
	XMLName    xml.Name     `xml:"DJ_PLAYLISTS"`
	Version    string       `xml:"Version,attr"`
	Product    []Product    `xml:"PRODUCT"`
	Collection []Collection `xml:"COLLECTION"`
	Node       []Node       `xml:"NODE"`
}

type Product struct {
	XMLName xml.Name `xml:"PRODUCT"`
	Name    string   `xml:"Name,attr"`
	Version string   `xml:"Version,attr"`
	Company string   `xml:"Company,attr"`
}

type Collection struct {
	XMLName         xml.Name          `xml:"COLLECTION"`
	Entries         string            `xml:"Entries,attr"`
	CollectionTrack []CollectionTrack `xml:"TRACK"`
}

type CollectionTrack struct {
	XMLName             xml.Name `xml:"COLLECTION>TRACK“`
	Name                string   `xml:"Name,attr"`
	Artist              string   `xml:"Artist,attr"`
	Composer            string   `xml:"Composer,attr"`
	Album               string   `xml:"Album,attr"`
	Genre               string   `xml:"Genre,attr"`
	Kind                string   `xml:"Kind,attr"`
	Size                int64    `xml:"Size,attr"`
	TotalTime           int64    `xml:"TotalTime,attr"`
	TrackNumber         int      `xml:"TrackNumber,attr"`
	Year                int      `xml:"Year,attr"`
	DateModified        string   `xml:"DateModified,attr"`
	DateAdded           string   `xml:"DateAdded,attr"`
	BitRate             int      `xml:"BitRate,attr"`
	SampleRate          int      `xml:"SampleRate,attr"`
	PlayCount           int      `xml:"PlayCount,attr"`
	PlayDate            int64    `xml:"PlayDate,attr"`
	PlayDateUTC         string   `xml:"PlayDateUTC,attr"`
	SkipCount           int      `xml:"SkipCount,attr"`
	SkipDate            string   `xml:"SkipDate,attr"`
	Rating              int      `xml:"Rating,attr"`
	AlbumRating         int      `xml:"AlbumRating,attr"`
	AlbumRatingComputed int      `xml:"AlbumRatingComputed,attr"`
	PersistentID        string   `xml:"PersistentID,attr"`
	TrackType           string   `xml:"TrackType,attr"`
	Location            string   `xml:"Location,attr"`
	FileFolderCount     int      `xml:"FileFolderCount,attr"`
	LibraryFolderCount  int      `xml:"LibraryFolderCount,attr"`
}

type Node struct {
	XMLName xml.Name `xml:"NODE"`
	Name    string   `xml:"Name,attr"`
	Type    string   `xml:"Type,attr"`
	KeyType string   `xml:"KeyType,attr"`
	Entries string   `xml:"Entries,attr"`
	Tracks  []Track  `xml:"TRACK"`
}

type Track struct {
	Key string `xml:"Key,attr"`
}
