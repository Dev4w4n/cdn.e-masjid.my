package model

type Request struct {
	MimeType       string `json:"mime_type"`
	SubDomain      string `json:"sub_domain"`
	TableReference string `json:"table_reference"`
	MarkAsDelete   bool   `json:"mark_as_delete"`
	Base64File     string `json:"base64_file"`
}
