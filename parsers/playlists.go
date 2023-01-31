package parsers

type DJPlaylists struct {
	Version    string       `xml:"Version,attr"`
	Product    []Product    `xml:"PLAYLISTS>PRODUCT"`
	Collection []Collection `xml:"COLLECTION"`
	Node       []Node       `xml:"PLAYLISTS>NODE"`
}

type Product struct {
	Name    string `xml:"Name,attr"`
	Version string `xml:"Version,attr"`
	Company string `xml:"Company,attr"`
}

type Collection struct {
	CollectionTrack []CollectionTrack `xml:"COLLECTION>TRACK"`
}

type CollectionTrack struct {
	Name                string `xml:"Name"`
	Artist              string `xml:"Artist"`
	Composer            string `xml:"Composer"`
	Album               string `xml:"Album"`
	Genre               string `xml:"Genre"`
	Kind                string `xml:"Kind"`
	Size                int64  `xml:"Size"`
	TotalTime           int64  `xml:"TotalTime"`
	TrackNumber         int    `xml:"TrackNumber"`
	Year                int    `xml:"Year"`
	DateModified        string `xml:"DateModified"`
	DateAdded           string `xml:"DateAdded"`
	BitRate             int    `xml:"BitRate"`
	SampleRate          int    `xml:"SampleRate"`
	PlayCount           int    `xml:"PlayCount"`
	PlayDate            int64  `xml:"PlayDate"`
	PlayDateUTC         string `xml:"PlayDateUTC"`
	SkipCount           int    `xml:"SkipCount"`
	SkipDate            string `xml:"SkipDate"`
	Rating              int    `xml:"Rating"`
	AlbumRating         int    `xml:"AlbumRating"`
	AlbumRatingComputed int    `xml:"AlbumRatingComputed"`
	PersistentID        string `xml:"PersistentID"`
	TrackType           string `xml:"TrackType"`
	Location            string `xml:"Location"`
	FileFolderCount     int    `xml:"FileFolderCount"`
	LibraryFolderCount  int    `xml:"LibraryFolderCount"`
}

type Node struct {
	Name    string  `xml:"Name,attr"`
	Type    string  `xml:"Type,attr"`
	KeyType string  `xml:"KeyType,attr"`
	Entries string  `xml:"Entries,attr"`
	Tracks  []Track `xml:"TRACK"`
}

type Track struct {
	Key string `xml:"Key,attr"`
}
