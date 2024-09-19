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

func GetCloudinaryInstance(ctx context.Context) *cloudinary.Cloudinary {

	if cld == nil {
		vaultClient := VaultConnection()
		cloudinary_credentials, err := vaultClient.KVv2("secret").Get(ctx, "cloudinary_credentials")

		if err != nil {
			log.Error("unable to read secret: %v", err)
		}

		cld, err = cloudinary.NewFromParams(
			cloudinary_credentials.Data["cloud"].(string),
			cloudinary_credentials.Data["key"].(string),
			cloudinary_credentials.Data["secret"].(string))

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
