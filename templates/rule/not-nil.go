package rule

const notNilTpl = `
		if utf8.RuneCountInString({{ .Key }}) == nil {
			return {{ .Field.Parent.GoIdent.GoName }}ValidationError {
				field:  "{{ .Field.GoName }}",
				reason: "value must not be nil",
			}
		}
`
