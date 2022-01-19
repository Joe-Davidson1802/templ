// Code generated by templ@(devel) DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "fmt"
import "time"

func headerTemplate(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<header")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-testid=\"headerTemplate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(name))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</header>")
		if err != nil {
			return err
		}
		return err
	})
}

func footerTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<footer")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-testid=\"footerTemplate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		var_1 := `&copy; `
		_, err = io.WriteString(w, var_1)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(fmt.Sprintf("%d", time.Now().Year())))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</footer>")
		if err != nil {
			return err
		}
		return err
	})
}

func navTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<nav")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-testid=\"navTemplate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<ul>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<li>")
		if err != nil {
			return err
		}
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"/\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		var_2 := `Home`
		_, err = io.WriteString(w, var_2)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</li>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<li>")
		if err != nil {
			return err
		}
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"/posts\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		var_3 := `Posts`
		_, err = io.WriteString(w, var_3)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</li>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</ul>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</nav>")
		if err != nil {
			return err
		}
		return err
	})
}

func layout(name string, content templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<html>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(name))
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
		err = headerTemplate(name).Render(ctx, w)
		if err != nil {
			return err
		}
		err = navTemplate().Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<main>")
		if err != nil {
			return err
		}
		err = content.Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</main>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</body>")
		if err != nil {
			return err
		}
		err = footerTemplate().Render(ctx, w)
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

func homeTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-testid=\"homeTemplate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		var_4 := `Welcome to my website.`
		_, err = io.WriteString(w, var_4)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func postsTemplate(posts []Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-testid=\"postsTemplate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		for _, p := range posts {
			err = templ.RenderScripts(ctx, w, )
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " data-testid=\"postsTemplatePost\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			err = templ.RenderScripts(ctx, w, )
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " data-testid=\"postsTemplatePostName\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, templ.EscapeString(p.Name))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			err = templ.RenderScripts(ctx, w, )
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " data-testid=\"postsTemplatePostAuthor\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, templ.EscapeString(p.Author))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = layout("Home", homeTemplate()).Render(ctx, w)
		if err != nil {
			return err
		}
		return err
	})
}

func posts(posts []Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = layout("Posts", postsTemplate(posts)).Render(ctx, w)
		if err != nil {
			return err
		}
		return err
	})
}

