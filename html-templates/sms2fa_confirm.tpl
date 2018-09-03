<h1>Confirm your sms code to complete setup</h1>
<form action="{{mountpathed "2fa/sms/confirm"}}" method="POST">
    {{with .error}}{{.}}<br />{{end}}
    {{with .errors}}{{range .code}}<span>{{.}}</span><br />{{end}}{{end -}}
    <input type="text" class="form-control" name="code" placeholder="Code" autocomplete="off"><br />
    {{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
    <button type="submit">Resend</button>
    <button type="submit">Ok</button>
</form>
