package url

import (
	"strings"
	"url-shortner/util"
)

type UrlServiceStruct struct{}

var UrlService = UrlServiceStruct{}

func (s *UrlServiceStruct) Create(urlDto URLDto) (URL, error) {

	var url URL
	var newId uint32 = 1
	err := util.DB.Last(&url).Error
	if err == nil {
		newId = url.ID + 1
	}

	code := toBase62(int64(newId))
	newUrl := URL{
		OriginalUrl: urlDto.OriginalUrl,
		UserId:      uint32(urlDto.UserId),
		Code:        code,
	}

	if err := util.DB.Create(&newUrl).Error; err != nil {
		return URL{}, err
	}

	return newUrl, nil
}

func (s *UrlServiceStruct) FindOneById(id int) (URL, error) {
	var url URL
	err := util.DB.Select(
		"ID",
		"OriginalUrl",
		"Code",
		"CreatedAt",
	).Where("id = ?", id).First(&url).Error
	if err != nil {
		return URL{}, err
	}

	return url, nil
}

func (s *UrlServiceStruct) FindOneByCode(code string) (URL, error) {
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

func (s *UrlServiceStruct) FindAll() ([]URL, error) {
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
