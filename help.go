package main

const appHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   crass {{if .VisibleFlags}}[options]{{end}} <GitHub ID>{{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}`
