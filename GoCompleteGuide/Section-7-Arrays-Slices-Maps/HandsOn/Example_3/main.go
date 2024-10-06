package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}
	fmt.Println(websites)

	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

}
