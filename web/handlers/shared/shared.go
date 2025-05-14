package shared

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(ctx *fiber.Ctx, view templ.Component) error {
	ctx.Set("Content-Type", "text/html")
	return view.Render(ctx.UserContext(), ctx.Response().BodyWriter())
}

func RenderError(ctx *fiber.Ctx, statuscode int, err error) error {
	errorHtml := fmt.Sprintf(`
	<div id="error-notifier">
		<div class="alert alert-danger d-flex justify-content-between align-items-center fade-in" role="alert">
			<div>
				<span class="icon-holder"><i class="ti-alert"></i></span>
				<strong>Error:</strong> %s
			</div>
			<span class="close-icon" style="cursor: pointer;" hx-get="/core/none" hx-swap="innterHTML" hx-target="#error-notifier" hx-swap="delete" hx-trigger="click"><i class="ti-close"></i></span>
		</div>
	</div>
`, err.Error())

	return ctx.Status(statuscode).SendString(errorHtml)
}
