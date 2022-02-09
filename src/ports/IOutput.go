package ports

type IOutput interface {
	createBMP(data []byte, w int, h int, fName string)
}
