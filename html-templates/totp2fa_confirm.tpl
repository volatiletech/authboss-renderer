<h1>Confirm your authenticator code to complete setup</h1>
<img src="{{mountpathed "2fa/totp/qr"}}" alt="2fa setup qr code" /><br />
<span>Key: {{.totp_secret}}</span>
<form action="{{mountpathed "2fa/totp/confirm"}}" method="POST">
    {{with .error}}{{.}}<br />{{end}}
    {{with .errors}}{{range .code}}<span>{{.}}</span><br />{{end}}{{end -}}
    <input type="text" class="form-control" name="code" placeholder="Code" autocomplete="off"><br />
    {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
    <button type="submit">Ok</button>
</form>
