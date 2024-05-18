package helper

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func RenderView(ctx context.Context, w http.ResponseWriter, component templ.Component) error {
	return component.Render(ctx, w)
}

func IsHtmxRequest(ctx context.Context) bool {
	htmxRequest := ctx.Value("hx-request")
	return htmxRequest != nil && htmxRequest == true
}

func RandomDigit(min, max int) (int, error) {
	if min >= max {
		return 0, fmt.Errorf("min should be less than max")
	}

	buffer := make([]byte, 1)
	_, err := rand.Read(buffer)
	if err != nil {
		return 0, err
	}

	num := int(buffer[0])
	return min + num%(max-min+1), nil
}
