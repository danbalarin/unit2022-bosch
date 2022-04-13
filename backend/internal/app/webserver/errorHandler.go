package webserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"log"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Modified default fiber error handling
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = err.Error()
	}

	if code >= 500 {
		errWithStack, hasStack := err.(stackTracer)
		if hasStack {
			log.Println("Occured error:", err.Error())
			for _, f := range errWithStack.StackTrace() {
				fmt.Printf("%+s:%d\n", f, f)
			}
		} else {
			log.Printf("Occurred unknown error: %s (%T)\n", err, err)
		}
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Status(code).JSON(fiber.Map{
		"status":  code,
		"message": msg,
	})
}
