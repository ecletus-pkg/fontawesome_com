package fontawesome_com

import (
	"fmt"
	"path"

	"github.com/aghape/assets"
	"github.com/gobwas/glob"

	"github.com/aghape/core"
	"github.com/aghape/plug"
	"github.com/aghape/render"
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

func (p *Plugin) OnRegister() {
	assets.Dis(p).OnSyncConfig(func(e *assets.PreRepositorySyncEvent) {
		m1 := glob.MustCompile("static/" + PTH + "/{css,js}/*").Match
		accept := map[string]bool{"all.min.css": true, "all.min.js": true}
		e.Repo.IgnorePath(
			func(pth string) bool {
				if m1(pth) {
					if _, ok := accept[path.Base(pth)]; !ok {
						return true
					}
				}
				return false
			},
			glob.MustCompile("static/"+PTH+"/{scss,less,metadata,LICENSE.txt}").Match,
		)
	})
}

func (p *Plugin) Init(options *plug.Options) {
	Render := options.GetInterface(p.RenderKey).(*render.Render)
	Render.RegisterFuncMapMaker("fontawesome_com", func(values *template.FuncValues, render *render.Render, context *core.Context) error {
		values.Set("fontawesome_css", func() template.HTML {
			pth := context.GenStaticURL(PTH, "css", "all.min.css")
			return template.HTML(fmt.Sprintf(`<link rel="stylesheet" media="all" href="%s">`, pth))
		})
		values.Set("fontawesome_js", func() template.HTML {
			pth := context.GenStaticURL(PTH, "js", "all.min.js")
			return template.HTML(fmt.Sprintf(`<script type="text/javascript" src="%v"></script>`, pth))
		})
		return nil
	})
}
