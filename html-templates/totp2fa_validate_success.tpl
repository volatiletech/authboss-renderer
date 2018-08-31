{{if eq .validate_mode "validate" -}}
<h1>Validation successful</h1>
{{else if eq .validate_mode "setup" -}}
<h1>Setup complete</h1>
{{else if eq .validate_mode "confirm" -}}
<h1>Authenticator two-factor enabled</h1>
<p>
    <span>Recovery Codes:</span></br>
    {{range .recovery_codes -}}
    <span>{{.}}</span><br />
    {{end -}}
</p>
{{else if eq .validate_mode "remove" -}}
<h1>Authenticator two-factor successfully removed from account</h1>
{{- end}}
