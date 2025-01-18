package constant

type CtxKey string

const (
	DB          CtxKey = "db"
	Lang        CtxKey = "lang"
	LangID             = "ID"
	LangEN             = "EN"
	LangDefault        = LangEN
)

var (
	AcceptLanguage = map[string]bool{
		LangID: true,
		LangEN: true,
	}
)

const (
	AlphaNumeric        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	AlphaCapitalNumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric             = "0123456789"
	FileExtensionJPG    = ".jpg"
	FileExtensionJPEG   = ".jpeg"
	FileExtensionPNG    = ".png"
	FileExtensionHEIC   = ".heic"
	FileExtensionPDF    = ".pdf"
	FileExtensionZIP    = ".zip"
)

var FileTypeImage = []string{
	FileExtensionHEIC,
	FileExtensionPNG,
	FileExtensionJPG,
	FileExtensionJPEG,
}

var FileTypeRichMedia = []string{
	FileExtensionPNG,
	FileExtensionJPG,
	FileExtensionJPEG,
}

var FileTypeImageAndPDF = append(FileTypeImage, FileExtensionPDF)
