<h1>Enter your authenticator code</h1>
<form action="{{mountpathed "2fa/totp/validate"}}" method="POST">
    {{with .error}}{{.}}<br />{{end}}
    {{with .errors}}{{range .code}}<span>{{.}}</span><br />{{end}}{{end -}}
    <input type="text" class="form-control" name="code" placeholder="Code" autocomplete="off"><br />
    <input type="text" class="form-control" name="recovery_code" placeholder="Recovery Code" autocomplete="off"><br />
    {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
    <button type="submit">Ok</button>
</form>
