{{if eq .validate_mode "setup" -}}
    <h1>Setup two-factor authentication</h1>
    <form action="{{mountpathed "2fa/sms/setup"}}" method="POST">
        <input type="text" class="form-control" name="phone_number" placeholder="Phone Number" {{with .sms_phone_number}}value="{{.}}"{{end}}/>
        <button type="submit">Begin Setup</button>
        {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
    </form>
{{else -}}
    {{if eq .validate_mode "validate" -}}
    <h1>Enter your sms code</h1>
    {{else if eq .validate_mode "confirm" -}}
    <h1>Confirm your sms code to complete setup</h1>
    {{else if eq .validate_mode "remove" -}}
    <h1>Confirm your sms code to remove 2fa from your account</h1>
    {{- end}}
    <form action="{{mountpathed (printf "2fa/sms/%s" .validate_mode)}}" method="POST">
        {{with .error}}{{.}}<br />{{end}}
        {{with .errors}}{{range .code}}<span>{{.}}</span><br />{{end}}{{end -}}
        <input type="text" class="form-control" name="code" placeholder="Code" autocomplete="off"><br />
        {{if or (eq .validate_mode "validate") (eq .validate_mode "remove") -}}
        <input type="text" class="form-control" name="recovery_code" placeholder="Recovery Code" autocomplete="off"><br />
        {{end -}}
        {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
        <button type="submit">Resend</button>
        <button type="submit">Ok</button>
    </form>
{{end}}
