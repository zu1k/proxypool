package tool

import (
	"fmt"
	"testing"
)

func TestBase64DecodeString(t *testing.T) {
	str := "OTEuMjA2LjkyLjM4OjI5MzE6b3JpZ2luOnJjNDpwbGFpbjpiRzVqYmk1dmNtY2dNak5zLz9vYmZzcGFyYW09JnJlbWFya3M9NTctNzVhS1o1WVdhWm1GdWNXbGhibWRrWVc1bkxtTnZiUSY9NUwtRTZMLWM1TGljUkEmZ3JvdXA9NTctNzVhS1o1WVdhWm1GdWNXbGhibWRrWVc1bkxtTnZiUSY9VEc1amJpNXZjbWM"
	fmt.Println(Base64DecodeString(str))
	fmt.Println(string([]byte(str)[232]))
}
