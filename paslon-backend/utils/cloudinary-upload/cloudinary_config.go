package cloudinaryconfig

import (
	"context"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var CloudinaryService *cloudinary.Cloudinary

func UploadToCloudinary(file multipart.File, folder string, publicId string) (string, error) {
	ctx := context.Background()
	cldService, _ := cloudinary.NewFromParams("dxirtmo5t", "572915325566226", "wG56_pbhLPHf2yYtfqH1_3WNiSo")
	paslonPublicId := strings.Replace(strings.ToLower(publicId), " ", "_", -1)
	resp, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folder, PublicID: paslonPublicId})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}

// func UpdateCloudinary(file multipart.File, folder string, publicId string) (string, error) {
// 	ctx := context.Background()
// 	cldService, _ := cloudinary.NewFromParams("dxirtmo5t", "572915325566226", "wG56_pbhLPHf2yYtfqH1_3WNiSo")
// 	paslonPublicId := strings.Replace(strings.ToLower(publicId), " ", "_", -1)
// 	resp, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folder, PublicID: paslonPublicId})
// 	if err != nil {
// 		return "", err
// 	}

// 	return resp.SecureURL, nil
// }

func DeleteCloudinary(folder string, publicId string) (string, error) {
	ctx := context.Background()
	cldService, _ := cloudinary.NewFromParams("dxirtmo5t", "572915325566226", "wG56_pbhLPHf2yYtfqH1_3WNiSo")
	paslonPublicId := strings.Replace(strings.ToLower(publicId), " ", "_", -1)
	resp, err := cldService.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: folder+"/"+paslonPublicId})
	if err != nil {
		return "", err
	}

	return resp.Result, nil
}