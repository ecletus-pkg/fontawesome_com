package fontawesome_com

import (
	"fmt"

	"github.com/ecletus/core"
	"github.com/ecletus/plug"
	"github.com/ecletus/render"
	"github.com/moisespsena/template/html/template"
)

const PTH = "fontawesome.com/fontawesome-free-5.3.1-web"

type Plugin struct {
	plug.EventDispatcher
	RenderKey string
}

func (p *Plugin) RequireOptions() []string {
	return []string{p.RenderKey}
}

func (p *Plugin) Init(options *plug.Options) {
	Render := options.GetInterface(p.RenderKey).(*render.Render)
	Render.RegisterFuncMapMaker("fontawesome_com", func(values *template.FuncValues, Render *render.Render, context *core.Context) error {
		values.Set("fontawesome_css", func(s *template.State) template.HTML {
			pth := render.Context(s).JoinStaticURL(PTH, "css", "all.min.css")
			return template.HTML(fmt.Sprintf(`<link rel="stylesheet" media="all" href="%s">`, pth))
		})
		values.Set("fontawesome_js", func(s *template.State) template.HTML {
			pth := render.Context(s).JoinStaticURL(PTH, "js", "all.min.js")
			return template.HTML(fmt.Sprintf(`<script type="text/javascript" src="%v"></script>`, pth))
		})
		return nil
	})
}
