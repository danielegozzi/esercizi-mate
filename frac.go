package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"
)

func frac(w http.ResponseWriter, r *http.Request) {
	var err error
	var columns, rows, scale, filled int
	var columnsParam, rowsParam, scaleParam, filledParam []string
	var etag string
	var present bool
	query := r.URL.Query()
	scaleParam, present = query["scale"]
	if present && len(scaleParam) > 0 {
		scale, err = strconv.Atoi(scaleParam[0])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	} else {
		scale = 20
	}
	if scale < 10 {
		scale = 10
	}
	if scale > 30 {
		scale = 30
	}
	columnsParam, present = query["columns"]
	if present && len(columnsParam) > 0 {
		err = nil
		columns, err = strconv.Atoi(columnsParam[0])
	}
	if !present || len(columnsParam) <= 0 || err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if columns < 1 {
		columns = 1
	}
	if columns > 20 {
		columns = 20
	}
	rowsParam, present = query["rows"]
	if present && len(rowsParam) > 0 {
		err = nil
		rows, err = strconv.Atoi(rowsParam[0])
	}
	if !present || len(rowsParam) <= 0 || err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if rows < 1 {
		rows = 1
	}
	if rows > 20 {
		rows = 20
	}
	filledParam, present = query["filled"]
	if present && len(filledParam) > 0 {
		filled, err = strconv.Atoi(filledParam[0])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	} else {
		filled = 0
	}
	if filled < 0 {
		filled = 0
	}
	if filled > rows*columns {
		filled = rows * columns
	}
	etag = fmt.Sprintf("%d.%d.%d.%d", scale, rows, columns, filled)
	w.Header().Set("ETag", etag)
	w.Header().Set("Content-Type", "image/png")
	grid := drawGrid(rows, columns, scale, filled)
	png.Encode(w, grid)
}

func drawSquare(image *image.RGBA, x int, y int, scale int, fillColor color.Color) {
	for i := 0; i < scale; i++ {
		image.Set(0+x*scale, i+y*scale, color.Black)
		image.Set(i+x*scale, 0+y*scale, color.Black)
		image.Set(scale-1+x*scale, i+y*scale, color.Black)
		image.Set(i+x*scale, scale-1+y*scale, color.Black)
	}
	for i := 1; i < scale-1; i++ {
		for j := 1; j < scale-1; j++ {
			image.Set(i+x*scale, j+y*scale, fillColor)
		}
	}
}

func drawGrid(rows int, cols int, scale int, filled int) *image.RGBA {
	var fillCounter int = 1
	upLeft := image.Point{0, 0}
	lowRight := image.Point{cols * scale, rows * scale}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if fillCounter <= filled {
				drawSquare(img, c, r, scale, color.RGBA{128, 128, 128, 255})
				fillCounter++
			} else {
				drawSquare(img, c, r, scale, color.White)
			}
		}
	}
	return img
}
