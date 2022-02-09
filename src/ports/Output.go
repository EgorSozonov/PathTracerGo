package ports

import (
	"os"
	"fmt"
)

func CreateBMP(data []byte, w int, h int, fName string) {
	fileSize := 54 + 3*w*h
	img := make([]byte, 3*w*h)

	for i := 0; i < w; i += 1 {
		for j := 0; j < h; j += 1 {
			indSource := 3 * (j*w + i)
			indTarget := 3 * (j*w + i)
			img[indTarget    ] = data[indSource + 2];
			img[indTarget + 1] = data[indSource + 1];
			img[indTarget + 2] = data[indSource    ];
			
			

		}
	}
	bmpHeader := [14]byte{byte('B'), byte('M'), 0, 0,  0, 0, 0, 0,  0, 0, 54, 0,  0, 0,}
    bmpHeader[2] = byte(fileSize);
    bmpHeader[3] = byte(fileSize >> 8);
    bmpHeader[4] = byte(fileSize >> 16);
    bmpHeader[5] = byte(fileSize >> 24);

	bmpInfoHeader := [40]byte {
                40, 0, 0, 0,  0, 0, 0, 0,  0, 0, 0, 0,  1, 0, 24, 0,  0, 0, 0, 0,  0, 0, 0, 0,  0, 0, 0, 0,  0, 0, 0, 0,
                0, 0, 0, 0,  0, 0, 0, 0,
            };
    bmpInfoHeader[ 4] = byte(w      );
    bmpInfoHeader[ 5] = byte(w >>  8);
    bmpInfoHeader[ 6] = byte(w >> 16);
    bmpInfoHeader[ 7] = byte(w >> 24);
    bmpInfoHeader[ 8] = byte(h      );
    bmpInfoHeader[ 9] = byte(h >>  8);
    bmpInfoHeader[10] = byte(h >> 16);
    bmpInfoHeader[11] = byte(h >> 24);

	f, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        fmt.Println("Error opening file " + fName)
		return
    }
    defer func () { 
		if errCl := f.Close(); errCl != nil {
        	fmt.Println("Error closing file " + fName)
		}
    }()

	bmpPad := [3]byte{ 0, 0, 0 }
	f.Write(bmpHeader[:])
	f.Write(bmpInfoHeader[:])
	lenPad := (4 - (w*3)%4) %4;
	if lenPad > 0 {
		for i := 0; i < h; i += 1 {
			f.Write(img[3*i*w : (3*(i + 1)*w)]);
			f.Write(bmpPad[:])
		}
	} else {
		for i := 0; i < h; i += 1 {
			f.Write(img[3*i*w : (3*(i + 1)*w)]);
		}
	}
}
