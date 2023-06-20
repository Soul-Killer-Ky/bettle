package util

import (
	pb "beetle/api/filesystem/v1"
)

func GetUploadType(contentType string) pb.UploadType {
	if InArrayString(contentType, []string{
		"image/jpeg",
	}) {
		return pb.UploadType_UT_Image
	} else {
		return pb.UploadType_UT_File
	}
}

func GetSubPath(uploadType pb.UploadType) string {
	if uploadType == pb.UploadType_UT_Image {
		return "image"
	} else {
		return "file"
	}
}
