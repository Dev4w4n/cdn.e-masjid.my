package model

type Request struct {
	Data      []byte `json:"data"`      // The image data (PNG or JPG)
	DataType  string `json:"dataType"`  // Type of the image data ("png" or "jpg")
	ImageName string `json:"imageName"` // Name of the image
	Namespace string `json:"namespace"` // Namespace of the image
	ImageType int    `json:"imageType"` // Type of the image (1: main.png, 2: public.png, 3: private.png)
}
