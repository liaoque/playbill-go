package playbill

import (
	vips "github.com/davidbyttow/govips/v2/vips"
	"palybill/app/model"
)

func Load(data model.Poster, changeData map[string]interface{}) {

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
	return vips.ImageTypeSVG

	//	if (!$rows->data->opacity) {
	//		return Image::black($rows->data->width / $rows->data->zoom, $rows->data->height / $rows->data->zoom);
	//	}
	//
	//	$width = $rows->data->width / $rows->data->zoom;
	//	$height = $rows->data->height / $rows->data->zoom;
	//	$render = <<<EOF
	//	<svg viewBox="0 0 {width} {height}">
	//	<rect x="0" y="0"
	//	height="{height}"
	//	width="{width}"
	//	fill-opacity="0"
	//	></rect>
	//	</svg>
	//		EOF;
	//	$render = strtr($render, [
	//	'{width}' => $width,
	//		'{height}' => $height,
	//]);
	//	return Image::svgload_buffer($render
}
