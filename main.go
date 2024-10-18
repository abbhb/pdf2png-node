package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/gen2brain/go-fitz"
	"github.com/nfnt/resize"
)

func main() {
	// 打开PDF文件
	doc, err := fitz.New("computernet.pdf")
	if err != nil {
		fmt.Printf("Error opening PDF file: %v\n", err)
		return
	}
	defer doc.Close()

	// 获取页面数
	pageCount := doc.NumPage()
	fmt.Printf("Total pages: %d\n", pageCount)

	// 创建空白图像
	const thumbSize = 600 // 每个缩略图的宽高
	const gridSize = 3    // 九宫格为3x3
	gridWidth := thumbSize * gridSize
	gridHeight := thumbSize * gridSize

	gridImage := image.NewRGBA(image.Rect(0, 0, gridWidth, gridHeight))

	// 填充空白背景
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(gridImage, gridImage.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// 迭代每一页的缩略图
	for i := 0; i < pageCount && i < gridSize*gridSize; i++ {
		img, err := doc.Image(i)
		if err != nil {
			fmt.Printf("Error rendering page %d: %v\n", i+1, err)
			continue
		}

		// 缩放图像到指定缩略图大小
		thumb := resize.Resize(thumbSize, thumbSize, img, resize.Lanczos3)

		// 计算九宫格中缩略图的位置
		row := i / gridSize
		col := i % gridSize
		x := col * thumbSize
		y := row * thumbSize

		// 粘贴缩略图
		rect := image.Rect(x, y, x+thumbSize, y+thumbSize)
		draw.Draw(gridImage, rect, thumb, image.Point{0, 0}, draw.Over)
	}

	// 保存为PNG文件
	outFile, err := os.Create("output1.png")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outFile.Close()

	if err := png.Encode(outFile, gridImage); err != nil {
		fmt.Printf("Error encoding PNG: %v\n", err)
		return
	}

	fmt.Println("Thumbnail grid saved as output.png")
}
