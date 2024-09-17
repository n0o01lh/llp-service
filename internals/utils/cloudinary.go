package utils

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2/log"
)

type Cloudinary struct {
}

var cld *cloudinary.Cloudinary

func GetCloudinaryInstance() *cloudinary.Cloudinary {
	var err error
	if cld == nil {
		cld, err = cloudinary.NewFromParams("deiz84fxy", "596158677423497", "boDgQ86Sdlhe4umhy7HizwBgmZw")
		if err != nil {
			log.Error(err)
		}
	}
	return cld

}

func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context, image string) (string, error) {

	resp, err := cld.Upload.Upload(ctx, image, uploader.UploadParams{
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}
