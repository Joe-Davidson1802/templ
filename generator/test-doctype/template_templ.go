// Code generated by templ@(devel) DO NOT EDIT.

package testdoctype

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"

func Layout(title, content string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, `<!doctype html>`)
		if err != nil {
			return err
		}
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<html")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " lang=\"en\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<meta")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " charset=\"UTF-8\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<meta")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " http-equiv=\"X-UA-Compatible\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " content=\"IE=edge\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<meta")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"viewport\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " content=\"width=device-width, initial-scale=1.0\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(title))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<body>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(content))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</body>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</html>")
		if err != nil {
			return err
		}
		return err
	})
}

