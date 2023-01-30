package parsers

import (
	"encoding/xml"
	"os"
)

// ParseXML function to parse an XML file
func ParseXML(filePath string) (*Playlists, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	var playlists Playlists
	if err := xml.NewDecoder(xmlFile).Decode(&playlists); err != nil {
		return nil, err
	}

	return &playlists, nil
}
