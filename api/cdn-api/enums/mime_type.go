package enums

import "errors"

type MIMEType string

const (
	ImageGIF  MIMEType = "image/gif"
	ImageJPEG MIMEType = "image/jpeg"
	ImagePNG  MIMEType = "image/png"
	ImageWEBP MIMEType = "image/webp"
	PDF       MIMEType = "application/pdf"
)

func GetFileExtension(mimeType string) string {
	switch MIMEType(mimeType) {
	case ImageGIF:
		return "gif"
	case ImageJPEG:
		return "jpg"
	case ImagePNG:
		return "png"
	case ImageWEBP:
		return "webp"
	case PDF:
		return "pdf"
	default:
		return ""
	}
}

// if mimtype Not in the enum list, throw error
func AllowedMimeTypes(mimeType string) error {
	switch MIMEType(mimeType) {
	case ImageGIF, ImageJPEG, ImagePNG, ImageWEBP, PDF:
		return nil
	default:
		return errors.New("invalid mime type")
	}
}
