package url

import (
	"strings"
	common "url-shortner/src/common"
	"url-shortner/src/util"
)

type UrlServiceStruct struct {
	common.ServiceStruct[URL]
}

var UrlService = UrlServiceStruct{
	ServiceStruct: common.ServiceStruct[URL]{},
}

// create url
func (s *UrlServiceStruct) CreateUrl(urlDto URLDto) (URL, error) {
	lastUrl, _ := s.FindLastRecord()
	var newId = lastUrl.ID + 1

	code := toBase62(int64(newId))
	newUrl := URL{
		OriginalUrl: urlDto.OriginalUrl,
		UserID:      uint32(urlDto.UserID),
		Code:        code,
	}

	err := s.CreateOneRecord(&newUrl)
	return newUrl, err
}

// find url by id
func (s *UrlServiceStruct) FindUrlById(id int) (URL, error) {
	url, _ := s.FindOneRecordById(id)

	if url.ID == 0 {
		return URL{}, util.NotFoundErr("url not found")
	}
	return url, nil
}

// find url by code
func (s *UrlServiceStruct) FindUrlByCode(code string) (URL, error) {
	var url URL
	err := util.DB.Select(
		"ID",
		"OriginalUrl",
		"Code",
		"CreatedAt",
	).Where("code = ?", code).First(&url).Error
	if err != nil {
		return URL{}, err
	}

	return url, nil
}

func (s *UrlServiceStruct) FindAllUrls() ([]URL, error) {
	var urls []URL
	err := util.DB.Select(
		"ID",
		"UserId",
		"OriginalUrl",
		"Code",
		"CreatedAt",
	).Find(&urls).Error

	if err != nil {
		return []URL{}, err
	}

	return urls, nil
}

// Convert an integer to a Base62 encoded string
func toBase62(num int64) string {

	const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if num == 0 {
		return "0"
	}

	var result strings.Builder
	base := int64(len(base62Chars))

	// Convert to Base62
	for num > 0 {
		remainder := num % base
		result.WriteByte(base62Chars[remainder])
		num /= base
	}

	// Reverse the string since the conversion adds characters in reverse order
	encoded := result.String()
	return reverse(encoded)
}

// Helper function to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
