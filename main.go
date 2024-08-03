package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {

	playlistPath := flag.String("playlist", "", "Path to the playlist")
	outputPath := flag.String("output", "", "Path to the output directory")

	flag.Parse()

	if *playlistPath == "" || *outputPath == "" {
		fmt.Println("Usage: parser -playlist <path to playlist> -output <path to output directory>")
		return
	}

	playlistDirPath := filepath.Dir(*playlistPath)

	allChannels, err := parsePlaylist(*playlistPath)
	if err != nil {
		fmt.Println("Error parsing playlist:", err)
		return
	}

	for _, channel := range allChannels {
		downloadPlaylistFile(channel, playlistDirPath)
	}

	availableChannelNames, _ := getAvailableChannelNames(playlistDirPath)
	availableChannels := filterAvailableChannels(allChannels, availableChannelNames)

	for _, channel := range availableChannels {
		setMediaURLs(&channel, playlistDirPath)
		saveChannelInfo(channel, playlistDirPath)
	}

}
