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
