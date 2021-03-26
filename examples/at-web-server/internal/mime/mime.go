package mime

import "mime"

var mimeTypes = map[string]string{
	".3gp":   "video/3gpp",
	".mp4":   "video/mp4",
	".mp4v":  "video/mp4",
	".mpg4v": "video/mp4",
	".m4v":   "video/mp4",
	".vrm":   "model/gltf+json",
	".png":   "image/png",
	".jpeg":  "image/jpeg",
	".jpg":   "image/jpeg",
	".jpe":   "image/jpeg",
	".gif":   "image/gif",
}

func init() {
	for ext, typ := range mimeTypes {
		mime.AddExtensionType(ext, typ)
	}
}

func TypeByExtension(ext string) string {
	t := mime.TypeByExtension(ext)
	if t != "" {
		return t
	}

	return "application/octet-stream"
}
