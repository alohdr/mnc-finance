package usecase

import (
	"mnc-finance/internal/contract/model"
	"mnc-finance/util"
)

func (u *UsecaseImpl) Palindrome(param model.PalindromeRequest) (string, bool) {
	palindrome := util.IsPalindrome(param.Text)

	if palindrome {
		return "is palindrome", palindrome
	}

	return "not palindrome", palindrome
}
