package parsers

import (
	"encoding/xml"
	"os"

	"github.com/ozzxzzo/rekordbox_List/parsers"
)

// ParseXML function to parse an XML file
func ParseXML(filePath string) (*DJPlaylists, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	var playlists DJPlaylists
	if err := xml.NewDecoder(xmlFile).Decode(&playlists); err != nil {
		return nil, err
	}

	return &playlists, nil
}

// Parse the Collection structure
func ParseCollection(collection []parsers.Collection) map[string]parsers.CollectionTrack {
	trackMap := make(map[string]parsers.CollectionTrack)

	for _, playlist := range collection {
		for _, track := range playlist.CollectionTrack {
			trackMap[track.TrackID] = track
		}
	}

	return trackMap
}
