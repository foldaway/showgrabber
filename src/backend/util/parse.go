package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bottleneckco/showgrabber/src/backend/graph/model"
)

func ParseNewznab(title string) (*model.ParsedMetadata, error) {
	var result model.ParsedMetadata
	var err error
	var isResultAvailable = false

	var seasonRegex = regexp.MustCompile(`[Ss](?P<SeasonNumber>\d+)`)
	var seasonRaw = seasonRegex.FindStringSubmatch(title)

	if len(seasonRaw) > 0 {
		seasonNum, err := strconv.Atoi(seasonRaw[1])
		if err != nil {
			return nil, err
		}
		result.SeasonNumber = &seasonNum
		isResultAvailable = true
	}

	var episodeRegex = regexp.MustCompile(`[Ee](?P<EpisodeNumber>\d+)`)
	var episodeRaw = episodeRegex.FindStringSubmatch(title)

	if len(episodeRaw) > 0 {
		episodeNum, err := strconv.Atoi(episodeRaw[1])
		if err != nil {
			return nil, err
		}
		result.EpisodeNumber = &episodeNum
		isResultAvailable = true
	}

	var resolutionLiteralRegex = regexp.MustCompile(`(\d{3,4})x(\d{3,4})`)
	var resolutionLiteralRaw = resolutionLiteralRegex.FindStringSubmatch(title)

	if len(resolutionLiteralRaw) > 0 {
		var res = fmt.Sprintf("%sp", resolutionLiteralRaw[2])
		result.Resolution = &res
		isResultAvailable = true
	}

	var resolutionHDTVRegex = regexp.MustCompile(`(\d+[PpIi])`)
	var resolutionHDTVRaw = resolutionHDTVRegex.FindStringSubmatch(title)

	if len(resolutionHDTVRaw) > 0 {
		var res = resolutionHDTVRaw[1]
		result.Resolution = &res
		isResultAvailable = true
	}

	var videoCodecRegex = regexp.MustCompile(`(?i)(?:h|x)?((?:264|265)|xvid)`)
	var videoCodecRaw = videoCodecRegex.FindStringSubmatch(title)

	if len(videoCodecRaw) > 0 {
		var vCodec = videoCodecRaw[1]
		result.VideoCodec = &vCodec
		isResultAvailable = true
	}

	if strings.Contains(title, "HEVC") {
		var vCodec = "265"
		result.VideoCodec = &vCodec
		isResultAvailable = true
	}

	var audioCodecRegex = regexp.MustCompile(`(?i)(aac|ac3)`)
	var audioCodecRaw = audioCodecRegex.FindStringSubmatch(title)

	if len(audioCodecRaw) > 0 {
		var aCodec = audioCodecRaw[1]
		result.AudioCodec = &aCodec
		isResultAvailable = true
	}

	var releaseFormatRegex = regexp.MustCompile(`(?i)(WEBDL|WEBRip|WEB|HDTV|Xvid)`)
	var releaseFormatRaw = releaseFormatRegex.FindStringSubmatch(strings.ReplaceAll(title, "-", ""))

	if len(releaseFormatRaw) > 0 {
		var format = strings.ToUpper(releaseFormatRaw[1])
		result.ReleaseFormat = &format
		isResultAvailable = true
	}

	if !isResultAvailable {
		return nil, err
	}

	return &result, err
}
