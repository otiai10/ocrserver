package controllers

import (
	"github.com/otiai10/gosseract/v2"
)

func scanText(client *gosseract.Client) (string, float64, error) {
	boxes, err := client.GetBoundingBoxes(gosseract.RIL_BLOCK)
	if err != nil {
		return "", 0.0, err
	}
	var content string
	var confidence float64
	for _, box := range boxes {
		content = content + box.Word
		confidence += box.Confidence
	}
	return content, confidence / float64(len(boxes)), nil
}
