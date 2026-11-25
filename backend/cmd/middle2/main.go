package main

import (
	"fmt"
	"slices"
	"strings"
)

type ContentFunc func(content string) (string, error)
type ContentMiddleware func(next ContentFunc) ContentFunc

func main() {
	content := "saya"

	handler := createMiddleware([]ContentMiddleware{
		f1,
		f2,
		f3,
	})

	res, _ := handler(content)

	fmt.Println(res)
}

func createMiddleware(middlewares []ContentMiddleware) ContentFunc {
	var handler ContentFunc = func(content string) (string, error) {
		return content, nil
	}

	slices.Reverse(middlewares)
	for _, mid := range middlewares {
		handler = mid(handler)
	}

	return handler
}

func f1(next ContentFunc) ContentFunc {
	return func(content string) (string, error) {
		jml := 0
		for _, v := range content {
			if v == 'a' || v == 'A' {
				jml++
			}
		}

		kata := content
		for i := 1; i < jml; i++ {
			kata += " " + kata
		}

		return next(content)
	}
}

func f2(next ContentFunc) ContentFunc {
	return func(content string) (string, error) {
		word := strings.Split(content, " ")
		return next(fmt.Sprintf("%d %s", len(word), content))
	}
}

func f3(next ContentFunc) ContentFunc {
	return func(content string) (string, error) {
		abjad := "0123456789abcdefghijklmnopqrstuvwxyzABCDFGHIJKLMNOPQRSTUVWXYZ"
		abjads := strings.Split(abjad, "")
		kata := strings.Split(content, "")
		word := ""

		for _, v := range kata {
			if v == " " {
				word += " "
			}
			if v == "z" {
				word += "a"
				continue
			}
			if v == "Z" {
				word += "A"
				continue
			}
			if v == "9" {
				word += "0"
				continue
			}
			for i, a := range abjads {
				if a == v {
					word += abjads[i+1]
				}
			}
		}

		return next(word)
	}
}
