# Twitter Lists

{{- range .Lists}}
* [{{ .Name }}](#{{ .Name }})
{{- end}}

Instructions on how to [Run your own](#run-your-own) at the bottom.

{{range .Lists}}
## <a href="https://twitter.com/i/lists/{{ .ID }}">{{ .Name }}</a>
<table>
{{range .Members}}<tr><td><a href="https://twitter.com/{{ .ScreenName }}"><img src="{{ .ProfileImage }}"></a></td><td>
<b><a href="https://twitter.com/{{ .ScreenName }}">@{{ .ScreenName }}</a> ({{ .Name }})</b><br />
<ul>
<li>{{ if .LastTweet }}Last Tweet: {{ .LastTweet }}{{else}}<i>Protected</i>{{end}}</li>
<li>{{ .Description }}</li>
</ul>
</td></tr>
{{end}}
</table>
{{end}}
