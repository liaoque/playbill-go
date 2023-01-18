package playbill

import (
	"database/sql"
	vips "github.com/davidbyttow/govips/v2/vips"
	"palybill/app/model"
	"strings"
)

func Load(data model.Poster, changeData map[string]interface{}) {

	image, err := createArea(data)
	if err != nil {
		return
	}

	if data.Data["backgroundImage"] != nil || data.Data["background"] != nil {
		if data.Data["background"] != nil {

		} else {

		}
	} else {

	}

	objects := data.Data["objects"]

}

func createArea(data model.Poster) (*vips.ImageRef, error) {
	width := data.Data["width"].(int)
	height := data.Data["height"].(int)
	zoom := data.Data["zoom"].(int)
	if data.Data["opacity"] != 0 {
		return vips.Black(
			width/zoom,
			height/zoom,
		)
	}
	width = width / zoom
	height = height / zoom

	render := "<svg viewBox=\"0 0 {width} {height}\">\n" +
		"<rect x=\"0\" y=\"0\"\n\theight=\"{height}\"\n\twidth=\"{width}\"\n\tfill-opacity=\"0\"\n>" +
		"</rect>\n" +
		"</svg>"

	render = strings.ReplaceAll(render, "{width}", string(width))
	render = strings.ReplaceAll(render, "{height}", string(height))

	return vips.LoadImageFromBuffer([]byte(render), &vips.ImportParams{})
}
