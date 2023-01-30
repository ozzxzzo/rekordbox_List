package parsers

type Playlists struct {
	Nodes []Node `xml:"NODE"`
}

type Node struct {
	Name    string  `xml:"Name,attr"`
	Type    string  `xml:"Type,attr"`
	Entries string  `xml:"Entries,attr"`
	Tracks  []Track `xml:"TRACK"`
}

type Track struct {
	Key string `xml:"Key,attr"`
}
