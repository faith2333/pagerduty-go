package event

type Image struct {
	// The source of the image being attached to the incident or alert. This image must be served via HTTPS.
	Src string `json:"src"`
	// Optional link for the image.
	Href string `json:"href"`
	// Optional alternative text for the image.
	Alt string `json:"alt"`
}
