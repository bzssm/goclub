import (
    "sync"

    deco "github.com/bzssm/goclub/decorator"
)

{{ $decorator := (or .Vars.DecoratorName (printf "Detorated%s" .Interface.Name)) }}

type {{$decorator}} struct {
	preRun  []deco.Handler
	postRun []deco.Handler
	base   {{.Interface.Type}}
	ctxPool sync.Pool
}

func New{{$decorator}}(_base {{.Interface.Type}}, preRun, postRun []deco.Handler) *{{$decorator}} {
	return &{{$decorator}}{
		preRun:  preRun,
		postRun: postRun,
		base:   _base,
		ctxPool: sync.Pool{New: func() interface{} { return &deco.Context{Keys: make(map[string]interface{})} }},
	}
}

{{range $method := .Interface.Methods}}
  // {{$method.Name}} implements {{$.Interface.Type}}
  func (_d {{$decorator}}) {{$method.Declaration}} {
      _ctx := _d.ctxPool.Get().(*deco.Context)
      _ctx.Reset()
      _ctx.FuncName = "{{$method.Name}}"
      _ctx.InputParams = {{$method.ParamsMap}}
      defer _d.ctxPool.Put(_ctx)
      for _, h := range _d.preRun {
      	  h(_ctx)
      }
      defer func() {
      	  for _, h := range _d.postRun {
      	      h(_ctx)
      	  }
      }()
      {{$method.ResultsNames}} = _d.base.{{$method.Call}}
      _ctx.OutputParams = {{$method.ResultsMap}}
      return
  }
{{end}}
