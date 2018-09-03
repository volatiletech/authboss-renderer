<h1>Confirm your authenticator code to remove 2fa from your account</h1>
<form action="{{mountpathed "2fa/totp/remove"}}" method="POST">
    {{with .error}}{{.}}<br />{{end}}
    {{with .errors}}{{range .code}}<span>{{.}}</span><br />{{end}}{{end -}}
    <input type="text" class="form-control" name="code" placeholder="Code" autocomplete="off"><br />
    <input type="text" class="form-control" name="recovery_code" placeholder="Recovery Code" autocomplete="off"><br />
    {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
    <button type="submit">Ok</button>
</form>
