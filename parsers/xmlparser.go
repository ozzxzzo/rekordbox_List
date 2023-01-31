package parsers

import (
	"encoding/xml"
	"os"
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
func ParseCollection(collection []Collection) map[string]CollectionTrack {
	trackMap := make(map[string]CollectionTrack)

	for _, playlist := range collection {
		for _, track := range playlist.CollectionTrack {
			trackMap[track.TrackID] = track
		}
	}

	return trackMap
}
